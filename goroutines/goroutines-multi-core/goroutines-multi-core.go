package main

import (
	"fmt"
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
void set_affinity(int cpuid) {
        cpu_set_t my_set;        // Define your cpu_set bit mask
        CPU_ZERO(&my_set);       // Initialize it all to 0, i.e. no CPUs selected
        CPU_SET(cpuid, &my_set);     // Set the bit that represents core
        sched_setaffinity(0, sizeof(cpu_set_t), &my_set); // Set affinity of this process to
}
*/
import "C"

// GO RUNTIME & OS
const lockThread = true        // locked the goroutine to a thread (Done in go runtime)
const useParticularCPUs = true // Do you want to use particular CPUs?
var usetheseCPUs = []int{5}    // Which CPU/Cores to use. These will rotate
const lockCore = true          // locked the thread to a core (Done in C)
const setPriorityThread = 7    // Set the thread priority the goroutine is on (0 to 39 with 0 highest)

// WORKERS
const useGoroutine = true // Do you want to use goroutines
const numberWorkers = 5   // Number of workers
const timeWork = 30       // Amount of time it takes a worker to finish

// BUFFER CHANNEL
var channelBufferSize = numberWorkers + 1 // How many channel buffers

// struct to pass in channel
type workerStats struct {
	id             int
	timeTook       time.Duration
	cpuID          C.int
	pid            int
	tid            int
	threadPriority int
}

func doWork(msgCh chan *workerStats, wg *sync.WaitGroup, id int, useCPU int) {

	timeStart := time.Now().UTC()

	// Lock this goroutine to a particular thread (go runtime won't change threads)
	if lockThread {
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()
	}

	// Set goroutine to a particular Core/CPU - Set affinity()
	if useParticularCPUs {
		C.set_affinity(C.int(useCPU))
	}

	// Now lock this thread to the Core/CPU you are on
	if lockCore {
		// Get the cpu your are on
		cpuID := C.sched_getcpu()
		// lock the thread to a cpu
		C.lock_thread(C.int(cpuID))
	}

	// Get the start specs
	startcpuID := C.sched_getcpu()
	startpid := syscall.Getpid()
	starttid := syscall.Gettid()
	//startthreadPriority, _ := syscall.Getpriority(syscall.PRIO_PROCESS, 0)
	startthreadPriority, _, _ := syscall.Syscall(syscall.SYS_GETPRIORITY, 0, 0, 0)

	// Set the priority of the thread using system call
	err := syscall.Setpriority(syscall.PRIO_PROCESS, 0, setPriorityThread)
	if err != nil {
		println("Setpriority failed")
		return
	}
	// Put it back to what it was (I can't get this to work correctly)
	//defer syscall.Setpriority(syscall.PRIO_PROCESS, syscall.Getpid(), startthreadPriority)

	// Print out some stats
	fmt.Printf("    Worker: %v, cpuID: %v, pid: %v, tid: %v, threadPriority: %v\n", id, startcpuID, startpid, starttid, startthreadPriority)

	// Do something
	time.Sleep(timeWork * time.Second)
	timeEnd := time.Now().UTC()
	timeTook := timeEnd.Sub(timeStart)

	// Get the end specs
	endcpuID := C.sched_getcpu()
	endpid := syscall.Getpid()
	endtid := syscall.Gettid()
	//endthreadPriority, _ := syscall.Getpriority(syscall.PRIO_PROCESS, 0)
	endthreadPriority, _, _ := syscall.Syscall(syscall.SYS_GETPRIORITY, 0, 0, 0)

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
	if startthreadPriority != endthreadPriority {
		fmt.Printf("    ***WARNING*** Worker: %v used different threadPriority\n", id)
	}

	msgCh <- &workerStats{
		id:             id,
		timeTook:       timeTook,
		cpuID:          endcpuID,
		pid:            endpid,
		tid:            endtid,
		threadPriority: int(endthreadPriority),
	}
	wg.Done()

}

func getStats(msgCh chan *workerStats, wg *sync.WaitGroup, numberWorkers int, start time.Time) {

	// WAIT FOR ALL WORKERS TO FINISH
	fmt.Println("getStats - Waiting for all the workers to finish")
	fmt.Println("getStats - There are currently this many goroutines", runtime.NumGoroutine())

	wg.Wait()

	// Print out stats from each worker
	for i := 0; i < numberWorkers; i++ {
		r := <-msgCh
		fmt.Printf("getStats - From Worker %v, Took: %v, cpuID: %v, pid: %v, tid: %v, threadPriority: %v\n", r.id, r.timeTook, r.cpuID, r.pid, r.tid, r.threadPriority)
	}

	fmt.Printf("End time: %f seconds\n", time.Since(start).Seconds())
	fmt.Println("Press Return to exit")

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
	fmt.Printf("    const lockThread = %v\n", lockThread)
	fmt.Printf("    const useParticularCPUs = %v\n", useParticularCPUs)
	fmt.Printf("    usetheseCPUs = %v\n", usetheseCPUs)
	fmt.Printf("    const lockCore = %v\n", lockCore)
	fmt.Printf("    const useGoroutine = %v\n", useGoroutine)
	fmt.Printf("    const numberWorkers = %v\n", numberWorkers)
	fmt.Printf("    const timeWork = %v\n", timeWork)
	fmt.Printf("    const channelBufferSize = %v\n", channelBufferSize)
	// How many cores can i use?
	fmt.Println("Total number of cores on this machine: ", runtime.NumCPU())
	// Limit number of cores to numberCores
	// runtime.GOMAXPROCS(numberCores)
	fmt.Println("Number of Cores you set ", runtime.GOMAXPROCS(-1))
	fmt.Println("")

	// Make Channel Buffer
	msgCh := make(chan *workerStats, channelBufferSize)

	// Create Workers
	c := 0
	for id := 1; id < numberWorkers+1; id++ {
		wg.Add(1)
		if c == len(usetheseCPUs) {
			c = 0
		}
		useCPU := usetheseCPUs[c]
		c++
		fmt.Printf("Starting worker %v using CPU %v\n", id, useCPU)
		if useGoroutine {
			go doWork(msgCh, &wg, id, useCPU)
		} else {
			doWork(msgCh, &wg, id, useCPU)
		}
	}
	fmt.Printf("Finished Kicking off all goroutines in %f seconds\n", time.Since(start).Seconds())

	// Gather Stats from workers
	go getStats(msgCh, &wg, numberWorkers, start)

	// Press return to exit
	fmt.Scanln()
	fmt.Println("EOF...")
}
