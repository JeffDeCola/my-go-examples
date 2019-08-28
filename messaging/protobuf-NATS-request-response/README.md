# protobuf-NATS-request-response example

`protobuf-NATS-request-response` _is an example of
sending a protobuf msg over NATS from a
client to a server using request and response._

These are my 3 main example of using protobuf,

* [protobuf](https://github.com/JeffDeCola/my-go-examples/tree/master/messaging/protobuf)
* [protobuf-NATS-publish-subscribe](https://github.com/JeffDeCola/my-go-examples/tree/master/messaging/protobuf-NATS-publish-subscribe)
* **protobuf-NATS-request-response** You are here

Documentation and reference,

* My [protobuf cheat sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/messaging/protobuf-cheat-sheet)
* My [NATS cheat sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/messaging/NATS-cheat-sheet)
* Official [NATS go client library](https://github.com/nats-io/nats.go)
  at github

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## START YOUR NATS SERVER

Time to use NATS as a pipe.  First, lets start your NATS server,

```bash
nats-server -v
nats-server -DV -p 4222 -a 127.0.0.1
```

Where -DV is both debug and trace log.

## GET NATS GO CLIENT LIBRARY

You must have this library to use go,

```go
go get -v -u github.com/nats-io/nats.go/
```



## BASED ON PREVIOUS EXAMPLE protobuf-NATS

This example will add workers(servers) and the client will request
and get a response back.

Refer to
[protobuf-NATS-publish-subscribe](https://github.com/JeffDeCola/my-go-examples/tree/master/protobuf-NATS-publish-subscribe).

## HIGH-LEVEL-VIEW

![IMAGE - protobuf-NATS-request-response - IMAGE](../../docs/pics/protobuf-NATS-request-response.jpg)
