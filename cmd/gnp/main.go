package main

import (
	"flag"
	"fmt"
	"soub4i/gnp/internal"
)

const (
	// Version is the current version of the application
	Version = "v1"
	Usage   = `
	gnp - glovo network protocol
	Usage: gnp [options] <command> .
	commands:
		- server: start the server
		- client [ip] -d "Yo my name is mo" : send a message to the server
		- help: show help
		- version: show version
	options:
		- -p: port to listen on
		- -ip: server address
		- -d: data to send
	`
)

// parseFlags parses the command line flags

func parseFlags() {

	flag.Parse()
}

// Main is the entry point for the application

func main() {

	port := flag.Int("p", 8008, "Port to listen on")
	serverAddr := flag.String("ip", "0.0.0.0", "Server address")
	data := flag.String("d", "", "Data to send")
	flag.Parse()
	action := flag.Arg(0)

	options := internal.Options{
		Port:       *port,
		ServerAddr: *serverAddr,
	}

	payload := internal.Payload{
		Data:    *data,
		Version: Version,
	}

	switch action {
	case "server":
		fmt.Println("Starting server on port", options.Port)

		err, server := internal.CreateServer(options)

		internal.Panic("Error creating server", err)

		server.Run()

	case "client":

		err, client := internal.CreateClient(options)

		internal.Panic("Error creating client", err)

		err = client.Send(payload)

		internal.Panic("Error sending data", err)

	case "help":
		fmt.Println(Usage)

	case "version":
		fmt.Println(Version)

	default:
		internal.Panic("Invalid command", nil)
		fmt.Println(Usage)
	}

}
