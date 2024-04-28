package internal

import (
	"fmt"
	"net"
)

// this module is used to setup the protocol connection
// it's a useless function

func Ping(ip net.IP, port int) {

	network := "udp"
	address := fmt.Sprintf("%s:%d", ip.String(), port)

	conn, err := net.Dial(network, address)

	if err != nil {
		Panic("Error connecting to server", err)
	}

	defer conn.Close()

	data := []byte("salam")

	res, err := conn.Write(data)

	if err != nil {
		Panic("Error sending data", err)
	}

	if res != len(data) {
		Panic("Handshake error", nil)
	}
}
