# pipes-unnamed-simple

_A pipe provides a uni-directional communication channel.
This is a very simple example of an unnamed pipe._

Other communication examples using,

**I - IN-PROCESS COMMUNICATION**

* **SHARED MEMORY**
  * ASYNCHRONOUS
  * SYNCHRONOUS
    * [pipes-unnamed-simple](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/shared-memory/synchronous/pipes-unnamed-simple)
      **<- YOU ARE HERE**
    * [pipes-unnamed](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/shared-memory/synchronous/pipes-unnamed)
    * [pipes-unnamed-io](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/shared-memory/synchronous/pipes-unnamed-io)
* **MESSAGE PASSING**
  * ASYNCHRONOUS
    * [channels-buffered](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/message-passing/asynchronous/channels-buffered)
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

* [PIPES OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/shared-memory/synchronous/pipes-unnamed-simple#pipes-overview)
* [CODE](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/shared-memory/synchronous/pipes-unnamed-simple#code)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/shared-memory/synchronous/pipes-unnamed-simple#run)
* [IN-PROCESS AND INTER-PROCESS COMMUNICATION OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/shared-memory/synchronous/pipes-unnamed-simple#in-process-and-inter-process-communication-overview)

Documentation and references,

* Refer to the
  [io.Pipe](https://pkg.go.dev/io#Pipe)
  package for more info
* Refer to
  [my cheat sheets](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/communication/in-process-and-inter-process-communications-ipc-overview-cheat-sheet)
  for a communications overview
* This repos [github webpage](https://jeffdecola.github.io/my-go-examples/)

## PIPES OVERVIEW

A pipe provides a uni-directional in-process communication channel, where two
ends are involved: reader and writer. Data written to the write end of
the pipe can be read from the read end. A named pipe (FIFO) can be accessed from
different processes whereas an unnamed pipe can be accessed from the same process.

* **MACHINE**: SAME
* **PROCESSES**: IN-PROCESS
* **TYPE**: SHARED MEMORY
* **TIMING**: SYNCHRONOUS
* **DIRECTION**: UNI-DIRECTIONAL

Pipes in go can be used to connect code expecting an io.Reader with
code expecting an io.Writer.

![IMAGE - pipes-unnamed-named.jpg - IMAGE](../../../../docs/pics/in-process-communication/pipes-unnamed-named.jpg)

## CODE

Create the pipe using the "io" package,

```go
pr, pw := io.Pipe()
```

Send data (write) to the pipe and close it,

```go
_, err := pw.Write([]byte(data))
pw.Close()
```

This writes the data to memory that is shared.

Now when ready, receive data (read) from the pipe (i.e. read from memory),

```go
buffer := make([]byte, 100)
_, err := pr.Read(buffer)
```

## RUN

Run,

```bash
go run pipes-unnamed.go
```

## IN-PROCESS AND INTER-PROCESS COMMUNICATION OVERVIEW

Refer to
[my cheat sheets](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/communication/in-process-and-inter-process-communications-ipc-overview-cheat-sheet)
for a more thorough communications overview.

![IMAGE - in-process-and-inter-process-communication.jpg - IMAGE](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/docs/pics/in-process-and-inter-process-communication.jpg)
