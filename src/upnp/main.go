package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/NebulousLabs/go-upnp"
)

func main() {
	// connect to router
	d, err := upnp.Discover()
	if err != nil {
		log.Fatal(err)
	}

	// discover external IP
	ip, err := d.ExternalIP()
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("Your external IP is:%v", ip)
}