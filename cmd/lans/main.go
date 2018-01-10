package main

import (
	"encoding/hex"
	"net"

	"github.com/Akagi201/lancaster/mcast"
	log "github.com/sirupsen/logrus"
)

func msgHandler(src *net.UDPAddr, n int, b []byte) {
	log.Println(n, "bytes read from", src)
	log.Println(hex.Dump(b[:n]))
}

func main() {
	mcast.Listen(opts.McastAddr, msgHandler)
}
