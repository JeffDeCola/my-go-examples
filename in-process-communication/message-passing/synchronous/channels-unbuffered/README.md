# channels-unbuffered

_Unbuffered channels are unidirectional and synchronous with blocking._

Other examples using,

* **IN-PROCESS COMMUNICATION**
  * **SHARED MEMORY**
    * ASYNCHRONOUS
    * SYNCHRONOUS
      * [pipes-unnamed](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/shared-memory/synchronous/pipes-unnamed)
      * [pipes-unnamed-io](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/shared-memory/synchronous/pipes-unnamed-io)
  * **MESSAGE PASSING**
    * ASYNCHRONOUS
      * [channels-buffered](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/message-passing/asynchronous/channels-buffered)
    * SYNCHRONOUS
      * [channels-unbuffered](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/message-passing/synchronous/channels-unbuffered)
        **<- YOU ARE HERE**
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

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/message-passing/synchronous/channels-unbuffered#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/message-passing/synchronous/channels-unbuffered#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/message-passing/synchronous/channels-unbuffered#test)
* [COMMUNICATIONS ILLUSTRATION](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/message-passing/synchronous/channels-unbuffered#communications-illustration)

Documentation and references,

* This repos [github webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

* Synchronous – The client expects a timely response and might
  even block while it waits
* Asynchronous – The client doesn’t block while waiting for a response,
  and the response, if any, may not be immediate

**Unbuffered channels are unidirectional and synchronous with blocking.**

This illustration  may help,

![IMAGE - channels-unbuffered.jpg - IMAGE](../../../../docs/pics/in-process-communication/channels-unbuffered.jpg)

## RUN

Run,

```bash
go run channels-unbuffered.go
```

```bash
SEND:     I am the data that will be sent 
SEND FAST - WAITS
RECEIVED: I am the data that was received 1 - Took 0 seconds 
RECEIVED: I am the data that was received 2 - Took 2 seconds 
RECEIVED: I am the data that was received 3 - Took 4 seconds 
RECEIVED: I am the data that was received 4 - Took 4 seconds 
RECEIVED: I am the data that was received 5 - Took 4 seconds 
RECEIVED: I am the data that was received 6 - Took 4 seconds 
SEND SLOW
RECEIVED: I am the data that was received A - Took 0 seconds 
RECEIVED: I am the data that was received B - Took 0 seconds 
RECEIVED: I am the data that was received C - Took 0 seconds 
RECEIVED: I am the data that was received D - Took 0 seconds 
RECEIVED: I am the data that was received E - Took 0 seconds 
```

Note how it takes 4 seconds, because it's a 2 hop process.

## TEST

To create _test files,

```bash
gotests -w -all channels-unbuffered.go
```

To unit test the code,

```bash
go test -cover ./... 
```

## COMMUNICATIONS ILLUSTRATION

This illustration may help,

![IMAGE - communications-overview.jpg - IMAGE](../../../../docs/pics/in-process-communication/communications-overview.jpg)
