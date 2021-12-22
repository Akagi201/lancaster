package main

import (
	"net"

	"github.com/Akagi201/lancaster/mcast"
	log "github.com/sirupsen/logrus"
)

func msgHandler(src *net.UDPAddr, n int, b []byte) {
	log.Infof("%v bytes read from %v", n, src)
	log.Infof("Discovered %v", string(b))
}

func main() {
	log.Infof("Lancaster server listening on %v", opts.McastAddr)
	mcast.Listen(opts.McastAddr, msgHandler)
}
