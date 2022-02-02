# grpc

_tbd._

Other examples using,

* **IN-PROCESS COMMUNICATION**
  * **GOROUTINES**
    * **ASYNCHRONOUS**
      * [CHANNELS (BUFFERED)](https://github.com/JeffDeCola/my-go-examples/tree/master/communication/in-process-communication/goroutines/asynchronous/channels-buffered)
    * **SYNCHRONOUS**
      * [CHANNELS (UNBUFFERED)](https://github.com/JeffDeCola/my-go-examples/tree/master/communication/in-process-communication/goroutines/synchronous/channels-unbuffered)
  * **SYNCHRONOUS**
    * [PIPES](https://github.com/JeffDeCola/my-go-examples/tree/master/communication/in-process-communication/synchronous/pipes)
* **INTER-PROCESS COMMUNICATION (IPC)**
  * **ONE-TO-MANY**
    * **ASYNCHRONOUS**
      * [PUB/SUB](https://github.com/JeffDeCola/my-go-examples/tree/master/communication/inter-process-communication-ipc/one-to-many/asynchronous/pub-sub)
  * **ONE-TO-ONE**
    * **ASYNCHRONOUS**
      * [MESSAGE QUEUES](https://github.com/JeffDeCola/my-go-examples/tree/master/communication/inter-process-communication-ipc/one-to-one/asynchronous/message-queues)
      * [TCP/IP](https://github.com/JeffDeCola/my-go-examples/tree/master/communication/inter-process-communication-ipc/one-to-one/asynchronous/tcp-ip)
    * **SYNCHRONOUS**
      * [gRPC](https://github.com/JeffDeCola/my-go-examples/tree/master/communication/inter-process-communication-ipc/one-to-one/synchronous/gRPC)
      * [REST](https://github.com/JeffDeCola/my-go-examples/tree/master/communication/inter-process-communication-ipc/one-to-one/synchronous/rest)

Table of Contents,

* tbd

Documentation and references,

* Refer to the
  [???](https://pkg.go.dev/????)
  package for more info
* This repos [github webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

* One-to-one – Each client request is processed by one Server
* One-to-many – Each Client request is processed by Multiple Servers
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
