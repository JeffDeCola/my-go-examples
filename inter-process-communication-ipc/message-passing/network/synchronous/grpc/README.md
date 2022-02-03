# grpc

_tbd._

Other examples using,

* **IN-PROCESS COMMUNICATION**
  * **SHARED MEMORY**
    * ASYNCHRONOUS
    * SYNCHRONOUS
      * [PIPES (UNNAMED)](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/shared-memory/synchronous/pipes-unnamed)
  * **MESSAGE PASSING**
    * ASYNCHRONOUS
      * [CHANNELS (BUFFERED)](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/message-passing/asynchronous/channels-buffered)
    * SYNCHRONOUS
      * [CHANNELS (UNBUFFERED)](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/message-passing/synchronous/channels-unbuffered)
* **INTER-PROCESS COMMUNICATION (IPC)**
  * **SHARED MEMORY**
    * ASYNCHRONOUS
    * SYNCHRONOUS
      * [PIPES (NAMED)](https://github.com/JeffDeCola/my-go-examples/tree/master/communication/inter-process-communication-ipc/shared-memory/synchronous/pipes-named)
  * **MESSAGE PASSING**
    * **OPERATING SYSTEM**
      * ASYNCHRONOUS
        * [MESSAGE QUEUES](https://github.com/JeffDeCola/my-go-examples/tree/master/inter-process-communication-ipc/message-passing/operating-system/asynchronous/message-queues)
        * [PUB/SUB (NATS)](https://github.com/JeffDeCola/my-go-examples/tree/master/inter-process-communication-ipc/message-passing/operating-system/asynchronous/pub-sub-nats)
        * [TCP](https://github.com/JeffDeCola/my-go-examples/tree/master/communication/inter-process-communication-ipc/message-passing/operating-system/asynchronous/tcp)
      * SYNCHRONOUS
        * [gRPC](https://github.com/JeffDeCola/my-go-examples/tree/master/inter-process-communication-ipc/message-passing/operating-system/synchronous/gRPC)
    * **NETWORK**
      * ASYNCHRONOUS
        * [PUB/SUB (NATS)](https://github.com/JeffDeCola/my-go-examples/tree/master/inter-process-communication-ipc/message-passing/network/asynchronous/pub-sub-nats)
        * [TCP/IP](https://github.com/JeffDeCola/my-go-examples/tree/master/inter-process-communication-ipc/message-passing/network/asynchronous/tcp)
      * SYNCHRONOUS
        * [gRPC](https://github.com/JeffDeCola/my-go-examples/tree/master/inter-process-communication-ipc/message-passing/network/synchronous/gRPC)
        * [REST](https://github.com/JeffDeCola/my-go-examples/tree/master/inter-process-communication-ipc/message-passing/network/synchronous/rest)

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/inter-process-communication-ipc/message-passing/network/synchronous/grpc#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/inter-process-communication-ipc/message-passing/network/synchronous/grpc#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/inter-process-communication-ipc/message-passing/network/synchronous/grpc#test)
* [COMMUNICATIONS ILLUSTRATION](https://github.com/JeffDeCola/my-go-examples/tree/master/inter-process-communication-ipc/message-passing/network/synchronous/grpc#communications-illustration)

Documentation and references,

* Refer to the
  [???](https://pkg.go.dev/????)
  package for more info
* This repos [github webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

* Synchronous – The client expects a timely response and might
  even block while it waits
* Asynchronous – The client doesn’t block while waiting for a response,
  and the response, if any, may not be immediate

## RUN

Run,

```bash
go run grpc.go
```

## TEST

To create _test files,

```bash
gotests -w -all grpc.go
```

To unit test the code,

```bash
go test -cover ./... 
```

## COMMUNICATIONS ILLUSTRATION

This illustration may help,

![IMAGE - communications-overview.jpg - IMAGE](../../../docs/pics/in-process-communications/communications-overview.jpg)
