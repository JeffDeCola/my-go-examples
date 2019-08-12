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
#include <pthread.h>

void lock_thread(int cpuid) {
        pthread_t tid;
        cpu_set_t cpuset;

        tid = pthread_self();
        CPU_ZERO(&cpuset);
        CPU_SET(cpuid, &cpuset);
    pthread_setaffinity_np(tid, sizeof(cpu_set_t), &cpuset);
}
*/
import "C"

// GO RUNTIME
const numberCores = 5   // Number of CPU you want to use
const lockThread = true // locked the goroutine to a thread (Done in go runtime)
const lockCore = true   // locked the thread to a core (Done in C)

// WORKERS
const useGoroutine = true // Do you want to use goroutines
const numberWorkers = 10  // Number of workers
const timeWork = 5        // Amount of time it takes a worker to finish

// BUFFER CHANNEL
var channelBufferSize = numberWorkers + 1 // How many channel buffers

// struct to pass in channel
type workerStats struct {
	id       int
	timeTook time.Duration
	cpuID    C.int
	pid      int
	tid      int
}

func doWork(msgCh chan *workerStats, wg *sync.WaitGroup, id int) {

	timeStart := time.Now().UTC()

	// Lock this goroutine to a particular thread (go runtime won't change threads)
	if lockThread {
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()
	}

	// Now lock this thread to CPU (This is outside of go)
	if lockCore {
		// Get the cpu your are on
		cpuID := C.sched_getcpu()
		// lock the thread to a cpu
		C.lock_thread(C.int(cpuID))
	}

	// Get the start specs
	startcpuID := C.sched_getcpu()
	startpid := os.Getpid()
	starttid := syscall.Gettid()

	// Print out some stats
	fmt.Printf("    Worker: %v, cpuID: %v, pid: %v, tid: %v\n", id, startcpuID, startpid, starttid)

	// Do something
	time.Sleep(timeWork * time.Second)
	timeEnd := time.Now().UTC()
	time.Sleep(timeWork * time.Second)
	time.Sleep(timeWork * time.Second)
	timeTook := timeEnd.Sub(timeStart)

	// Get the end specs
	endcpuID := C.sched_getcpu()
	endpid := os.Getpid()
	endtid := syscall.Gettid()

	// Check if anything changed
	if startcpuID != endcpuID {
		fmt.Printf("    ***WARNING*** Worker: %v used different cpuID\n", id)
	}
	if startpid != endpid {
		fmt.Printf("    ***WARNING*** Worker: %v used different pid\n", id)
	}
	if starttid != endtid {
		fmt.Printf("    ***WARNING*** Worker: %v used different tid\n", id)
	}

	msgCh <- &workerStats{
		id:       id,
		timeTook: timeTook,
		cpuID:    endcpuID,
		pid:      endpid,
		tid:      endtid,
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
	fmt.Printf("    const lockThread = %v\n", lockThread)
	fmt.Printf("    const lockCore = %v\n", lockCore)
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
