# protobuf example

`protobuf` _is an example of
using protobuf to show `data -> marshal -> snd -> rcv -> unmarshal -> data`._

These are my 5 main example of using protobuf,

* **protobuf** <- Your are here
* [protobuf-NATS-publish-subscribe](https://github.com/JeffDeCola/my-go-examples/tree/master/messaging/protobuf-NATS-publish-subscribe)
* [protobuf-NATS-queue-groups](https://github.com/JeffDeCola/my-go-examples/tree/master/messaging/protobuf-NATS-queue-groups)
* [protobuf-NATS-request-reply](https://github.com/JeffDeCola/my-go-examples/tree/master/messaging/protobuf-NATS-request-reply)
* [protobuf-NATS-request-reply-goroutines](https://github.com/JeffDeCola/my-go-examples/tree/master/messaging/protobuf-NATS-request-reply-goroutines)

Documentation and reference,

* Refer to my
  [protobuf cheat sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/messaging/protobuf-cheat-sheet)
  for information on installation and use.

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## STEP 1 - DEFINE .proto FILE

Define a protocol buffer file `messages.proto` that
declares the messages that are going to be serialized.
A message is just an aggregate containing a set of typed fields.

Structure for this example is,

```go
message Person {
    string name = 1;
    int32 age = 2;
    string email = 3;
    string phone = 4;
    uint32 count = 5;
}
```

## STEP 2 - COMPILE .proto FILE

Compile the protocol buffer file to get the wrappers and
place this file in the same directory as protobuf.go.

```bash
cd protoc
protoc --go_out=. messages.proto
cp messages.pb.go ..
```

Results in `messages.pb.go` that
implements all messages as go structs and types.

## STEP 3 - IMPLEMENT (RUN)

Run the code,

```bash
go run protobuf.go messages.pb.go
```

## THE FLOW

The flow to send data over a pipe is,

* DATA
* MARSHALL
* SEND
* RECEIVE
* UNMARSHAL
* DATA

Usually, you have two separate processes but I kept everything
inside one process to keep it simple.

### DATA

Now lets create the message to send. Create a pointer
to a type Person struct.

```go
sndPerson := &Person{
    Name:  "Jeff",
    Age:   20,
    Email: "blah@blah.com",
    Phone: "555-555-5555",
    Count: 1,
}
```

### MARSHAL

```go
sndMsg, err := proto.Marshal(sndPerson)
```

### SEND

Now lets pretend we're sending the message sndMsg over a pipe.

```go
pipe := sndMsg
```

### RECEIVE

```go
rcvMsg := pipe
```

### UNMARSHAL -> DATA

Now lets create an empty pointer to the
same struct from our protobuf file and unmarshal it.

```go
rcvPerson := &Person{}
err = proto.Unmarshal(rcvMsg, rcvPerson)
```
