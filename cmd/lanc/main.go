package main

import (
	"encoding/json"
	"time"

	"github.com/Akagi201/lancaster/mcast"
	"github.com/Akagi201/utilgo/ips"
	log "github.com/sirupsen/logrus"
)

func ping(addr string) {
	conn, err := mcast.NewBroadcaster(addr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		myIPs, _ := ips.LocalIPv4s()
		b, _ := json.Marshal(myIPs)
		conn.Write(b)
		time.Sleep(5 * time.Second)
	}
}

func main() {
	ping(opts.McastAddr)
}
