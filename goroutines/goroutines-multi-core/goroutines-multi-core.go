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
void set_affinity(int cpuid) {
        cpu_set_t my_set;        // Define your cpu_set bit mask
        CPU_ZERO(&my_set);       // Initialize it all to 0, i.e. no CPUs selected
        CPU_SET(cpuid, &my_set);     // Set the bit that represents core
        sched_setaffinity(0, sizeof(cpu_set_t), &my_set); // Set affinity of this process to
}
*/
import "C"

const debug = false

// GO RUNTIME & OS FEATURES
// FEATURE 1 - LOCK A GOROUTINE TO A THREAD
const lockThread = true // locked the goroutine to a thread (Done in go runtime)
// FEATURE 2 - PIN A GOROUTINE TO A CPU (set affinity)
const useParticularCPUs = true                                                 // Do you want to use particular CPUs?
var usetheseCPUs = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15} // Which CPU/Cores to use. These will rotate
// FEATURE 3 - LOCK A THREAD TO A CPU/CORE
const lockCore = true // locked the thread to a core (Done in C)
// FEATURE 4 - SET PRIORITY ON THREAD
const setPriority = true   // Set the thread priority the goroutine
const setPriorityLevel = 0 // (0 to 39 with 0 highest)

// WORKERS
const useGoroutine = true    // Do you want to use goroutines
const numberWorkers = 50000  // Number of workers
const testforPrimes = 200000 // Find all prime numbers up to this number (brute force way)
// This must be divisible by the numberWorkers

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
	foundPrimes    int
}

func doWork(msgCh chan *workerStats, wg *sync.WaitGroup, id int, useCPU int, startNumber int, endNumber int) {

	timeStart := time.Now().UTC()

	// FEATURE 1 - LOCK A GOROUTINE TO A THREAD
	if lockThread {
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()
	}

	// FEATURE 2 - PIN A GOROUTINE TO A CPU (set affinity)
	if useParticularCPUs {
		C.set_affinity(C.int(useCPU))
	}

	// FEATURE 3 - LOCK A THREAD TO A CPU/CORE
	if lockCore {
		// Get the cpu your are on
		cpuID := C.sched_getcpu()
		// lock the thread to a cpu
		C.lock_thread(C.int(cpuID))
	}

	// FEATURE 4 - SET PRIORITY ON THREAD
	if setPriority {
		err := syscall.Setpriority(syscall.PRIO_PROCESS, 0, setPriorityLevel)
		if err != nil {
			println("Setpriority failed")
			return
		}
	}
	// Put it back to what it was (I can't get this to work correctly)
	//defer syscall.Setpriority(syscall.PRIO_PROCESS, syscall.Getpid(), startthreadPriority)

	// Get the start specs
	startcpuID := C.sched_getcpu()
	startpid := syscall.Getpid()
	starttid := syscall.Gettid()
	startthreadPriority, _ := syscall.Getpriority(syscall.PRIO_PROCESS, 0)

	// Print out some stats
	if debug {
		fmt.Printf("    Worker: %v, cpuID: %v, pid: %v, tid: %v, threadPriority: %v, startNumber: %v, endNumber: %v\n", id, startcpuID, startpid, starttid, startthreadPriority, startNumber, endNumber)
	}
	// DO SOMETHING (DON'T SLEEP) TAX THE PROCESSOR A BIT
	// Find prime number the brute force way, by dividing.
	// Find all Prime numbers up to n
	// A prime is defined as any counting number that is divisible
	// by exactly two distinct numbers: 1 and itself. Therefore:
	foundPrimes := 0
	for n := startNumber; n < endNumber; n++ {
		// fmt.Printf("For number n %v\n", n)
		// Divide by all number before it and check the remiander
		count := 0
		for i := 1; i < n+1; i++ {
			remainder := n % i
			// fmt.Printf("   Divide by i %v remainder is %v\n", i, remainder)
			if remainder == 0 {
				count++
			}
		}
		// So did you have only two counts (i.e. 1 and itself)? If so, you have a prime number.
		if count == 2 {
			//fmt.Printf("Number n %v is a prime\n", n)
			foundPrimes++
		}
	}

	timeEnd := time.Now().UTC()
	timeTook := timeEnd.Sub(timeStart)

	// Get the end specs
	endcpuID := C.sched_getcpu()
	endpid := syscall.Getpid()
	endtid := syscall.Gettid()
	endthreadPriority, _ := syscall.Getpriority(syscall.PRIO_PROCESS, 0)

	// Check if anything changed
	if debug {
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
	}

	msgCh <- &workerStats{
		id:             id,
		timeTook:       timeTook,
		cpuID:          endcpuID,
		pid:            endpid,
		tid:            endtid,
		threadPriority: int(endthreadPriority),
		foundPrimes:    foundPrimes,
	}
	wg.Done()

}

func getStats(msgCh chan *workerStats, wg *sync.WaitGroup, numberWorkers int, start time.Time) {

	// WAIT FOR ALL WORKERS TO FINISH
	fmt.Println("getStats - Waiting for all the workers to finish")
	fmt.Printf("getStats - There are currently %v goroutines running\n", runtime.NumGoroutine())

	wg.Wait()

	// Print out stats from each worker
	var cpuSet []int
	totalCPUsUsed := 0
	threadSet := make(map[int][]int) // A map of slices (for each cpu)
	totalThreadsUsed := make(map[int]int)
	totalWorkersUsed := make(map[int]int)
	totalPrimesFound := 0

	for i := 0; i < numberWorkers; i++ {
		r := <-msgCh
		if debug {
			fmt.Printf("getStats - From Worker %v, Took: %v, cpuID: %v, pid: %v, tid: %v, threadPriority: %v foundPrimes: %v\n", r.id, r.timeTook, r.cpuID, r.pid, r.tid, r.threadPriority, r.foundPrimes)
		}

		// totalCPUUsed - check if you have CPU in your slice
		newCPUFound := true
		for _, element := range cpuSet {
			if element == int(r.cpuID) {
				newCPUFound = false
				break
			}
		}
		if newCPUFound {
			totalCPUsUsed++
			cpuSet = append(cpuSet, int(r.cpuID))
		}

		// totalThreadsUsed[cpuID] - check if you have a thread in your map for a particular CPU
		newThreadFound := true
		for _, element := range threadSet[int(r.cpuID)] {
			if element == int(r.tid) {
				newThreadFound = false
				break
			}
		}
		if newThreadFound {
			totalThreadsUsed[int(r.cpuID)]++
			threadSet[int(r.cpuID)] = append(threadSet[int(r.cpuID)], int(r.tid))
		}

		// Total workers
		totalWorkersUsed[int(r.cpuID)]++

		// Total Primes found
		totalPrimesFound = totalPrimesFound + r.foundPrimes

	}

	fmt.Printf("")
	fmt.Printf("** The total CPU/Cores used was %v\n", totalCPUsUsed)
	for _, cpuID := range cpuSet {
		fmt.Printf("  For CPU/Core %v", cpuID)
		fmt.Printf("  %v Workers used %v Threads\n", totalWorkersUsed[cpuID], totalThreadsUsed[cpuID])
	}
	fmt.Printf("** The total Primes under %v is %v\n", testforPrimes, totalPrimesFound)
	fmt.Printf("** End time: %f seconds\n", time.Since(start).Seconds())
	fmt.Println("** Press Return to exit")

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
	fmt.Println("Your settings are:")
	fmt.Println("FEATURE 1 - LOCK A GOROUTINE TO A THREAD")
	fmt.Printf("    const lockThread = %v\n", lockThread)
	fmt.Println("FEATURE 2 - PIN A GOROUTINE TO A CPU (set affinity)")
	fmt.Printf("    const useParticularCPUs = %v\n", useParticularCPUs)
	fmt.Printf("    usetheseCPUs = %v\n", usetheseCPUs)
	fmt.Println("FEATURE 3 - LOCK A THREAD TO A CPU/CORE")
	fmt.Printf("    const lockCore = %v\n", lockCore)
	fmt.Println("FEATURE 4 - SET PRIORITY ON THREAD")
	fmt.Printf("    const setPriority = %v\n", setPriority)
	fmt.Printf("    const setPriorityLevel = %v\n", setPriorityLevel)
	fmt.Println("WORKERS")
	fmt.Printf("    const useGoroutine = %v\n", useGoroutine)
	fmt.Printf("    const numberWorkers = %v\n", numberWorkers)
	fmt.Printf("    const testforPrimes = %v\n", testforPrimes)
	fmt.Println("BUFFER CHANNEL")
	fmt.Printf("    const channelBufferSize = %v\n", channelBufferSize)
	// How many cores can i use?
	fmt.Println("Total number of cores on this machine: ", runtime.NumCPU())
	// Limit number of cores to numberCores
	// runtime.GOMAXPROCS(numberCores)
	fmt.Println("Number of Cores you set ", runtime.GOMAXPROCS(-1))
	fmt.Println("")

	// Make Channel Buffer
	msgCh := make(chan *workerStats, channelBufferSize)

	// Divide the workload
	if testforPrimes%numberWorkers != 0 {
		fmt.Println("Can't split workload evenly - change testforPrimes or numberWorkers")
		os.Exit(3)
	}
	splitWorkload := testforPrimes / numberWorkers
	startNumber := 1
	endNumber := splitWorkload

	// Create Workers and split workload (start and end numbers)
	c := 0
	for id := 1; id < numberWorkers+1; id++ {
		wg.Add(1)
		if c == len(usetheseCPUs) {
			c = 0
		}
		useCPU := usetheseCPUs[c]
		c++
		if debug {
			fmt.Printf("Starting worker %v using CPU %v\n", id, useCPU)
		}
		if useGoroutine {
			go doWork(msgCh, &wg, id, useCPU, startNumber, endNumber)
		} else {
			doWork(msgCh, &wg, id, useCPU, startNumber, endNumber)
		}
		startNumber = startNumber + splitWorkload
		endNumber = endNumber + splitWorkload
	}
	fmt.Printf("Finished Kicking off all goroutines in %f seconds\n", time.Since(start).Seconds())

	// Gather Stats from workers
	go getStats(msgCh, &wg, numberWorkers, start)

	// Press return to exit
	fmt.Scanln()
	fmt.Println("EOF...")
}
