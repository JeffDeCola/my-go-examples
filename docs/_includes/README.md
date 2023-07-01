  _built with
  [concourse](https://github.com/JeffDeCola/my-go-examples/blob/master/ci-README.md)_

# ARCHITECTURES

* BLOCKCHAIN

  * [bitcoin-ledger](https://github.com/JeffDeCola/my-go-examples/tree/master/architectures/blockchain/bitcoin-ledger)

    _Demonstrates a bitcoin ledger in a blockchain using the
    unspent transaction output model._

  * [create-bitcoin-address-from-ecdsa-publickey]()

    _Create a bitcoin address from your ecdsa public key
    using the `crypto/ecdsa` standard package._
  
  * [single-node-blockchain-with-REST]()

    _A simple single node sha256 blockchain with a REST JSON API
    (to view (GET) the blockchain and add (POST) a block)._

* SCRAPING

  * _Coming soon._

## CGO

* C CODE

  * [simple-c-code](https://github.com/JeffDeCola/my-go-examples/tree/master/cgo/c-code/simple-c-code)

    _A very simple example to show you how to write a c function in go._

  * [simple-c-code-using-stdio](https://github.com/JeffDeCola/my-go-examples/tree/master/cgo/c-code/simple-c-code-using-stdio)

    _A c function in go using stdio.h for printf._

## CLIENT/SERVER

* gRPC

  * _Coming soon._

* HTTP

  * _Coming soon._

* TCP/IP

  * _Coming soon._

## CLOUD SERVICE PROVIDERS

* GOOGLE CLOUD PLATFORM

  * _Coming soon._

* MICROSOFT AZURE

  * _Coming soon._

## COMMON GO

* ERROR REPORTING

  * [error-example](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/error-reporting/error-example)

    _Error Handling using the standard `error` package._

* FLAGS

  * [flags](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/flags/flags)

    _The standard `flag` package makes it easy to implement command-line flag parsing._

* LOGGING

  * [logrus](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/logging/logrus)

    _Logging using `logrus` package._

* MY GENERIC GO TEMPLATE

  * [jeffs-basic-go-template](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/my-generic-go-template/jeffs-basic-go-template)

    _A simple go template with flags, logging & error handling._

* TESTING

  * _Coming soon._

## CRYPTOGRAPHY

* ASYMMETRIC CRYPTOGRAPHY
  _(Great for digital signatures (verify sender) and receiving encrypted data)_

  * _Coming soon._

* HASHING
  _(Great for getting fingerprints)_

  * _Coming soon._

* SYMMETRIC CRYPTOGRAPHY
  _(Using the same key to encrypt and decrypt)_

  * _Coming soon._

## DATABASES

* NON-RELATIONAL

  * _Coming soon._

* RELATIONAL

  * _Coming soon._

## FUNCTIONS, METHODS AND INTERFACES

* FUNCTIONS

  * [functions](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/functions/functions)

    _Using functions to calculate the area of a rectangle and circle._

  * [functions-pointers-arguments](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/functions/functions-pointers-arguments)

    _Using functions to calculate the area of a rectangle and circle
    by passing pointers._

* METHODS

  * [methods](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/methods/methods)

    _Using methods to calculate the area of a rectangle and circle._

  * [methods-pointers-arguments](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/methods/methods-pointers-arguments)

    _Using methods to calculate the area of a rectangle and circle
    by passing pointers._

  * [methods-pointers-receivers](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/methods/methods-pointers-receivers)

    _Using methods to calculate the area of a rectangle and circle
    by passing pointers and using pointer receivers._

* INTERFACES

  * [interfaces](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/interfaces/interfaces)

    _Using an interface to calculate the area of a rectangle and circle._

  * [interfaces-pointers-arguments](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/interfaces/interfaces-pointers-arguments)

    _Using an interface to calculate the area of a rectangle and circle
    by passing pointers._

  * [interfaces-pointer-receivers](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/interfaces/interfaces-pointers-receivers)

    _Using an interface to calculate the area of a rectangle and circle
    by passing pointers and using pointer receivers._

  * [shapes-package](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/interfaces/shapes-package)

    _Using an interface to calculate and manipulate simple 2D and 3D geometric shapes._

## GO RUNTIME

* GO ROUTINES

  * [goroutines-multi-core](https://github.com/JeffDeCola/my-go-examples/tree/master/go-runtime/goroutines/goroutines-multi-core)

    _Concurrency across multiples cores.  You can play around with workers,
    threads, cpus/cores and nice to find the fastest performance.
    It will find the total number of prime numbers within a range._

  * [goroutines-waitgroup](https://github.com/JeffDeCola/my-go-examples/tree/master/go-runtime/goroutines/goroutines-waitgroup)

    _Concurrency using a waitgroup (waiting for a collection of goroutines to finish)._

  * [goroutines-worker-pools](https://github.com/JeffDeCola/my-go-examples/tree/master/go-runtime/goroutines/goroutines-worker-pools)

    _Concurrency using worker pools._

* INTERACT GO RUNTIME

  * [simple-go-runtime-interactions](https://github.com/JeffDeCola/my-go-examples/tree/master/go-runtime/interact-host-os/simple-external-commands)

    _A few go runtime interactions using the `runtime` package._

* INTERACT HOST OS

  * [simple-external-commands](https://github.com/JeffDeCola/my-go-examples/tree/master/go-runtime/interact-host-os/simple-external-commands)

    _Run a few os commands using the `exec` package._

  * [simple-os-interactions](https://github.com/JeffDeCola/my-go-examples/tree/master/go-runtime/interact-host-os/simple-os-interactions)

    _A few os interactions using the `syscall` package._

## INPUT/OUTPUT

* IO READER

  * [io-reader](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader/io-reader)

    _Read data (a stream of bytes) from a string, buffer, file, stdin and
    a pipe to a buffer using the standard `io` package._

  * [io-reader-simple](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader/io-reader-simple)

    _Read data (a stream of bytes) from a buffer to a buffer
    using the standard `io` package._

  * [read-file](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader/read-file)

    _Read a file (*os.File) to a buffer._

  * [read-user-input](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader/read-user-input)

    _Read user input (os.Stdin) to a buffer (using Read method)
    and string (using Fscan)._

* IO WRITER

  * [io-writer](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-writer/io-writer)

    _Write data (a stream of bytes) to a buffer, file, stdout and a pipe
    from a buffer using the standard `io` package._

  * [io-writer-simple](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-writer/io-writer-simple)

    _Write data (a stream of bytes) to a buffer
    from a buffer using the standard `io` package._

## IN-PROCESS COMMUNICATION

* SHARED MEMORY

  * ASYNCHRONOUS

    * _Coming soon._

  * SYNCHRONOUS

    * [pipes-unnamed-simple](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/shared-memory/synchronous/pipes-unnamed-simple)

      _A pipe provides a uni-directional communication channel.
      This is a very simple example of an unnamed pipe._

    * [pipes-unnamed](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/shared-memory/synchronous/pipes-unnamed)

      _This is a more robust example of an unnamed pipe showing multiple reads._

    * [pipes-unnamed-io](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/shared-memory/synchronous/pipes-unnamed-io)

      _This example of an unnamed pipe connects an io.Writer and io.Reader._

* MESSAGE PASSING

  * ASYNCHRONOUS

    * [channels-buffered](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/message-passing/asynchronous/channels-buffered)

      _Buffered channels are uni-directional, asynchronous with no blocking._

  * SYNCHRONOUS

    * [channels-unbuffered](https://github.com/JeffDeCola/my-go-examples/tree/master/in-process-communication/message-passing/synchronous/channels-unbuffered)

      _Unbuffered channels are uni-directional, synchronous with blocking._

## INTER-PROCESS COMMUNICATION (IPC)

* SHARED MEMORY

  * ASYNCHRONOUS

    * _Coming soon._

  * SYNCHRONOUS

    * [pipes-named](https://github.com/JeffDeCola/my-go-examples/tree/master/inter-process-communication-ipc/shared-memory/synchronous/pipes-named)

      _Sending data over a named-pipe (FIFO) from one process to another process._

* MESSAGE PASSING

  * OPERATING SYSTEM

    * ASYNCHRONOUS

      * [message-queues](https://github.com/JeffDeCola/my-go-examples/tree/master/inter-process-communication-ipc/message-passing/operating-system/asynchronous/message-queues)

        _tbd._

      * [pub-sub-nats-os](https://github.com/JeffDeCola/my-go-examples/tree/master/inter-process-communication-ipc/message-passing/operating-system/asynchronous/pub-sub-nats-os)

        _tbd._

      * [tcp](https://github.com/JeffDeCola/my-go-examples/tree/master/inter-process-communication-ipc/message-passing/operating-system/asynchronous/tcp)

        _tbd._

    * SYNCHRONOUS

      * [grpc-os](https://github.com/JeffDeCola/my-go-examples/tree/master/inter-process-communication-ipc/message-passing/operating-system/synchronous/grpc-os)

        _tbd._

  * NETWORK

    * ASYNCHRONOUS

      * [pub-sub-nats-network](https://github.com/JeffDeCola/my-go-examples/tree/master/inter-process-communication-ipc/message-passing/network/asynchronous/pub-sub-nats-network)

        _tbd._

      * [tcp-ip](https://github.com/JeffDeCola/my-go-examples/tree/master/inter-process-communication-ipc/message-passing/network/asynchronous/tcp-ip)

        _tbd._

    * SYNCHRONOUS

      * [grpc-network](https://github.com/JeffDeCola/my-go-examples/tree/master/inter-process-communication-ipc/message-passing/network/synchronous/grpc-network)

        _tbd._

      * [rest](https://github.com/JeffDeCola/my-go-examples/tree/master/inter-process-communication-ipc/message-passing/network/synchronous/rest)

        _tbd._

## IoT

* RASPBERRY PI

  * [blink-led-raspberry-pi-gpio-periph](https://github.com/JeffDeCola/my-go-examples/tree/master/iot/raspberry-pi/blink-led-raspberry-pi-gpio-periph)

    _GPIO OUTPUT -
    Blink an LED
    via a Raspberry Pi GPIO
    using the `periph.io/...` packages._

  * [push-button-raspberry-pi-gpio-periph](https://github.com/JeffDeCola/my-go-examples/tree/master/iot/raspberry-pi/push-button-raspberry-pi-gpio-periph)

    _GPIO INPUT -
    Push a button
    via a Raspberry Pi GPIO
    using the `periph.io/...` packages._

  * [toggle-led-with-button-raspberry-pi-gpio-periph](https://github.com/JeffDeCola/my-go-examples/tree/master/iot/raspberry-pi/toggle-led-with-button-raspberry-pi-gpio-periph)

    _Toggle an LED with a button push
    via a Raspberry Pi GPIO
    using the `periph.io/...` packages._

## MODULES AND PACKAGES

* LOCAL PACKAGES

  * [module-with-local-package](https://github.com/JeffDeCola/my-go-examples/tree/master/modules-and-packages/local-packages/module-with-local-package)

    _A go module with a local package._

* REMOTE PACKAGES
  * [module-with-remote-package](https://github.com/JeffDeCola/my-go-examples/tree/master/modules-and-packages/remote-packages/module-with-remote-package)

    _A go module with a remote (public) package._
