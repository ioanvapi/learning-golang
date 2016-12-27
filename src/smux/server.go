package main

import (
	"net"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/jessevdk/go-flags"
	"github.com/xtaci/smux"
)

var opts struct {
	ListenAddr string `long:"listen" default:"0.0.0.0:8327" description:"TCP address to listen at"`
}

var log *logrus.Logger

func init() {
	log = logrus.New()
	log.Level = logrus.InfoLevel
	f := new(logrus.TextFormatter)
	f.TimestampFormat = "2006-01-02 15:04:05"
	f.FullTimestamp = true
	log.Formatter = f
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	// Setup server side of smux
	session, err := smux.Server(conn, nil)
	if err != nil {
		log.Errorf("smux setup server session failed, err: %v", err)
		return
	}

	// Accept a stream
	stream, err := session.AcceptStream()
	if err != nil {
		log.Errorf("session accept stream failed, err: %v", err)
		return
	}
	defer stream.Close()

	buf := make([]byte, 1024)
	n, err := stream.Read(buf)
	if err != nil {
		log.Fatalf("stream Read failed, err: %v", err)
	}
	log.Infof("Read buf: %v", string(buf[:n]))

	_, err = stream.Write([]byte("pong"))
	if err != nil {
		log.Fatalf("stream Write failed, err: %v", err)
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

	listener, err := net.Listen("tcp", opts.ListenAddr)
	if err != nil {
		log.Fatalf("TCP Listen failed, err: %v", err)
	}
	log.Infof("TCP listening at: %v", opts.ListenAddr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Errorf("TCP accept failed, err: %v", err)
			continue
		}
		go handleConn(conn)
	}
}
