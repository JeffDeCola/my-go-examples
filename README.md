# my-go-examples

[![Go Report Card](https://goreportcard.com/badge/github.com/JeffDeCola/my-go-examples)](https://goreportcard.com/report/github.com/JeffDeCola/my-go-examples)
[![GoDoc](https://godoc.org/github.com/JeffDeCola/my-go-examples?status.svg)](https://godoc.org/github.com/JeffDeCola/my-go-examples)
[![Maintainability](https://api.codeclimate.com/v1/badges/3c5477c63d77a071fdff/maintainability)](https://codeclimate.com/github/JeffDeCola/my-go-examples/maintainability)
[![Issue Count](https://codeclimate.com/github/JeffDeCola/my-go-examples/badges/issue_count.svg)](https://codeclimate.com/github/JeffDeCola/my-go-examples/issues)
[![License](http://img.shields.io/:license-mit-blue.svg)](http://jeffdecola.mit-license.org)

_A place to keep my go snippets and examples._

Table of Contents,

* [ARCHITECTURES](https://github.com/JeffDeCola/my-go-examples#architectures)
  * BLOCKCHAIN
  * SCRAPING
* [CGO](https://github.com/JeffDeCola/my-go-examples#cgo)
* [CLIENT/SERVER](https://github.com/JeffDeCola/my-go-examples#clientserver)
  * HTTP
  * gRPC
  * TCP/IP
* [CLOUD SERVICE PROVIDERS](https://github.com/JeffDeCola/my-go-examples#cloud-service-providers)
* [COMMON GO](https://github.com/JeffDeCola/my-go-examples#common-go)
  * ERROR REPORTING
  * FLAGS
  * LOGGING
  * MY GENERIC GO TEMPLATE
  * TESTING
* [COMMUNICATION](https://github.com/JeffDeCola/my-go-examples#communication)
  * API
  * INTER PROCESS COMMUNICATION (IPC)
* [CRYPTOGRAPHY](https://github.com/JeffDeCola/my-go-examples#cryptography)
  * ASYMMETRIC CRYPTOGRAPHY
  * HASHING
  * SYMMETRIC CRYPTOGRAPHY
* [DATABASES](https://github.com/JeffDeCola/my-go-examples#databases)
* [FUNCTIONS, METHODS AND INTERFACES](https://github.com/JeffDeCola/my-go-examples#functions-methods-and-interfaces)
  * FUNCTIONS
  * METHODS
  * INTERFACES
* [GOROUTINES](https://github.com/JeffDeCola/my-go-examples#goroutines)
* [INPUT/OUTPUT](https://github.com/JeffDeCola/my-go-examples#inputoutput)
  * IO READER
  * IO WRITER
* [IoT](https://github.com/JeffDeCola/my-go-examples#iot)
  * RASPBERRY PI

Documentation and reference,

* [go-cheat-sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet)
* [my-go-packages](https://github.com/JeffDeCola/my-go-packages)
* [my-go-tools](https://github.com/JeffDeCola/my-go-tools)
* This repos
  [github webpage](https://jeffdecola.github.io/my-go-examples/)
  _built with
  [concourse](https://github.com/JeffDeCola/my-go-examples/blob/master/ci-README.md)_

## GO EXAMPLES

_All sections in alphabetical order._

### ARCHITECTURES
  
_Using go to build fun things._

* BLOCKCHAIN
  
* SCRAPING

  _coming soon_

### CGO

_Using c with go._

### CLIENT/SERVER

_Lets make a server._

* HTTP

* gRPC

* TCP/IP

### CLOUD SERVICE PROVIDERS

_AWS and Google services._

### COMMON GO

_A good place to start._

* ERROR REPORTING

  * [error-example](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/error-reporting/error-example)

    _Error Handling using the standard `error` package._

* FLAGS

  * [flag](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/flags/flag)

    _The standard `flag` package makes it easy to implement command-line flag parsing._

* LOGGING

  * [logrus](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/logging/logrus)

    _Logging using `logrus` package._

* MY GENERIC GO TEMPLATE

  * [jeffs-basic-go-template](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/my-generic-go-template/jeffs-basic-go-template)

    _A simple go template with flags, logging & error handling._

* TESTING

### COMMUNICATION

_Talk to me._

* API

  * ASYNCHRONOUS
  
    * MESSAGE QUEUES
  
    * TCP/IP
  
  * SYNCHRONOUS
  
    * gRPC
  
    * REST

* INTER PROCESS COMMUNICATION (IPC)

  * ASYNCHRONOUS

    * CHANNELS (UNBUFFERED)

    * SOCKETS

  * SYNCHRONOUS
  
    * CHANNELS (BUFFERED)
  
    * PIPES

### CRYPTOGRAPHY

_Secure communication techniques._

* ASYMMETRIC CRYPTOGRAPHY

  _Great for digital signatures (verify sender) and receiving encrypted data._

* HASHING

  _Great for getting fingerprints._

* SYMMETRIC CRYPTOGRAPHY

  _Using the same key to encrypt and decrypt._

### DATABASES

_Save and organize your data._

### FUNCTIONS, METHODS AND INTERFACES

_The following code demonstrates functions, methods and interfaces using
simple area and perimeter calculations for different shapes._

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

  * [geometry-package](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/interfaces/geometry-package)

    _Using an interface to calculate the area and perimeter of a rectangle,
    circle and triangle via a geometry package.
    It uses pointers for arguments and receivers._

### GOROUTINES

_Go is written for concurrency. The go runtime schedules goroutines on threads.
The OS schedules these threads on cpus/cores._

### INPUT/OUTPUT

_Using the
[io](https://pkg.go.dev/io)
package to read data and write data to/from various IO
(files, buffers, pipes)._

* IO READER

  * [io-reader](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader/io-reader)

    _Read data (a stream of bytes) from a string, buffer, file, stdin and
    a pipe to a buffer using the standard `io` package._

  * [io-reader-simple](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader/io-reader-simple)

    _Read data (a stream of bytes) from a buffer to a buffer
    using the standard `io` package._

* IO WRITER

  * [io-writer](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-writer/io-writer)

    _Write data (a stream of bytes) to a buffer, file, stdout and a pipe
    from a buffer using the standard `io` package._

  * [io-writer-simple](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-writer/io-writer-simple)

    _Write data (a stream of bytes) to a buffer
    from a buffer using the standard `io` package._

### IoT

_Using go with the Internet of Things (IoT)._

* RASPBERRY PI
