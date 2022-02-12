# simple-os-interactions

 _A few os interactions using the `syscall` package._

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/interact-os/simple-os-interactions#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/interact-os/simple-os-interactions#run)

Documentation and references,

* Refer to the
  [syscall](https://pkg.go.dev/syscall)
  package for more info
* This repos [github webpage](https://jeffdecola.github.io/my-go-examples/)

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
go run simple-os-interactions.go
```
