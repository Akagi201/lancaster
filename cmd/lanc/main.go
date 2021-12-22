package main

import (
	"encoding/json"
	"net"
	"time"

	"github.com/Akagi201/lancaster/mcast"
	"github.com/Akagi201/utilgo/ips"
	"github.com/avast/retry-go"
	log "github.com/sirupsen/logrus"
)

func ping(addr string) {
	log.Infof("Lanscaster client sending IPs to %v", addr)
	var conn *net.UDPConn
	var err error
	if err = retry.Do(
		func() error {
			conn, err = mcast.NewBroadcaster(addr)
			return err
		},
		retry.Attempts(10000),
		retry.Delay(10*time.Second),
	); err != nil {
		log.Fatal(err)
	}

	log.Infof("%v multicast address connected", addr)

	for {
		myIPs, _ := ips.LocalIPv4s()
		b, _ := json.Marshal(myIPs)
		conn.Write(b)
		time.Sleep(10 * time.Second)
	}
}

func main() {
	ping(opts.McastAddr)
}
