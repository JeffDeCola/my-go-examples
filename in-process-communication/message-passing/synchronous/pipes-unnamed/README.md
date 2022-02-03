# pipes-unnamed

_tbd._

Other examples using,

* **IN-PROCESS COMMUNICATION**
  * **SHARED MEMORY**
    * ASYNCHRONOUS
    * SYNCHRONOUS
      * [PIPES (UNNAMED)](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/synchronous/pipes-unnamed)
  * **MESSAGE PASSING**
    * ASYNCHRONOUS
      * [CHANNELS (BUFFERED)](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/asynchronous/channels-buffered)
    * SYNCHRONOUS
      * [CHANNELS (UNBUFFERED)](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/synchronous/channels-unbuffered)
* **INTER-PROCESS COMMUNICATION**
  * **SHARED MEMORY**
    * ASYNCHRONOUS
    * SYNCHRONOUS
      * [PIPES (NAMED)](https://github.com/JeffDeCola/my-go-examples/tree/master/communication/inter-process-communication-ipc/synchronous/pipes-named)
  * **MESSAGE PASSING**
    * **OPERATING SYSTEM**
      * ASYNCHRONOUS
        * [MESSAGE QUEUES](https://github.com/JeffDeCola/my-go-examples/tree/master/inter-process-communication-ipc/operating-system/asynchronous/message-queues)
        * [PUB/SUB (NATS)](https://github.com/JeffDeCola/my-go-examples/tree/master/inter-process-communication-ipc/operating-system/asynchronous/pub-sub-nats)
        * [TCP](https://github.com/JeffDeCola/my-go-examples/tree/master/communication/inter-process-communication-ipc/operating-system/asynchronous/tcp)
      * SYNCHRONOUS
        * [gRPC](https://github.com/JeffDeCola/my-go-examples/tree/master/inter-process-communication-ipc/operating-system/synchronous/gRPC)
    * **NETWORK**
      * ASYNCHRONOUS
        * [PUB/SUB (NATS)](https://github.com/JeffDeCola/my-go-examples/tree/master/inter-process-communication-ipc/network/asynchronous/pub-sub-nats)
        * [TCP/IP](https://github.com/JeffDeCola/my-go-examples/tree/master/inter-process-communication-ipc/network/asynchronous/tcp)
      * SYNCHRONOUS
        * [gRPC](https://github.com/JeffDeCola/my-go-examples/tree/master/inter-process-communication-ipc/network/synchronous/gRPC)
        * [REST](https://github.com/JeffDeCola/my-go-examples/tree/master/inter-process-communication-ipc/network/synchronous/rest)

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/communication/in-process-communication/goroutines/asynchronous/channels-buffered#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/communication/in-process-communication/goroutines/asynchronous/channels-buffered#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/communication/in-process-communication/goroutines/asynchronous/channels-buffered#test)

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
go run pipes-unnamed.go
```

## TEST

To create _test files,

```bash
gotests -w -all pipes-unnamed.go
```

To unit test the code,

```bash
go test -cover ./... 
```
