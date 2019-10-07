# tcp-ip-server example

_Sending a message back-and-forth between two TCP servers using the
standard `net` package._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## RUN

In two separate terminals, run two instances,

```bash
go run tcp-ip.go server.go client.go -sa 127.0.0.1 -sp 3333 sa 127.0.0.1 -sp 4444 
go run tcp-ip.go client.go client.go -sa 127.0.0.1 -sp 4444 sa -sp 3333 -start
```

Where lp and sp are listen port and send port.

