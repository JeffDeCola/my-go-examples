# goroutines-multi-core

_Concurrency across multi-cores. You can play around with workers,
threads, cpus/cores and nice to find the fastest performance.
It will find the total number of prime numbers within a range._

This will show that **lightweight goroutines are amazing**.

Other examples using,

* [goroutines-multi-core](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines/goroutines-multi-core)
  **<- You are here**
* [goroutines-waitgroup](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines/goroutines-waitgroup)
* [goroutines-worker-pools](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines/goroutines-worker-pools)

Table of contents,

* [MACOS DOES NOT WORK](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines/goroutines-multi-core#macos-does-not-work)
* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines/goroutines-multi-core#overview)
* [SETUP](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines/goroutines-multi-core#setup)
* [GO RUNTIME FEATURES](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines/goroutines-multi-core#go-runtime-features)
  * [FEATURE 1 -  LOCK A GOROUTINE TO A THREAD](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines/goroutines-multi-core#feature-1----lock-a-goroutine-to-a-thread)
* [OS KERNEL FEATURES](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines/goroutines-multi-core#os-kernel-features)
  * [FEATURE 2 - LOCK A GOROUTINE TO A CPU](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines/goroutines-multi-core#feature-2---lock-a-goroutine-to-a-cpu)
  * [FEATURE 3 - LOCK A THREAD TO A CPU](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines/goroutines-multi-core#feature-3---lock-a-thread-to-a-cpu)
  * [FEATURE 4 - SET PRIORITY ON THREAD](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines/goroutines-multi-core#feature-4---set-priority-on-thread)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines/goroutines-multi-core#run)
* [SOME BENCHMARKS](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines/goroutines-multi-core#some-benchmarks)

Documentation and references,

* [simple-go-runtime-interactions](https://github.com/JeffDeCola/my-go-examples/tree/master/interact-go-runtime/simple-go-runtime-interactions)
show a few go runtime interactions using the `runtime` package
* [simple-os-interactions](https://github.com/JeffDeCola/my-go-examples/tree/master/interact-os/simple-os-interactions)
shows a few os interactions using the `syscall` package
* This repos [github webpage](https://jeffdecola.github.io/my-go-examples/)

## MACOS DOES NOT WORK

I could not get the c code to run on macOS and did not spend the time to
figure it out. But if someone figures this out, please let me know.

## OVERVIEW

Your go executable has a go runtime environment that schedules
where the goroutines run (which CPU/CORE and which Thread).
This is constantly changing.

This example will show you how to `lock the goroutine` to a particular
CPU/CORE and a particular Thread.

This example will also allow you to change variables (like how many workers) to
see if you can improve performance.

These illustrations may help,

![IMAGE - goroutines-multi-core - IMAGE](../../docs/pics/goroutines/goroutines-multi-core.jpg)

![IMAGE - goroutines-lock-threads-cores - IMAGE](../../docs/pics/goroutines/goroutines-lock-threads-cores.jpg)

## SETUP

This program will allow you to set the numbers of workers and check the
statistics of each goroutine and core.

```go
// WORKERS
const useGoroutine = false
const numberWorkers = 50

// PRIME NUMBER
const testForPrimes = 200000

// BUFFER CHANNEL
var channelBufferSize = numberWorkers + 1

// FEATURE 1 - LOCK A GOROUTINE TO A THREAD
const lockGoroutineToThread = false

// FEATURE 2 - LOCK A GOROUTINE TO A CPU
const lockGoroutineToCPU = false 
var useTheseCPUs = []int{9}

// FEATURE 3 - LOCK A THREAD TO A CPU
const lockThreadToCore = false

// FEATURE 4 - SET PRIORITY ON THREAD
const setPriority = false  
const setPriorityLevel = 0
```

The workload will simple split the workload (finding a prime number)
over the number of workers.  For example if you have
5 workers and you are trying to find out how many numbers are prime under 1000,
worker 1 will get numbers 1-200, worker 2, 201-400 and so on.

## GO RUNTIME FEATURES

The `go runtime` will schedule goroutines to cores and threads.  And this
can change a lot.  

### FEATURE 1 -  LOCK A GOROUTINE TO A THREAD

This feature will lock a goroutine (worker) to a thread.
The thread can bounce around to different CPUs,
carrying the goroutine with it.

```go
// FEATURE 1 - LOCK A GOROUTINE TO A THREAD
if lockGoroutineToThread {
    runtime.LockOSThread()
    defer runtime.UnlockOSThread()
}
```

## OS KERNEL FEATURES

These require c code or system calls to the kernel.

### FEATURE 2 - LOCK A GOROUTINE TO A CPU

This feature will lock a goroutine (worker) to a CPU.
The goroutine can bounce around to different threads
in the same CPU.

```go
// FEATURE 2 - LOCK A GOROUTINE TO A CPU
if lockGoroutineToCPU {
    C.set_affinity(C.int(useCPU))
}
```

### FEATURE 3 - LOCK A THREAD TO A CPU

This feature will lock a thread to a CPU.
Goroutines can go to other threads on other CPUs.

```go
// FEATURE 3 -  LOCK A THREAD TO A CPU
if lockThreadToCore {
    // Get the cpu your are on
    cpuID := C.sched_getcpu()
    // lock the thread to a cpu
    C.lock_thread(C.int(cpuID))
}
```

### FEATURE 4 - SET PRIORITY ON THREAD

I believe setting it works, but getting the value and printing it out is
always off a little. So take my nice code with a grain of salt.
I will have to revisit this at some point.

```go
// FEATURE 4 - SET PRIORITY ON THREAD
if setPriority {
    err := syscall.Setpriority(syscall.PRIO_PROCESS, 0, setPriorityLevel)
    if err != nil {
        println("Setpriority failed")
        return
    }
}
```

## RUN

```bash
go run goroutines-multi-core.go
go run goroutines-multi-core.go -loglevel info
```

Press return to exit.

Check a pid, nice levels and show threads,

```bash
ps -lTp <pid>
```

## SOME BENCHMARKS

OK, after all that work, here is the gravy.
For calculating all primes up to 200,000 (17,984 primes)

Running on my rig,

| FOCUS     | F1 LK-GR | F2-LK CPU | F3-LK T | F4-PRTY | Workers |   Time |           Comments |
|:----------|---------:|----------:|--------:|--------:|--------:|-------:|:-------------------|
| No GR     |        F |     F (1) |       F | default |      50 |  29.44 | The one Worker     |
|           |        F |     F (1) |       F | default |     500 |  29.42 | could use diff CPU |
|           |        F |     F (1) |       F | default |    2000 |  29.30 | and diff thread    |
|           |        F |     F (1) |       F | default |   20000 |  26.58 |                    |
| GR ON     |        F |    F (16) |       F | default |      50 |   4.12 | A Worker could use |
|           |        F |    F (16) |       F | default |     500 |   3.97 | diff CPU and and   |
|           |        F |    F (16) |       F | default |    2000 |   3.95 | diff Thread        |
|           |        F |    F (16) |       F | default |   20000 |   3.54 |                    |
|           |        F |    F (16) |       F | default |   50000 |**2.97**|                    |
|           |        F |    F (16) |       F | default |  100000 |   3.55 |                    |
| F1-LK GR  |        T |    F (16) |       F | default |      50 |   4.12 | A Worker will use  |
|           |        T |    F (16) |       F | default |     500 |   4.05 | same thread but    |
|           |        T |    F (16) |       F | default |    2000 |   3.93 | could use diff CPU |
|           |        T |    F (16) |       F | default |   20000 |   3.76 |                    |
| F2-1 CPU  |        F |       T-1 |       F | default |      50 |  29.47 | A Worker in same   |
|           |        F |       T-1 |       F | default |     500 |  29.40 | CPU but could use  |
|           |        F |       T-1 |       F | default |    2000 |  29.19 | diff thread        |
|           |        F |       T-1 |       F | default |   20000 |  26.89 |                    |
| F2-8 CPUs |        T |       T-8 |       F | default |      50 |   5.27 | 8 CPUs             |
|           |        T |       T-8 |       F | default |     500 |   5.21 |                    |
|           |        T |       T-8 |       F | default |    2000 |   5.46 |                    |
|           |        T |       T-8 |       F | default |   20000 |   5.51 |                    |
| F2-16 CPUs|        T |      T-16 |       F | default |      50 |   4.40 | 16 CPUs            |
|           |        T |      T-16 |       F | default |     500 |   4.50 |                    |
|           |        T |      T-16 |       F | default |    2000 |   4.65 |                    |
|           |        T |      T-16 |       F | default |   20000 |   4.22 |                    |
| F3-LK T   |        F |    F (16) |       T | default |      50 |   4.08 | Threads locked to  |
|           |        F |    F (16) |       T | default |     500 |   3.95 | CPU but Worker     |
|           |        F |    F (16) |       T | default |    2000 |   3.88 | could use diff CPU |
|           |        F |    F (16) |       T | default |   20000 |   4.77 |                    |
| F4 - PRTY |        F |    F (16) |       F |       0 |      50 |   4.02 |                    |
|           |        F |    F (16) |       F |       0 |     500 |   3.90 |                    |
|           |        F |    F (16) |       F |       0 |    2000 |   3.86 |                    |
|           |        F |    F (16) |       F |       0 |   20000 |   3.54 |                    |
| F1,2,3,4  |        T |      T-16 |       T |       0 |      50 |   4.71 | Workers locked to  |
|           |        T |      T-16 |       T |       0 |     500 |   4.78 | CPU and Thread     |
|           |        T |      T-16 |       T |       0 |    2000 |   4.90 |                    |
|           |        T |      T-16 |       T |       0 |   20000 |   4.96 |                    |

You can see the **more routines doing a small calculations** is the key.
Lightweight goroutines are amazing.
