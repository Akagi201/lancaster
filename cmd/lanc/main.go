package main

import (
	"time"

	"github.com/Akagi201/lancaster/mcast"
	log "github.com/sirupsen/logrus"
)

func ping(addr string) {
	conn, err := mcast.NewBroadcaster(addr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn.Write([]byte("hello, world\n"))
		time.Sleep(1 * time.Second)
	}
}

func main() {
	ping(opts.McastAddr)
}
