package internal

import (
	"fmt"
	"net"

	"github.com/golang/protobuf/proto"
)

type Server struct {
	Conn net.UDPConn
}

func CreateServer(option Options) (error, Server) {

	err := option.Validate()

	server := Server{}

	if err != nil {
		fmt.Println("Error: ", err)
		return err, server
	}

	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: option.Port,
		Zone: "",
	})

	server.Conn = *conn

	if err != nil {
		fmt.Println("Error: ", err)
		return err, server
	}

	return nil, server

}

func (s Server) Run() {

	var payload Payload
	buf := make([]byte, 1024)

	for {
		n, addr, err := s.Conn.ReadFromUDP(buf)

		if err != nil {
			fmt.Println("Error => ", err)
			continue
		}

		go func() {

			err = proto.Unmarshal(buf[:n], &payload)

			if err != nil {
				fmt.Println("Error => ", err)
			}

			fmt.Printf("Received data: %s from %s\n", payload.Data, addr)
		}()

	}

}
