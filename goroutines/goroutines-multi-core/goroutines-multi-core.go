package main

import (
	"fmt"
	"runtime"
	"sync"

	"time"
)

// WORKERS AND CORES
const useGoroutine = true
const numberWorkers = 10
const timeWork = 1
const numberCores = 2

// OTHER
var channelBufferSize = numberWorkers + 1 // How many channel buffers

// struct to pass in channel
type workerStats struct {
	id       int
	timeTook time.Duration
}

func doWork(msgCh chan *workerStats, wg *sync.WaitGroup, id int) {

	timeStart := time.Now().UTC()
	time.Sleep(timeWork * time.Second)
	timeEnd := time.Now().UTC()
	timeTook := timeEnd.Sub(timeStart)

	msgCh <- &workerStats{
		id:       id,
		timeTook: timeTook,
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
		fmt.Printf("From Worker %v, Took %v\n", r.id, r.timeTook)
	}

}

func main() {

	// START
	start := time.Now()

	// Create wait group
	var wg sync.WaitGroup

	// SET CORES
	runtime.GOMAXPROCS(numberCores)

	fmt.Println("Main start time", time.Since(start))

	// How many core amd i using
	fmt.Println("I have this many cores on this machine: ", runtime.NumCPU())

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

	// Gather Stats from workers
	getStats(msgCh, &wg, numberWorkers)

	// END
	fmt.Println("Main end time", time.Since(start))

}
