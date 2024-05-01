# GNP Golovo Network Protocol

This is a simple network protocol to send messages between different devices. built on top of UDP and named `Glovoo` because it does not guaranty delivery of the messages like Glovoo with your orders.

it's uses Google's protobuf to serialize and deserialize the messages. because why not.

## Usage

```sh
gnp help

        gnp - Glovoo network protocol
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

### Server

```sh
gnp-linux-amd64 server -p=8080
```

### Client

```sh
gnp-linux-amd64 client -p=8080 -d="Yo my name is mo"
```

## compile

```sh
make compile
```
