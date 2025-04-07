# SIMPLE GO RUNTIME INTERACTIONS EXAMPLE

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_A few go runtime interactions using the
[runtime](https://pkg.go.dev/runtime)
standard package._

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/go-runtime/interact-go-runtime/simple-go-runtime-interactions#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/go-runtime/interact-go-runtime/simple-go-runtime-interactions#run)

Documentation and Reference

* [runtime](https://pkg.go.dev/runtime)
  standard package

## OVERVIEW

Package runtime contains operations that interact with Go's runtime system,
such as functions to control goroutines.

For example,

```go
// NUMBER CPU
numCPUs := runtime.NumCPU()
fmt.Printf("Number of CPUS: %d\n", numCPUs)

// NUMBER GOROUTINES
numGoRoutines := runtime.NumGoroutine()
fmt.Printf("Number of Go Routines that exist: %d\n", numGoRoutines)
```

## RUN

Run,

```bash
go run main.go
```
