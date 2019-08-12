package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"syscall"

	"time"
)

/*
#define _GNU_SOURCE
#include <sched.h>
*/
import "C"

// GO RUNTIME
const numberCores = 5    // Number of CPU you want to use
const lockthreads = true // locked the goroutine to a thread

// WORKERS
const useGoroutine = true // Do you want to use goroutines
const numberWorkers = 3   // Number of workers
const timeWork = 5        // Amount of time it takes a worker to finish

// BUFFER CHANNEL
var channelBufferSize = numberWorkers + 1 // How many channel buffers

// struct to pass in channel
type workerStats struct {
	id       int
	timeTook time.Duration
	cpuID    _Ctype_int
	pid      int
	tid      int
}

func doWork(msgCh chan *workerStats, wg *sync.WaitGroup, id int) {

	// Lock this goroutine to a particular thread (go runtime won't change threads)
	if lockthreads {
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()
	}

	// Get the start thread id
	startTID := syscall.Gettid()

	// Print out stats
	fmt.Printf("    Worker: %v, cpuID: %v, pid: %v, tid: %v\n", id, C.sched_getcpu(), os.Getpid(), syscall.Gettid())

	timeStart := time.Now().UTC()
	time.Sleep(timeWork * time.Second)
	timeEnd := time.Now().UTC()
	timeTook := timeEnd.Sub(timeStart)

	// Check if the go runtime moved onto a different thread
	endTID := syscall.Gettid()
	if startTID != endTID {
		fmt.Printf("    *** The go runtime moved Worker: %v onto a different thread\n", id)
	}

	msgCh <- &workerStats{
		id:       id,
		timeTook: timeTook,
		cpuID:    C.sched_getcpu(),
		pid:      os.Getpid(),
		tid:      syscall.Gettid(),
	}
	wg.Done()

}

func getStats(msgCh chan *workerStats, wg *sync.WaitGroup, numberWorkers int) {

	// WAIT FOR ALL WORKERS TO FINISH
	fmt.Println("getStats - Waiting for all the workers to finish")
	fmt.Println("getStats - There are currently this many goroutines", runtime.NumGoroutine())
	wg.Wait()

	// Print out stats from each worker
	for i := 0; i < numberWorkers; i++ {
		r := <-msgCh
		fmt.Printf("getStats - From Worker %v, Took: %v, cpuID: %v, pid: %v, tid: %v\n", r.id, r.timeTook, r.cpuID, r.pid, r.tid)
	}

}

func main() {

	// START
	start := time.Now()
	fmt.Println("")

	// Create wait group
	var wg sync.WaitGroup

	// PRINT SOME STATS
	fmt.Printf("Main start time: %f seconds\n", time.Since(start).Seconds())
	// Print constants
	fmt.Printf("Your settings are:\n")
	fmt.Printf("    const numberCores = %v\n", numberCores)
	fmt.Printf("    const lockthreads = %v\n", lockthreads)
	fmt.Printf("    const useGoroutine = %v\n", useGoroutine)
	fmt.Printf("    const numberWorkers = %v\n", numberWorkers)
	fmt.Printf("    const timeWork = %v\n", timeWork)
	fmt.Printf("    const channelBufferSize = %v\n", channelBufferSize)
	// How many cores can i use?
	fmt.Println("Total number of cores on this machine: ", runtime.NumCPU())
	// Limit number of cores to numberCores
	runtime.GOMAXPROCS(numberCores)
	fmt.Println("Number of Cores you set ", runtime.GOMAXPROCS(-1))
	fmt.Println("")

	// Make Channel Buffer
	msgCh := make(chan *workerStats, channelBufferSize)

	// Create Workers on Core
	for id := 1; id < numberWorkers+1; id++ {
		wg.Add(1)
		fmt.Println("Starting worker", id)
		if useGoroutine {
			go doWork(msgCh, &wg, id)
		} else {
			doWork(msgCh, &wg, id)
		}
	}
	fmt.Printf("Kicked off all goroutines in %f seconds\n", time.Since(start).Seconds())

	// Gather Stats from workers
	getStats(msgCh, &wg, numberWorkers)

	// END
	fmt.Printf("Main end time: %f seconds\n", time.Since(start).Seconds())

}
