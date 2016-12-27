package main

import (
	"net"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/jessevdk/go-flags"
	"github.com/xtaci/smux"
)

var opts struct {
	TcpAddr string `long:"tcp" default:"127.0.0.1:8327" description:"TCP address to connect to"`
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

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		if !strings.Contains(err.Error(), "Usage") {
			log.Fatalf("cli params parse failed, err: %v", err)
		} else {
			return
		}
	}

	// Get a TCP connection
	conn, err := net.Dial("tcp", opts.TcpAddr)
	if err != nil {
		log.Fatalf("TCP Dial failed, err: %v", err)
	}
	defer conn.Close()
	log.Infof("TCP Dial address: %v", opts.TcpAddr)

	// Setup client side of smux
	session, err := smux.Client(conn, nil)
	if err != nil {
		log.Fatalf("smux setup client failed, err: %v", err)
	}

	// Open a new stream
	stream, err := session.OpenStream()
	if err != nil {
		log.Fatalf("session open stream failed, err: %v", err)
	}

	// Stream implements io.ReadWriteCloser
	_, err = stream.Write([]byte("ping"))
	if err != nil {
		log.Fatalf("stream Write failed, err: %v", err)
	}

	buf := make([]byte, 1024)
	n, err := stream.Read(buf)
	if err != nil {
		log.Fatalf("stream Read failed, err: %v", err)
	}
	log.Infof("Read buf: %v", string(buf[:n]))
}
