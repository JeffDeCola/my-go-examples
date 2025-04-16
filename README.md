# MY GO EXAMPLES

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)
[![Go Reference](https://pkg.go.dev/badge/github.com/JeffDeCola/my-go-examples.svg)](https://pkg.go.dev/github.com/JeffDeCola/my-go-examples)
[![Go Report Card](https://goreportcard.com/badge/github.com/JeffDeCola/my-go-examples)](https://goreportcard.com/report/github.com/JeffDeCola/my-go-examples)

_A place to keep my go examples._

Table of Contents

* [ARCHITECTURES](https://github.com/JeffDeCola/my-go-examples#architectures)
  * BLOCKCHAIN
  * SCRAPING
* [CGO](https://github.com/JeffDeCola/my-go-examples#cgo)
  * C CODE
* [CLIENT/SERVER](https://github.com/JeffDeCola/my-go-examples#clientserver)
  * gRPC
  * HTTP
  * TCP/IP
* [CLOUD SERVICE PROVIDERS](https://github.com/JeffDeCola/my-go-examples#cloud-service-providers)
  * GOOGLE CLOUD PLATFORM
  * MICROSOFT AZURE
* [COMMON GO](https://github.com/JeffDeCola/my-go-examples#common-go)
  * ERROR REPORTING
  * FLAGS
  * LOGGING
  * MY GENERIC GO TEMPLATE
  * TESTING
* [CRYPTOGRAPHY](https://github.com/JeffDeCola/my-go-examples#cryptography)
  * ASYMMETRIC CRYPTOGRAPHY
  * HASHING
  * SYMMETRIC CRYPTOGRAPHY
* [DATABASES](https://github.com/JeffDeCola/my-go-examples#databases)
  * NON-RELATIONAL
  * RELATIONAL
* [FUNCTIONS, METHODS AND INTERFACES](https://github.com/JeffDeCola/my-go-examples#functions-methods-and-interfaces)
  * FUNCTIONS
  * METHODS
  * INTERFACES
* [GO RUNTIME](https://github.com/JeffDeCola/my-go-examples#go-runtime)
  * GOROUTINES
  * INTERACT GO RUNTIME
  * INTERACT HOST OS
* [INPUT/OUTPUT](https://github.com/JeffDeCola/my-go-examples#inputoutput)
  * IO READER
  * IO WRITER
* [IN-PROCESS COMMUNICATION](https://github.com/JeffDeCola/my-go-examples#in-process-communication)
  * SHARED MEMORY
  * MESSAGE PASSING
* [INTER-PROCESS COMMUNICATION (IPC)](https://github.com/JeffDeCola/my-go-examples#inter-process-communication-ipc)
  * SHARED MEMORY
  * MESSAGE PASSING
* [IoT](https://github.com/JeffDeCola/my-go-examples#iot)
  * RASPBERRY PI
* [MODULES AND PACKAGES](https://github.com/JeffDeCola/my-go-examples#modules-and-packages)
  * LOCAL PACKAGES
  * REMOTE PACKAGES
* [STANDARD GO PACKAGES](https://github.com/JeffDeCola/my-go-examples#standard-go-packages)

Documentation and Reference

* [go](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet)
* [go standard library](https://pkg.go.dev/std)
* [my-go-packages](https://github.com/JeffDeCola/my-go-packages)
* [my-go-tools](https://github.com/JeffDeCola/my-go-tools)
* This repos
  [github webpage](https://jeffdecola.github.io/my-go-examples/)
  _built with
  [concourse](https://github.com/JeffDeCola/my-go-examples/blob/master/ci-README.md)_

## ARCHITECTURES

* BLOCKCHAIN

  * [bitcoin-ledger](https://github.com/JeffDeCola/my-go-examples/tree/master/architectures/blockchain/bitcoin-ledger)

    _Demonstrates a bitcoin ledger in a blockchain using the
    unspent transaction output model._

  * [create-bitcoin-address-from-ecdsa-publickey](https://github.com/JeffDeCola/my-go-examples/tree/master/architectures/blockchain/create-bitcoin-address-from-ecdsa-publickey)

    _Create a bitcoin address from your ecdsa public key
    using the
    [crypto/ecdsa](https://pkg.go.dev/crypto/ecdsa)
    standard package._

  * [single-node-blockchain-with-REST](https://github.com/JeffDeCola/my-go-examples/tree/master/architectures/blockchain/single-node-blockchain-with-REST)

    _A simple single node sha256 blockchain with a REST JSON API
    (to view (GET) the blockchain and add (POST) a block)._

* WEB SCRAPING

  * _[web scrape twitter for ???]_

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

  _Error Handling using the standard
  [errors](https://pkg.go.dev/errors)
  standard package._

* FLAGS

  * [flags](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/flags/flags)

    _The
    [flag](https://pkg.go.dev/flag)
    standard package makes it easy to implement command-line flag parsing._

* LOGGING

  * [jeffs-logger](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/logging/jeffs-logger)

    _Logging using my
    [github.com/JeffDeCola/my-go-packages/golang/logger](https://github.com/JeffDeCola/my-go-packages/tree/master/golang/logger)
    which uses the structured logging
    [log/slog](https://pkg.go.dev/log/slog)
    standard package._

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
    It will find the total number of prime numbers within a range.
    Lightweight goroutines are amazing._

  * [goroutines-waitgroup](https://github.com/JeffDeCola/my-go-examples/tree/master/go-runtime/goroutines/goroutines-waitgroup)

    _Concurrency using a waitgroup (waiting for a collection of goroutines to finish)._

  * [goroutines-worker-pools](https://github.com/JeffDeCola/my-go-examples/tree/master/go-runtime/goroutines/goroutines-worker-pools)

    _Concurrency using worker pools._

* INTERACT GO RUNTIME

  * [simple-go-runtime-interactions](https://github.com/JeffDeCola/my-go-examples/tree/master/go-runtime/interact-go-runtime/simple-go-runtime-interactions)

    _A few go runtime interactions using the
    [runtime](https://pkg.go.dev/runtime)
    standard package._

* INTERACT HOST OS

  * [simple-external-commands](https://github.com/JeffDeCola/my-go-examples/tree/master/go-runtime/interact-host-os/simple-external-commands)

    _Run a few os commands using the
    [os/exec](https://pkg.go.dev/os/exec)
    standard package._

  * [simple-os-interactions](https://github.com/JeffDeCola/my-go-examples/tree/master/go-runtime/interact-host-os/simple-os-interactions)

    _A few os interactions using the
    [syscall](https://pkg.go.dev/syscall)
    package._

## INPUT/OUTPUT

* IO READER

  * [io-reader](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader/io-reader)

    _Read data (a stream of bytes) from a string, buffer, file, stdin and
    a pipe to a buffer using the
    [io](https://pkg.go.dev/io)
    standard package._

  * [io-reader-simple](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader/io-reader-simple)

    _Read data (a stream of bytes) from a buffer to a buffer
    using the
    [io](https://pkg.go.dev/io)
    standard package._

  * [read-file](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader/read-file)

    _Read a file (*os.File) to a buffer._

  * [read-user-input](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader/read-user-input)

    _Read user input (os.Stdin) to a buffer (using Read method)
    and string (using Fscan)._

* IO WRITER

  * [io-writer](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-writer/io-writer)

    _Write data (a stream of bytes) to a buffer, file, stdout and a pipe
    from a buffer using the
    [io](https://pkg.go.dev/io)
    standard package._

  * [io-writer-simple](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-writer/io-writer-simple)

    _Write data (a stream of bytes) to a buffer
    from a buffer using the
    [io](https://pkg.go.dev/io)
    standard package._

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
    Blink an LED via a Raspberry Pi GPIO using the
    [periph.io/...](https://pkg.go.dev/periph.io/x/conn/v3)
    packages._

  * [push-button-raspberry-pi-gpio-periph](https://github.com/JeffDeCola/my-go-examples/tree/master/iot/raspberry-pi/push-button-raspberry-pi-gpio-periph)

    _GPIO INPUT -
    Push a button via a Raspberry Pi GPIO using the
    [periph.io/...](https://pkg.go.dev/periph.io/x/conn/v3)
    packages._

  * [toggle-led-with-button-raspberry-pi-gpio-periph](https://github.com/JeffDeCola/my-go-examples/tree/master/iot/raspberry-pi/toggle-led-with-button-raspberry-pi-gpio-periph)

    _Toggle an LED with a button push via a Raspberry Pi GPIO using the
    [periph.io/...](https://pkg.go.dev/periph.io/x/conn/v3)
    packages._

## MODULES AND PACKAGES

* LOCAL PACKAGES

  * [module-with-local-package](https://github.com/JeffDeCola/my-go-examples/tree/master/modules-and-packages/local-packages/module-with-local-package)

    _A go module with a local package._

* REMOTE PACKAGES
  * [module-with-remote-package](https://github.com/JeffDeCola/my-go-examples/tree/master/modules-and-packages/remote-packages/module-with-remote-package)

    _A go module with a remote (public) package._

## STANDARD GO PACKAGES

* archive
  * tar
  * zip
* bufio
* bytes
* cmp
* compress
  * bzip2
  * flate
  * gzip
  * lzw
  * zlib
* container
  * heap
  * list
  * ring
* context
* crypto
  * aes -
    [encryptfile](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/encryptfile)
    tool,
    [decryptfile](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/decryptfile)
    tool
  * ecdsa -
    [create-bitcoin-address-from-ecdsa-publickey](https://github.com/JeffDeCola/my-go-examples/tree/master/architectures/blockchain/create-bitcoin-address-from-ecdsa-publickey)
  * elliptic -
    [create-bitcoin-address-from-ecdsa-publickey](https://github.com/JeffDeCola/my-go-examples/tree/master/architectures/blockchain/create-bitcoin-address-from-ecdsa-publickey)
  * md5 -
    [md5-hash-file](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/md5-hash-file)
    tool
  * rand -
    [create-bitcoin-address-from-ecdsa-publickey](https://github.com/JeffDeCola/my-go-examples/tree/master/architectures/blockchain/create-bitcoin-address-from-ecdsa-publickey)
  * sha256 -
    [create-bitcoin-address-from-ecdsa-publickey](https://github.com/JeffDeCola/my-go-examples/tree/master/architectures/blockchain/create-bitcoin-address-from-ecdsa-publickey),
    [sha256-hash-file](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/sha256-hash-file)
    tool
  * x509 -
    [create-bitcoin-address-from-ecdsa-publickey](https://github.com/JeffDeCola/my-go-examples/tree/master/architectures/blockchain/create-bitcoin-address-from-ecdsa-publickey)
  * etc...
* database
  * sql
  * sql-driver
* debug
  * buildinfo
  * dwarf
  * etc...
* embed
* encoding
  * hex -
    [create-bitcoin-address-from-ecdsa-publickey](https://github.com/JeffDeCola/my-go-examples/tree/master/architectures/blockchain/create-bitcoin-address-from-ecdsa-publickey)
  * pem -
    [create-bitcoin-address-from-ecdsa-publickey](https://github.com/JeffDeCola/my-go-examples/tree/master/architectures/blockchain/create-bitcoin-address-from-ecdsa-publickey)
  * etc...
* errors -
  [error-example](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/error-reporting/error-example)
* expvar
* flag -
  [flags](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/flags/flags)
* fmt
* go
  * ast
  * build
  * etc...
* hash
  * adler32
  * crc32
  * crc64
  * fnv
  * maphash
* html
  * template
* image
  * color
  * color/palette
  * draw
  * gif
  * jpeg
  * png
* index
  * suffixarray
* io -
  [io-reader](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader/io-reader),
  [io-reader-simple](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader/io-reader-simple),
  [io-writer](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-writer/io-writer),
  [io-writer-simple](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-writer/io-writer-simple)
  * fs
  * ioutil
* iter
* log
  * slog -
    [jeffs-logger](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/logging/jeffs-logger)
  * syslog
* maps
* math
  * big
  * bits
  * cmplx
  * rand
  * rand/v2
* mime
  * multipart
  * quotedprintable
* net
  * http
  * http/cgi
  * etc.,,
* os
  * exec -
    [simple-external-commands](https://github.com/JeffDeCola/my-go-examples/tree/master/go-runtime/interact-host-os/simple-external-commands)
  * signal
  * user
* path
  * filepath
* reflect -
  [create-bitcoin-address-from-ecdsa-publickey](https://github.com/JeffDeCola/my-go-examples/tree/master/architectures/blockchain/create-bitcoin-address-from-ecdsa-publickey)
* regexp (syntax)
* runtime -
  [simple-go-runtime-interactions](https://github.com/JeffDeCola/my-go-examples/tree/master/go-runtime/interact-go-runtime/simple-go-runtime-interactions)
  * cgo
  * coverage
  * etc...
* slices
* sort
* strconv
* strings
* structs
* sync (atomic)
* syscall -
  [simple-os-interactions](https://github.com/JeffDeCola/my-go-examples/tree/master/go-runtime/interact-host-os/simple-os-interactions)
  * js
* testing
  * fstest
  * iotest
  * quick
  * slogtest
  * synctest
* text
  * scanner
  * tabwriter
  * template
  * template/parse
* time
  * tzdata
* unicode
  * utf1
  * utf8
* unique
* unsafe
* weak
