# GNP Golovo Network Protocol

This is a simple protocol to send messages between different devices. it is designed to be simple and easy to implement. built on top of UDP and named Glovo because it does not guaranty delivery of the messages Like Glovo with your .

## Usage

To use the protocol you need to implement the `gnp.Message` interface and then you can use the `gnp.Server` and `gnp.Client` to send and receive messages.

```sh
gnp help

        gnp - glovo network protocol
        Usage: gnp [options] <command> .
        commands:
                - server: start the server
                - client [ip] -d "Yo my name is mo" : send a message to the server
                - help: show help
        options:
                - -p: port to listen on
                - -ip: server address
                - -d: data to send
```

## compile

```sh
make compile
```

## Structure

The `cmd` directory will contain the different binaries, separated in directories.

The `bin` directory will contain the generated binaries.

The `internal` directory will contain all your reusable packages.
