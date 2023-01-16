# channels-buffered

_Buffered channels are unidirectional and asynchronous with no blocking._

Other communication examples using,

**I - IN-PROCESS COMMUNICATION**

* **SHARED MEMORY**
  * ASYNCHRONOUS
  * SYNCHRONOUS
    * [pipes-unnamed](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/shared-memory/synchronous/pipes-unnamed)
    * [pipes-unnamed-io](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/shared-memory/synchronous/pipes-unnamed-io)
* **MESSAGE PASSING**
  * ASYNCHRONOUS
    * [channels-buffered](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/message-passing/asynchronous/channels-buffered)
      **<- YOU ARE HERE**
  * SYNCHRONOUS
    * [channels-unbuffered](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/message-passing/synchronous/channels-unbuffered)

**II - INTER-PROCESS COMMUNICATION (IPC)**

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

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/message-passing/asynchronous/channels-buffered#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/message-passing/asynchronous/channels-buffered#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/message-passing/asynchronous/channels-buffered#test)
* [COMMUNICATIONS ILLUSTRATION](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/message-passing/asynchronous/channels-buffered#communications-illustration)

Documentation and references,

* This repos [github webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

**Buffered channels are unidirectional and asynchronous with no blocking.**

You create a buffered channel in go as follows,

```go
var chBuffered = make(chan string, 100) // BUFFERED
```

And place data on that channel as follows,

```go
chBuffered <- data // NO BLOCKING - NO WAITING
```

This illustration  may help,

![IMAGE - channels-buffered.jpg - IMAGE](../../../../docs/pics/in-process-communication/channels-buffered.jpg)

## RUN

Run,

```bash
go run channels-buffered.go
```

```bash
SEND:     I am the data that will be sent 
SEND FAST - WAITS
Chilling....
RECEIVED: I am the data that was received 6 - Took 2 seconds 
Chilling....
SEND SLOW
RECEIVED: I am the data that was received A - Took 1 seconds 
Chilling....
RECEIVED: I am the data that was received B - Took 2 seconds 
RECEIVED: I am the data that was received C - Took 1 seconds 
Chilling....
RECEIVED: I am the data that was received D - Took 2 seconds 
TOTAL TIME: 17 seconds
RECEIVED: I am the data that was received E - Took 1 seconds 
Chilling....
Chilling....
Chilling....
Chilling....
Chilling....
Chilling....
```

## TEST

To create _test files,

```bash
gotests -w -all channels-buffered.go
```

To unit test the code,

```bash
go test -cover ./... 
```

## COMMUNICATION OVERVIEW

### SYNCHRONOUS AND ASYNCHRONOUS COMMUNICATION

Synchronous and asynchronous communication really depends on
what you're referring too. For example, if you have a uni-directional
communication and you are referring to the data transfer. Or if it
a bi-directional communication and you're referring to
the reply.

Here is a high-level view,

* Synchronous – The client expects a timely response and might
  even block while it waits
* Asynchronous – The client doesn’t block while waiting for a response,
  and the response, if any, may not be immediate

![IMAGE - synchronous-and-asynchronous-communication.jpg - IMAGE](../../../../docs/pics/in-process-communication/synchronous-and-asynchronous-communication.jpg)

### CLIENT SERVER COMMUNICATION

![IMAGE - client-server-communication.jpg - IMAGE](../../../../docs/pics/in-process-communication/client-server-communication.jpg)

### IN-PROCESS AND INTER-PROCESS COMMUNICATION

![IMAGE - in-process-and-inter-process-communication.jpg - IMAGE](../../../../docs/pics/in-process-communication/in-process-and-inter-process-communication.jpg)
