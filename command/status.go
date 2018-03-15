package main

import (
	"net"
	"strconv"
)

var targetAds string

func initStatus() {
	targetAds = net.JoinHostPort(cfg.TargetAddress, strconv.Itoa(int(cfg.TargetPort)))
}
