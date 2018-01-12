package main

import (
	"time"

	"github.com/jackpal/gateway"
	log "github.com/sirupsen/logrus"
	"github.com/sparrc/go-ping"
)

func main() {
	for {
		ip, err := gateway.DiscoverGateway()
		if err != nil {
			log.Errorf("get gateway ip failed, err: %v", err)
			time.Sleep(1 * time.Minute)
			continue
		}

		log.Infof("gateway ip: %v", ip.String())

		pinger, err := ping.NewPinger(ip.String())
		if err != nil {
			panic(err)
		}
		pinger.Count = 3
		pinger.Run()                            // blocks until finished
		log.Infof("%+v\n", pinger.Statistics()) // get send/receive/rtt stats
		time.Sleep(1 * time.Minute)
	}
}
