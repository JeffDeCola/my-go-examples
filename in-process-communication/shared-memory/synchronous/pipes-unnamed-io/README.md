# PIPE UNNAMED IO EXAMPLE

_This example of an unnamed pipe connects an io.Writer and io.Reader._

**I - IN-PROCESS COMMUNICATION**

* **SHARED MEMORY**
  * ASYNCHRONOUS
  * SYNCHRONOUS
    * [pipes-unnamed-io](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/shared-memory/synchronous/pipes-unnamed-io)
    * [pipes-unnamed](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/shared-memory/synchronous/pipes-unnamed)
    * [pipes-unnamed-io](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/shared-memory/synchronous/pipes-unnamed-io)
      **<- YOU ARE HERE**
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

Table of Contents

* [PIPES OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/shared-memory/synchronous/pipes-unnamed-io#pipes-overview)
* [CODE](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/shared-memory/synchronous/pipes-unnamed-io#code)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/shared-memory/synchronous/pipes-unnamed-io#run)
* [IN-PROCESS AND INTER-PROCESS COMMUNICATION OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/shared-memory/synchronous/pipes-unnamed-io#in-process-and-inter-process-communication-overview)

Documentation and Reference

* Refer to the
  [io.Pipe](https://pkg.go.dev/io#Pipe)
  package for more info
* Refer to
  [my cheat sheets](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/communication/in-process-and-inter-process-communications-ipc-overview-cheat-sheet)
  for a communications overview

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

Send data (write) to the pipe using io.Writer.

```go
fmt.Fprint(pw, "I am some data using pipe\n")
```

This writes the data to memory that is shared.

Now when ready, receive data (read) from the pipe (i.e. read from memory)
using io.Reader,

```go
_, err := io.Copy(os.Stdout, pr)
```

## RUN

Run,

```bash
go run main.go
```

## IN-PROCESS AND INTER-PROCESS COMMUNICATION OVERVIEW

Refer to
[my cheat sheets](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/communication/in-process-and-inter-process-communications-ipc-overview-cheat-sheet)
for a more thorough communications overview.

![IMAGE - in-process-and-inter-process-communication.jpg - IMAGE](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/docs/pics/in-process-and-inter-process-communication.jpg)
