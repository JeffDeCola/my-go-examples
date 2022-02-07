# pipes-unnamed

_A pipe provides a unidirectional communication channel._

Other examples using,

* **IN-PROCESS COMMUNICATION**
  * **SHARED MEMORY**
    * ASYNCHRONOUS
    * SYNCHRONOUS
      * [pipes-unnamed](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/shared-memory/synchronous/pipes-unnamed)
        **<- YOU ARE HERE**
      * [pipes-unnamed-io](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/shared-memory/synchronous/pipes-unnamed-io)
  * **MESSAGE PASSING**
    * ASYNCHRONOUS
      * [channels-buffered](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/message-passing/asynchronous/channels-buffered)
    * SYNCHRONOUS
      * [channels-unbuffered](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/message-passing/synchronous/channels-unbuffered)
* **INTER-PROCESS COMMUNICATION (IPC)**
  * **SHARED MEMORY**
    * ASYNCHRONOUS
    * SYNCHRONOUS
      * [pipes-named](https://github.com/JeffDeCola/my-go-examples/tree/master/inter-process-communication-ipc/shared-memory/synchronous/pipes-named)
  * **MESSAGE PASSING**
    * **OPERATING SYSTEM**
      * ASYNCHRONOUS
        * [message-queues](https://github.com/JeffDeCola/my-go-examples/tree/master/inter-process-communication-ipc/message-passing/operating-system/asynchronous/message-queues)
        * [pub-sub-nats-os](https://github.com/JeffDeCola/my-go-examples/tree/master/inter-process-communication-ipc/message-passing/operating-system/asynchronous/pub-sub-nats-os)
        * [tcp](https://github.com/JeffDeCola/my-go-examples/tree/master/inter-process-communication-ipc/message-passing/operating-system/asynchronous/tcp)
      * SYNCHRONOUS
        * [grpc-os](https://github.com/JeffDeCola/my-go-examples/tree/master/inter-process-communication-ipc/message-passing/operating-system/synchronous/grpc-os)
    * **NETWORK**
      * ASYNCHRONOUS
        * [pub-sub-nats-network](https://github.com/JeffDeCola/my-go-examples/tree/master/inter-process-communication-ipc/message-passing/network/asynchronous/pub-sub-nats-network)
        * [tcp-ip](https://github.com/JeffDeCola/my-go-examples/tree/master/inter-process-communication-ipc/message-passing/network/asynchronous/tcp-ip)
      * SYNCHRONOUS
        * [grpc-network](https://github.com/JeffDeCola/my-go-examples/tree/master/inter-process-communication-ipc/message-passing/network/synchronous/grpc-network)
        * [rest](https://github.com/JeffDeCola/my-go-examples/tree/master/inter-process-communication-ipc/message-passing/network/synchronous/rest)

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/shared-memory/synchronous/pipes-unnamed#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/shared-memory/synchronous/pipes-unnamed#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/shared-memory/synchronous/pipes-unnamed#test)
* [COMMUNICATIONS ILLUSTRATION](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/shared-memory/synchronous/pipes-unnamed#communications-illustration)

Documentation and references,

* Refer to the
  [io.Pipe](https://pkg.go.dev/io#Pipe)
  package for more info
* This repos [github webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

* Synchronous – The client expects a timely response and might
  even block while it waits
* Asynchronous – The client doesn’t block while waiting for a response,
  and the response, if any, may not be immediate

A pipe provides a unidirectional inter-process communication channel, where two
ends are involved: reader and writer. Data written to the write end of
the pipe can be read from the read end.

**Pipes in go can be used to connect code expecting an io.Reader with
code expecting an io.Writer.**

## RUN

Run,

```bash
go run pipes-unnamed.go
```

## COMMUNICATIONS ILLUSTRATION

This illustration may help,

![IMAGE - communications-overview.jpg - IMAGE](../../../../docs/pics/in-process-communication/communications-overview.jpg)
