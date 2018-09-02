# protobuf example

`protobuf`  _protocol buffers serialize structured data, useful for messaging._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

Also check out
[protobuf over NATS](https://github.com/JeffDeCola/my-go-examples/tree/master/protobuf-NATS)

## USEFUL FOR RPC

Protobuf serves as a basis for a custom remote procedure
call (RPC) system that can be used for
inter-machine communication.

The whole purpose of using protocol buffers is to serialize
your data so that it can be parsed elsewhere.

## PROTOCOL .proto BUFFER FILE

Define a protocol buffer file `messages.proto` that
declares the messages that are going to be serialized.

Strucure for this example is,

A message is just an aggregate containing a set of typed fields.

```go
message Token {
    string AccessToken = 1;
    string TokenType = 2;
    string RefreshToken = 3;
    int64 ExpiresAt = 4;
}
```

## PROTOBUF COMPILER

In go, get the compiler,

```bash
brew install protobuf
apt-get install protobuf-compiler
protoc --version
```

Compile the protcol buffer file to get the wrappers,

```bash
protoc --go_out=. messages.proto
```

`--go_out=.` is the directory you want to output go code to.

Results in `messages.pb.go` that
implements all messages as go structs and types.

## CLIENT - MARSHAL - WRITE/SEND

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

## SERVER - RECEIVE - READ/UNMARSHAL

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
