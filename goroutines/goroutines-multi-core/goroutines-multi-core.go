package main

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
	"syscall"

	"time"
)

/*
#define _GNU_SOURCE
#include <sched.h>
*/
import "C"

//CORES
const numberCores = 10

// WORKERS
const useGoroutine = true
const numberWorkers = 50
const timeWork = 1

// OTHER
var channelBufferSize = numberWorkers + 1 // How many channel buffers

// struct to pass in channel
type workerStats struct {
	id        int
	timeTook  time.Duration
	pid1      int
    pid2      int
    cpu1      int
	threadID1 uint64
	threadID2 int
}

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func doWork(msgCh chan *workerStats, wg *sync.WaitGroup, id int) {

	// Lock this goroutine to a particular thread
	// runtime.LockOSThread()
	//defer runtime.UnlockOSThread()

	timeStart := time.Now().UTC()
	time.Sleep(timeWork * time.Second)
	timeEnd := time.Now().UTC()
	timeTook := timeEnd.Sub(timeStart)

	fmt.Printf("worker: %d, CPU: %d\n", id, C.sched_getcpu())

	msgCh <- &workerStats{
		id:        id,
		timeTook:  timeTook,
		pid1:      os.Getpid(),
		pid2:      syscall.Getpid(),
		threadID1: getGID(),
		threadID2: syscall.Gettid(),
	}
	wg.Done()

}

func getStats(msgCh chan *workerStats, wg *sync.WaitGroup, numberWorkers int) {

	// WAIT FOR ALL WORKERS TO FINISH
	fmt.Println("Waiting for all the workers to finish")
	fmt.Println("There are currently this many goroutines", runtime.NumGoroutine())
	wg.Wait()

	// Print out stats from each worker
	for i := 0; i < numberWorkers; i++ {
		r := <-msgCh
		fmt.Printf("From Worker %v, Took %v on pid %v or pid %v and tid %v or tid %v\n", r.id, r.timeTook, r.pid2, r.pid2, r.threadID1, r.threadID2)
	}

}

func main() {

	// START
	start := time.Now()

	// Create wait group
	var wg sync.WaitGroup

	fmt.Printf("Main start time %f seconds\n", time.Since(start).Seconds())

	// How many cores can i use?
	fmt.Println("Total number of cores on this machine: ", runtime.NumCPU())
	runtime.GOMAXPROCS(numberCores)
	fmt.Println("Number of Cores that can be used here: ", runtime.GOMAXPROCS(-1))

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
	fmt.Printf("Main end time in %f seconds\n", time.Since(start).Seconds())

}
