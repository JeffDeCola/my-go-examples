# simple-go-runtime-interactions

_A few go runtime interactions using the `runtime` package._

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/interact-go-runtime/simple-go-runtime-interactions#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/interact-go-runtime/simple-go-runtime-interactions#run)

Documentation and references,

* Refer to the
  [runtime](https://pkg.go.dev/runtime)
  package for more info
* This repos [github webpage](https://jeffdecola.github.io/my-go-examples/)

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
go run simple-go-runtime-interactions.go
```
