# SIMPLE OS INTERACTIONS EXAMPLE

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_A few os interactions using the
[syscall](https://pkg.go.dev/syscall)
package._

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/go-runtime/interact-host-os/simple-os-interactions#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/go-runtime/interact-host-os/simple-os-interactions#run)

Documentation and Reference

* [syscall](https://pkg.go.dev/syscall)
  standard package

## OVERVIEW

Package syscall contains an interface to the low-level operating system primitives.

For example,

```go
// GET PROCESS ID (pid)
pid := syscall.Getpid()
fmt.Printf("Process ID (pid): %d\n", pid)

// GET THREAD ID (tid)
tid := syscall.Gettid()
fmt.Printf("Thread ID (id):   %d\n", tid)
```

## RUN

Run,

```bash
go run main.go
```
