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
* [BASIC SYNTAX](https://github.com/JeffDeCola/my-go-examples#basic-syntax)
  * DATA TYPES
  * FUNCTIONS
  * INTERFACES
  * METHODS
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

### BASIC SYNTAX

_Just your basic syntax examples in go._

* DATA TYPES

* FUNCTIONS

* INTERFACES

* METHODS

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

  * [errors](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/error-reporting/errors)

    _Error Handling using the standard `error` package._

* FLAGS

  * [flag](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/flags/flag)

    _The standard `flag` package makes it easy to implement command-line flag parsing._

* LOGGING

  * [logrus](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/logging/logrus)

    _Logging using non-standard `logrus` package._

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

### GOROUTINES

_Go is written for concurrency. The go runtime schedules goroutines on threads.
The OS schedules these threads on cpus/cores._

### INPUT/OUTPUT

_Using the
[io](https://pkg.go.dev/io)
package to read data and write data to/from various IO
(files, buffers, pipes)._

* IO READER

* IO WRITER

### IoT

_Using go with the Internet of Things (IoT)._

* RASPBERRY PI
