package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
)

var opts struct {
	LogLevel  string `long:"log_level" default:"info" description:"log level"`
	McastAddr string `long:"mcast_addr" default:"239.0.0.0:9999" description:"multicast address"`
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func init() {
	parser := flags.NewParser(&opts, flags.HelpFlag|flags.PassDoubleDash|flags.IgnoreUnknown)

	_, err := parser.Parse()
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(-1)
	}
}

func init() {
	if level, err := log.ParseLevel(strings.ToLower(opts.LogLevel)); err != nil {
		log.SetLevel(level)
	}

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
}
