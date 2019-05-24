# protobuf example

`protobuf`  _is an example of
protocol buffers serialize structured data, useful for messaging._

Also check out
[protobuf-NATS-request-response](https://github.com/JeffDeCola/my-go-examples/tree/master/messaging/protobuf-NATS-request-response)
and
[protobuf-NATS-publish-subscribe](https://github.com/JeffDeCola/my-go-examples/tree/master/messaging/protobuf-NATS-publish-subscribe).

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## .proto FILE

Define a protocol buffer file `messages.proto` that
declares the messages that are going to be serialized.

Structure for this example is,

A message is just an aggregate containing a set of typed fields.

```txt
message Person {
    string name = 1;
    int32 age = 2;
    string email = 3;
    string phone = 4;
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
