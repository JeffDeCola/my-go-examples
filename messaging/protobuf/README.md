# protobuf example

`protobuf`  _is an example of
protocol buffers serialize structured data, useful for messaging._

These are my 3 main example of using protobuf,

* `protobuf` You are here
* [protobuf-NATS-publish-subscribe](https://github.com/JeffDeCola/my-go-examples/tree/master/messaging/protobuf-NATS-publish-subscribe)
* [protobuf-NATS-request-response](https://github.com/JeffDeCola/my-go-examples/tree/master/messaging/protobuf-NATS-request-response)

Refer to my
[protobuf cheat sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/messaging/protobuf-cheat-sheet)
for information on installation and use.

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## STEP 1 - DEFINE .proto FILE

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

## STEP 2 - COMPILE .proto FILE

Compile the protocol buffer file to get the wrappers,

```bash
protoc --go_out=. messages.proto
```

Place this file in the same directory as protobuf.go.

Results in `messages.pb.go` that
implements all messages as go structs and types.

## STEP 3 - IMPLEMENT

Run the code,

```bash
go run protobuf.go messages.pb.go
```

## WHAT IT DOES

Usually, you have two separate processes to show the message being passed, but I
didn't want to have the pipes, so I kept everything inside one process.

### CLIENT - MARSHAL - WRITE/SEND

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

### SERVER - RECEIVE - READ/UNMARSHAL

Now lets create an empty pointer to the
proto stuct, receive the message and unmarshal it.

```go
rcvToken := &Token{}
err = proto.Unmarshal(msg, rcvToken)
```
