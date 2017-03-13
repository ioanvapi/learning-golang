package main

import (
	"crypto/sha1"
	"encoding/csv"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/getlantern/smux"
	"github.com/golang/snappy"
	"github.com/jessevdk/go-flags"
	"github.com/pkg/errors"
	"golang.org/x/crypto/pbkdf2"
)

var opts struct {
	KcpAddr      string `long:"kcp" default:"127.0.0.1:8327" description:"KCP address to connect to"`
	Key          string `long:"key" default:"it's a secrect" description:"pre-shared secret between client and server"`
	Crypt        string `long:"crypt" default:"aes" description:"tea, xor, none, aes-128, aes-192, blowfish, twofish, cast5, 3des, xtea, salsa20 aes"`
	Mode         string `long:"mode" default:"fast" description:"profiles: normal, fast, fast2, fast3, manual"`
	Conn         int    `long:"conn" default:"1" description:"set num of UDP connections to server"`
	AutoExpire   int    `long:"autoexpire" default:"0" description:"set auto expiration time(in seconds) for a single UDP connection, 0 to disable"`
	MTU          int    `long:"mtu" default:"1350" description:"set maximum transmission unit for UDP packets"`
	SndWnd       int    `long:"sndwnd" default:"128" description:"set send window size(num of packets)"`
	RcvWnd       int    `long:"rcvwnd" default:"512" description:"set receive window size(num of packets)"`
	DataShard    int    `long:"datashard" default:"10" description:"set reed-solomon erasure coding - datashard"`
	ParityShard  int    `long:"parityshard" default:"3" description:"set reed-solomon erasure coding - datashard"`
	DSCP         int    `long:"dscp" default:"0" description:"set DSCP(6bit)"`
	NoComp       bool   `long:"nocomp" description:"disable compression"`
	AckNodelay   bool   `long:"acknodelay" description:"flush ack immediately when a packet is received"`
	NoDelay      int    `long:"nodelay" default:"0" description:"nodelay mode"`
	Interval     int    `long:"interval" default:"40" description:"interval"`
	Resend       int    `long:"resend" default:"0" description:"resend"`
	NoCongestion int    `long:"nc" default:"0" description:"no congestion"`
	SockBuf      int    `long:"sockbuf" default:"4194304" description:"socket buffer size in bytes"`
	KeepAlive    int    `long:"keepalive" default:"10" description:"nat keepalive interval in seconds"`
	SnmpLog      string `long:"snmplog" default:"" description:"collect snmp to file, aware of timeformat in golang, like: ./snmp-20060102.log"`
	SnmpPeriod   int    `long:"snmpperiod" default:"60" description:"snmp collect period, in seconds"`
}

const (
	saltPbkdf2     = "Akagi201"
	maxScavengeTTL = 10 * time.Minute
)

var log *logrus.Logger

func init() {
	log = logrus.New()
	log.Level = logrus.InfoLevel
	f := new(logrus.TextFormatter)
	f.TimestampFormat = "2006-01-02 15:04:05"
	f.FullTimestamp = true
	log.Formatter = f
}

type compStream struct {
	conn net.Conn
	w    *snappy.Writer
	r    *snappy.Reader
}

func (c *compStream) Read(p []byte) (int, error) {
	return c.r.Read(p)
}

func (c *compStream) Write(p []byte) (int, error) {
	var n int
	var err error
	n, err = c.w.Write(p)
	if err != nil {
		return n, err
	}
	err = c.w.Flush()
	return n, err
}

func (c *compStream) Close() error {
	return c.conn.Close()
}

func newCompStream(conn net.Conn) *compStream {
	c := new(compStream)
	c.conn = conn
	c.w = snappy.NewBufferedWriter(conn)
	c.r = snappy.NewReader(conn)
	return c
}

type scavengeSession struct {
	session *smux.Session
	ttl     time.Time
}

func scavenger(ch chan *smux.Session) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	var sessionList []scavengeSession
	for {
		select {
		case sess := <-ch:
			sessionList = append(sessionList, scavengeSession{sess, time.Now()})
		case <-ticker.C:
			var newList []scavengeSession
			for k := range sessionList {
				s := sessionList[k]
				if s.session.NumStreams() == 0 || s.session.IsClosed() || time.Since(s.ttl) > maxScavengeTTL {
					log.Infof("session scavenged")
					s.session.Close()
				} else {
					newList = append(newList, sessionList[k])
				}
			}
			sessionList = newList
		}
	}
}

func snmpLogger(path string, interval int) {
	if path == "" || interval == 0 {
		return
	}
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			f, err := os.OpenFile(time.Now().Format(path), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
				log.Errorf("Open snmp log failed, err: %v", err)
				return
			}
			w := csv.NewWriter(f)
			// write header in empty file
			if stat, err := f.Stat(); err == nil && stat.Size() == 0 {
				if err := w.Write(append([]string{"Unix"}, kcp.DefaultSnmp.Header()...)); err != nil {
					log.Errorf("Write snmp log snmp header failed, err: %v", err)
				}
			}
			if err := w.Write(append([]string{fmt.Sprint(time.Now().Unix())}, kcp.DefaultSnmp.ToSlice()...)); err != nil {
				log.Errorf("Write snmp log timestamp failed, err: %v", err)
			}
			kcp.DefaultSnmp.Reset()
			w.Flush()
			f.Close()
		}
	}
}

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		if !strings.Contains(err.Error(), "Usage") {
			log.Fatalf("cli params parse failed, err: %v", err)
		} else {
			return
		}
	}

	rand.Seed(int64(time.Now().Nanosecond()))

	switch opts.Mode {
	case "normal":
		opts.NoDelay, opts.Interval, opts.Resend, opts.NoCongestion = 0, 30, 2, 1
	case "fast":
		opts.NoDelay, opts.Interval, opts.Resend, opts.NoCongestion = 0, 20, 2, 1
	case "fast2":
		opts.NoDelay, opts.Interval, opts.Resend, opts.NoCongestion = 1, 20, 2, 1
	case "fast3":
		opts.NoDelay, opts.Interval, opts.Resend, opts.NoCongestion = 1, 10, 2, 1
	case "manual":
		log.Infof("manual mode, please manual set nodelay, interval, resend, nocongestion")
	default:
		log.Fatalf("Unsupported mode: %v", opts.Mode)
	}

	pass := pbkdf2.Key([]byte(opts.Key), []byte(saltPbkdf2), 4096, 32, sha1.New)

	var block kcp.BlockCrypt
	switch opts.Crypt {
	case "tea":
		block, _ = kcp.NewTEABlockCrypt(pass[:16])
	case "xor":
		block, _ = kcp.NewSimpleXORBlockCrypt(pass)
	case "none":
		block, _ = kcp.NewNoneBlockCrypt(pass)
	case "aes-128":
		block, _ = kcp.NewAESBlockCrypt(pass[:16])
	case "aes-192":
		block, _ = kcp.NewAESBlockCrypt(pass[:24])
	case "blowfish":
		block, _ = kcp.NewBlowfishBlockCrypt(pass)
	case "twofish":
		block, _ = kcp.NewTwofishBlockCrypt(pass)
	case "cast5":
		block, _ = kcp.NewCast5BlockCrypt(pass[:16])
	case "3des":
		block, _ = kcp.NewTripleDESBlockCrypt(pass[:24])
	case "xtea":
		block, _ = kcp.NewXTEABlockCrypt(pass[:16])
	case "salsa20":
		block, _ = kcp.NewSalsa20BlockCrypt(pass)
	case "aes":
		block, _ = kcp.NewAESBlockCrypt(pass)
	default:
		log.Fatalf("Unsupported crypt: %v", opts.Crypt)
	}

	log.Infof("Configs: %+v", opts)

	smuxConfig := smux.DefaultConfig()
	smuxConfig.MaxReceiveBuffer = opts.SockBuf

	createConn := func() (*smux.Session, error) {
		kcpconn, err := kcp.DialWithOptions(opts.KcpAddr, block, opts.DataShard, opts.ParityShard)
		if err != nil {
			return nil, errors.Wrap(err, "createConn()")
		}

		kcpconn.SetStreamMode(true)
		kcpconn.SetNoDelay(opts.NoDelay, opts.Interval, opts.Resend, opts.NoCongestion)
		kcpconn.SetWindowSize(opts.SndWnd, opts.RcvWnd)
		kcpconn.SetMtu(opts.MTU)
		kcpconn.SetACKNoDelay(opts.AckNodelay)
		kcpconn.SetKeepAlive(opts.KeepAlive)

		if err := kcpconn.SetDSCP(opts.DSCP); err != nil {
			log.Errorf("SetDSCP failed, err: %v", err)
		}

		if err := kcpconn.SetReadBuffer(opts.SockBuf); err != nil {
			log.Errorf("SetReadBuffer failed, err: %v", err)
		}

		if err := kcpconn.SetWriteBuffer(opts.SockBuf); err != nil {
			log.Errorf("SetWriteBuffer failed, err: %v", err)
		}

		// stream multiplex
		var session *smux.Session
		if opts.NoComp {
			session, err = smux.Client(kcpconn, smuxConfig)
		} else {
			session, err = smux.Client(newCompStream(kcpconn), smuxConfig)
		}

		if err != nil {
			return nil, errors.Wrap(err, "createConn()")
		}
		return session, nil
	}

	// wait until a connection is ready
	waitConn := func() *smux.Session {
		for {
			if session, err := createConn(); err != nil {
				return session
			} else {
				time.Sleep(time.Second)
			}
		}
	}

	muxes := make([]struct {
		session *smux.Session
		ttl     time.Time
	}, opts.Conn)

	for k := range muxes {
		sess, err := createConn()
		if err != nil {
			log.Fatalf("create connection failed, err: %v", err)
		}
		muxes[k].session = sess
		muxes[k].ttl = time.Now().Add(time.Duration(opts.AutoExpire) * time.Second)
	}

	chScavenger := make(chan *smux.Session, 128)
	go scavenger(chScavenger)
	go snmpLogger(opts.SnmpLog, opts.SnmpPeriod)
	rr := uint16(0)

	for {
		idx := rr % uint16(opts.Conn)

		// do auto expiration && reconnection
		if muxes[idx].session.IsClosed() || (opts.AutoExpire > 0 && time.Now().After(muxes[idx].ttl)) {
			chScavenger <- muxes[idx].session
			muxes[idx].session = waitConn()
			muxes[idx].ttl = time.Now().Add(time.Duration(opts.AutoExpire) * time.Second)
		}

		s, err := muxes[idx].session.OpenStream()
		if err != nil {
			log.Errorf("session open stream failed, err: %v", err)
			continue
		}

		if _, err := s.Write([]byte("PING")); err != nil {
			log.Errorf("session Write failed, err: %v", err)
			continue
		}

		buf := make([]byte, opts.SockBuf)
		n, err := s.Read(buf)
		if err != nil {
			log.Errorf("session Read failed, err: %v", err)
			continue
		}

		log.Infof("Read buffer: %v", string(buf[:n]))

		time.Sleep(3 * time.Second)
		rr++
	}
}
