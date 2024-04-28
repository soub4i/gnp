package internal

import (
	"fmt"
	"net"
	"os"
)

func Panic(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err)
		os.Exit(1)
	}
}

// parse string ip to net.IP

func ParseIP(ip string) (net.IP, error) {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return nil, fmt.Errorf("Invalid IP")
	}
	return parsedIP, nil
}
