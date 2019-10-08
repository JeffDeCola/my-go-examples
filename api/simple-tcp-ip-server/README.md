# simple-tcp-ip-server example

_Using the standard `net` package to build a simple tcp server
to handle requests concurrently over ip (tcp/ip)._

Table of Contents,

* [REST vs TCP/IP](https://github.com/JeffDeCola/my-go-examples/tree/master/api/simple-tcp-ip-server#rest-vs-tcpip)
* [HOW IT WORKS](https://github.com/JeffDeCola/my-go-examples/tree/master/api/simple-tcp-ip-server#how-it-works)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/api/simple-tcp-ip-server#run)
  * [USING NETCAT](https://github.com/JeffDeCola/my-go-examples/tree/master/api/simple-tcp-ip-server#using-netcat)
  
[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## REST vs TCP/IP

REST does not have state, whereas tcp has an open connection and you
can assume a lot about the server.

## HOW IT WORKS

The `net` package lets us listen for tcp connections and
handle those requests over that connection.

1. Set which IP and port you would like to listen on,

    ```go
    server, err := net.Listen("tcp", "127.0.0.1:1234")
    ```

1. Create a connection for each request (concurrently),

    ```go
    conn, err := server.Accept()
    go handleRequest(conn)
    ```

1. Create your handler `handleRequest()`,

    ```go
    rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
    cmd, err := rw.ReadString('\n')
    ```

1. Create your handler functions like `handleAdd()`,

   ```go
    handleAdd(rw)
   ```

This illustration may help,

![IMAGE - simple-tcp-ip-server - IMAGE](../../docs/pics/simple-tcp-ip-server.jpg)

## RUN

```bash
go run simple-tcp-ip-server.go \
       requests.go handlers.go
```

Press return to exit.

You can interact with the tcp server in many ways

### USING NETCAT

In another terminal, use the nc command which runs Netcat,
a utility for sending raw data over a network connection.

```bash
netcat -q -1 localhost 1234
nc -q -1 localhost 1234
```

Now you can issue commands such as,

```bash
ADD
4
5
SUBTRACT
5
4
```

You can have as many connections as you like open.
