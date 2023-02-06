// my-go-examples goroutines-multi-core.go

package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"runtime"
	"sync"
	"syscall"

	log "github.com/sirupsen/logrus"

	"time"
)

/*
#define _GNU_SOURCE
#include <sched.h>
#include <pthread.h>

// FEATURE 3 - LOCK THREAD TO CPU/CORE
void lock_thread(int cpuid) {
    pthread_t tid;
    cpu_set_t cpuset;
    tid = pthread_self();
    CPU_ZERO(&cpuset);
    CPU_SET(cpuid, &cpuset);
    pthread_setaffinity_np(tid, sizeof(cpu_set_t), &cpuset);
}

// FEATURE 2 - PIN THREAD TO CPU/CORE
void set_affinity(int cpuid) {
    cpu_set_t my_set;            // Define your cpu_set bit mask
    CPU_ZERO(&my_set);           // Initialize it all to 0, i.e. no CPUs selected
    CPU_SET(cpuid, &my_set);     // Set the bit that represents core
    sched_setaffinity(0, sizeof(cpu_set_t), &my_set); // Set affinity of this process to
}
*/
import "C"

const toolVersion = "2.0.0"

var errLogLevel = errors.New("please use trace, info or error")

// WORKERS
const useGoroutine = true // Do you want to use goroutines
const numberWorkers = 50  // Number of workers

// PRIME NUMBER
const testForPrimes = 200000 // Find all prime numbers up to this number (brute force way)
// This must be divisible by the numberWorkers

// BUFFER CHANNEL
var channelBufferSize = numberWorkers + 1 // How many channel buffers

// FEATURE 1 - LOCK A GOROUTINE TO A THREAD
const lockGoroutineToThread = false // locked the goroutine to a thread (Done in go runtime)

// FEATURE 2 - LOCK A GOROUTINE TO A CPU
const lockGoroutineToCPU = false // Do you want to use particular CPUs?
// var useTheseCPUs = []int{9}     // Which CPU/Cores to use. These will rotate
// var useTheseCPUs = []int{0, 1, 2, 3, 4, 5, 6, 7}
var useTheseCPUs = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

// FEATURE 3 -  LOCK A THREAD TO A CPU
const lockThreadToCore = false // locked the thread to a cpu/core (Done in C)

// FEATURE 4 - SET PRIORITY ON THREAD
const setPriority = false  // Set the thread priority the goroutine
const setPriorityLevel = 0 // (0 to 39 with 0 highest)

// struct to pass in channel message
type workerStats struct {
	id             int
	timeTook       time.Duration
	cpuID          C.int
	pid            int
	tid            int
	threadPriority int
	foundPrimes    int
}

func setLogLevel(logLevel string) error {

	// SET LOG LEVEL (trace, info or error) None of which exit
	log.Trace("Set Log Level")

	switch logLevel {
	case "trace":
		log.SetLevel(log.TraceLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.ErrorLevel)
		return fmt.Errorf("%s", errLogLevel)
	}

	// SET FORMAT
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: false,
	})

	// SET OUTPUT (DEFAULT stderr)
	log.SetOutput(os.Stdout)

	return nil

}

func doWork(chMsg chan *workerStats, wg *sync.WaitGroup, id int, useCPU int, startNumber int, endNumber int) {

	timeStart := time.Now().UTC()

	// FEATURE 1 - LOCK A GOROUTINE TO A THREAD
	if lockGoroutineToThread {
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()
	}

	// FEATURE 2 - LOCK A GOROUTINE TO A CPU
	if lockGoroutineToCPU {
		C.set_affinity(C.int(useCPU))
	}

	// FEATURE 3 - LOCK A THREAD TO A CPU
	if lockThreadToCore {
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

	// START SPECS
	startcpuID := C.sched_getcpu()                                         // CPU ID
	startpid := syscall.Getpid()                                           // PROCESS ID
	starttid := syscall.Gettid()                                           // THREAD ID
	startthreadPriority, _ := syscall.Getpriority(syscall.PRIO_PROCESS, 0) // THREAD PRIORITY

	// SOME STATS
	s := fmt.Sprintf("START:     Worker: %2d, cpuID: %2d, process id: %6d, thread id: %6d, threadPriority: %4d, startNumber: %10d, endNumber: %10d\n", id, startcpuID, startpid, starttid, startthreadPriority, startNumber, endNumber)
	log.Debug(s)

	// DO SOMETHING TAX THE PROCESSOR A BIT
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

	// END SPECS
	timeEnd := time.Now().UTC()
	timeTook := timeEnd.Sub(timeStart)
	endcpuID := C.sched_getcpu()                                         // CPU ID
	endpid := syscall.Getpid()                                           // PROCESS ID
	endtid := syscall.Gettid()                                           // THREAD ID
	endthreadPriority, _ := syscall.Getpriority(syscall.PRIO_PROCESS, 0) // THREAD PRIORITY

	// DID ANYTHING CHANGE
	if startcpuID != endcpuID {
		s := fmt.Sprintf("    ***WARNING*** Worker: %v used different CPU ID (From %d to %d)\n", id, startcpuID, endcpuID)
		log.Info(s)
	}
	if startpid != endpid {
		s := fmt.Sprintf("    ***WARNING*** Worker: %v used different Process ID (From %d to %d)\n", id, startpid, endpid)
		log.Info(s)
	}
	if starttid != endtid {
		s := fmt.Sprintf("    ***WARNING*** Worker: %v used different Thread ID (From %d to %d)\n", id, starttid, endtid)
		log.Info(s)
	}
	if startthreadPriority != endthreadPriority {
		s := fmt.Sprintf("    ***WARNING*** Worker: %v used different Thread Priority (From %d to %d)\n", id, startthreadPriority, endthreadPriority)
		log.Info(s)
	}

	// SOME STATS
	s = fmt.Sprintf("COMPLETE:  Worker: %2d, cpuID: %2d, process id: %6d, thread id: %6d, threadPriority: %4d, startNumber: %10d, endNumber: %10d\n", id, endcpuID, endpid, endtid, endthreadPriority, startNumber, endNumber)
	log.Debug(s)

	// SEND INFO FOR STATS
	chMsg <- &workerStats{
		id:             id,
		timeTook:       timeTook,
		cpuID:          endcpuID,
		pid:            endpid,
		tid:            endtid,
		threadPriority: int(endthreadPriority),
		foundPrimes:    foundPrimes,
	}

	// WAITGROUP DONE
	wg.Done()

}

func getStats(chMsg chan *workerStats, wg *sync.WaitGroup, numberWorkers int, start time.Time) {

	// BLOCK - WAIT FOR ALL WORKERS TO FINISH
	fmt.Println("Waiting for all the workers to finish")
	fmt.Printf("There are currently %v goroutines running\n", runtime.NumGoroutine())
	wg.Wait()

	// Print out stats from each worker
	var cpuSet []int
	totalCPUsUsed := 0
	threadSet := make(map[int][]int) // A map of slices (for each cpu)
	totalThreadsUsed := make(map[int]int)
	totalWorkersUsed := make(map[int]int)
	totalPrimesFound := 0

	// LOOP OVER ALL WORKERS
	for i := 0; i < numberWorkers; i++ {

		r := <-chMsg
		s := fmt.Sprintf("Worker %2d, cpuID: %2d, pid: %6d, tid: %6d, threadPriority: %4d foundPrimes: %6d (Took %v)\n", r.id, r.cpuID, r.pid, r.tid, r.threadPriority, r.foundPrimes, r.timeTook)
		log.Debug(s)

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
	fmt.Printf("\nThe total CPU/Cores used: %v\n", totalCPUsUsed)
	for _, cpuID := range cpuSet {
		fmt.Printf("  -cpuID %2d %6d Workers used %3d Threads\n", cpuID, totalWorkersUsed[cpuID], totalThreadsUsed[cpuID])
	}

	fmt.Printf("\nThe total Primes numbers under %d is %d\n", testForPrimes, totalPrimesFound)
	fmt.Printf("Total time to calculate was: %f seconds\n", time.Since(start).Seconds())
	fmt.Println("\nDone - Press Return to exit")

}

func main() {

	// FLAGS
	version := flag.Bool("v", false, "prints current version")
	logLevel := flag.String("loglevel", "error", "log level (trace, info or error)")
	flag.Parse()

	// CHECK VERSION
	if *version {
		fmt.Println(toolVersion)
		os.Exit(0)
	}

	// SET LOG LEVEL (trace, info or error) None of which exit
	err := setLogLevel(*logLevel)
	if err != nil {
		log.Errorf("Error getting logLevel: %s", err)
	}

	// PRINT OUT FLAGS
	log.Trace("Version flag = ", *version)
	log.Trace("Log Level = ", *logLevel)

	// START
	startTime := time.Now()
	fmt.Println("")

	// PRINT SOME STATS
	fmt.Printf("Main start time: %f seconds\n\n", time.Since(startTime).Seconds())

	// SETTINGS
	fmt.Println("Your settings are:")
	fmt.Println("WORKERS")
	fmt.Printf("    const useGoroutine = %v\n", useGoroutine)
	fmt.Printf("    const numberWorkers = %v\n", numberWorkers)
	fmt.Printf("    const testForPrimes = %v\n", testForPrimes)
	fmt.Println("BUFFER CHANNEL")
	fmt.Printf("    const channelBufferSize = %v\n", channelBufferSize)
	fmt.Println("FEATURE 1 - LOCK A GOROUTINE TO A THREAD")
	fmt.Printf("    const lockGoroutineToThread = %v\n", lockGoroutineToThread)
	fmt.Println("FEATURE 2 - LOCK A GOROUTINE TO A CPU")
	fmt.Printf("    const lockGoroutineToCPU = %v\n", lockGoroutineToCPU)
	fmt.Printf("    useTheseCPUs = %v\n", useTheseCPUs)
	fmt.Println("FEATURE 3 - LOCK A THREAD TO A CPU")
	fmt.Printf("    const lockThreadToCore = %v\n", lockThreadToCore)
	fmt.Println("FEATURE 4 - SET PRIORITY ON THREAD")
	fmt.Printf("    const setPriority = %v\n", setPriority)
	fmt.Printf("    const setPriorityLevel = %v\n", setPriorityLevel)
	fmt.Println("")

	// HOW MANY CORES CAN WE USE
	fmt.Println("Total number of cores on this machine: ", runtime.NumCPU())
	// Limit number of cores to numberCores
	// runtime.GOMAXPROCS(numberCores)
	fmt.Println("Number of Cores you set: ", runtime.GOMAXPROCS(-1))
	fmt.Println("")

	// CHANNEL BUFFER FOR MESSAGES
	chMsg := make(chan *workerStats, channelBufferSize)

	// DIVIDE WORKLOAD
	if testForPrimes%numberWorkers != 0 {
		log.Fatal("Can't split workload evenly - change testforPrimes or numberWorkers")
	}
	splitWorkload := testForPrimes / numberWorkers
	startNumber := 1
	endNumber := splitWorkload

	// CREATE WORKERS AND SPLIT WORKLOAD (START AND END NUMBERS)
	// Create wait group
	var wg sync.WaitGroup
	c := 0 // Used to rotate through CPUs
	fmt.Println("Kicking off workers...")
	for id := 1; id < numberWorkers+1; id++ {

		// ADD WAITGROUP TO SEND TO GOROUTINE
		wg.Add(1)

		// WHAT CPU TO USE
		if c == len(useTheseCPUs) {
			c = 0
		}
		useCPU := useTheseCPUs[c]
		c++

		// KICK OFF A WORKER
		if useGoroutine {
			log.Debugf("Kicking off goroutine worker %2d for CPU %2d (Numbers %8d - %8d)\n", id, useCPU, startNumber, endNumber)
			go doWork(chMsg, &wg, id, useCPU, startNumber, endNumber)
		} else {
			log.Debugf("Kicking off worker %2d for CPU %2d (Numbers %8d - %8d)\n", id, useCPU, startNumber, endNumber)
			doWork(chMsg, &wg, id, useCPU, startNumber, endNumber)
		}
		startNumber = startNumber + splitWorkload
		endNumber = endNumber + splitWorkload
	}

	// DONE KICKING OFF WORKERS
	if useGoroutine {
		fmt.Printf("Finished Kicking off all goroutines in %f seconds\n\n", time.Since(startTime).Seconds())
	} else {
		fmt.Printf("Finished all work in %f seconds\n\n", time.Since(startTime).Seconds())
	}

	// GATHER STATS FROM WORKERS
	go getStats(chMsg, &wg, numberWorkers, startTime)

	// PRESS RETURN TO EXIT
	fmt.Scanln()
	fmt.Println("EOF...")

}
