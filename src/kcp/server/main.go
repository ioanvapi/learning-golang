package main

import (
	"crypto/sha1"
	"encoding/csv"
	"fmt"
	"io"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/getlantern/smux"
	"github.com/golang/snappy"
	"github.com/jessevdk/go-flags"
	"github.com/xtaci/kcp-go"
	"golang.org/x/crypto/pbkdf2"
)

var opts struct {
	ListenAddr   string `long:"listen" default:"0.0.0.0:8327" description:"KCP address to listen at"`
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
	saltPbkdf2 = "Akagi201"
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

// handle multiplex-ed connection
func handleMux(conn io.ReadWriteCloser) {
	// stream multiplex
	smuxConfig := smux.DefaultConfig()
	smuxConfig.MaxReceiveBuffer = opts.SockBuf
	mux, err := smux.Server(conn, smuxConfig)
	if err != nil {
		log.Println(err)
		return
	}
	defer mux.Close()
	for {
		s, err := mux.AcceptStream()
		if err != nil {
			log.Println(err)
			return
		}
		buf := make([]byte, opts.SockBuf)
		n, err := s.Read(buf)
		if err != nil {
			log.Errorf("session Read failed, err: %v", err)
			continue
		}
		log.Infof("Read buf: %v", string(buf[:n]))

		if _, err := s.Write([]byte("PONG")); err != nil {
			log.Errorf("session Write failed, err: %v", err)
			continue
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

	listener, err := kcp.ListenWithOptions(opts.ListenAddr, block, opts.DataShard, opts.ParityShard)
	if err != nil {
		log.Fatalf("KCP listen failed, err: %v", err)
	}

	if err := listener.SetDSCP(opts.DSCP); err != nil {
		log.Errorf("SetDSCP failed, err: %v", err)
	}

	if err := listener.SetReadBuffer(opts.SockBuf); err != nil {
		log.Errorf("SetReadBuffer failed, err: %v", err)
	}

	if err := listener.SetWriteBuffer(opts.SockBuf); err != nil {
		log.Errorf("SetWriteBuffer failed, err: %v", err)
	}

	go snmpLogger(opts.SnmpLog, opts.SnmpPeriod)

	for {
		conn, err := listener.AcceptKCP()
		if err != nil {
			log.Errorf("KCP accept failed, err: %v", err)
			continue
		}

		log.Infof("remote address: %v", conn.RemoteAddr())
		conn.SetStreamMode(true)
		conn.SetNoDelay(opts.NoDelay, opts.Interval, opts.Resend, opts.NoCongestion)
		conn.SetMtu(opts.MTU)
		conn.SetWindowSize(opts.SndWnd, opts.RcvWnd)
		conn.SetACKNoDelay(opts.AckNodelay)
		conn.SetKeepAlive(opts.KeepAlive)

		if opts.NoComp {
			go handleMux(conn)
		} else {
			go handleMux(newCompStream(conn))
		}
	}
}
