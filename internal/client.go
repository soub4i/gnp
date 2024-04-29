package internal

import (
	"fmt"
	"net"

	"github.com/golang/protobuf/proto"
)

type Client struct {
	Conn net.UDPConn
}

func CreateClient(option Options) (error, Client) {

	var client Client

	var ip net.IP
	var err error

	if ip, err = ParseIP(option.ServerAddr); err != nil {
		fmt.Println("Error: ", err)
		ip = net.IPv4(0, 0, 0, 0)
	}

	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   ip,
		Port: option.Port,
		Zone: "",
	})

	if err != nil {
		fmt.Println("Error: ", err)
		return err, client
	}
	Panic("Error creating client", err)

	client = Client{
		Conn: *conn,
	}

	return nil, client
}

func (c Client) Send(payload *Payload) error {

	p := payload.Clone()

	data, err := proto.Marshal(p)

	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	_, err = c.Conn.Write([]byte(data))

	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	return nil
}
