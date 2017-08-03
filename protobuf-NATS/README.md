# protobuf-NATS example

`protobuf-NATS`  _sends a protobuf msg over NATS from a client to a server._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## PROTOCOL .proto BUFFER FILE

```go
message Token {
    string AccessToken = 1;
    string TokenType = 2;
    string RefreshToken = 3;
    int64 ExpiresAt = 4;
}
```

## PROTOBUF COMPILER

Compile the protcol buffer file to get the wrappers,

```bash
protoc --go_out=. messages.proto
```

## CLIENT.GO - MARSHAL - WRITE/SEND

Now lets create the message `msg` to send. Create a pointer
to a type Token struct and fill it with data.

```go
token := &Token{
    AccessToken:  "the access token",
    TokenType:    "this",
    RefreshToken: "and the refresh token",
    ExpiresAt:    5,
}
```

Then serialize it up,

```go
msg, err := proto.Marshal(token)
```

## SERVER.GO - RECEIVE - READ/UNMARSHAL

Now lets create an empty pointer to the
proto stuct, receive the message and unmarshal it.

```go
rcvToken := &Token{}
err = proto.Unmarshal(msg, rcvToken)
```

## RUN

```go
go run read.go messages.pb.go
```
