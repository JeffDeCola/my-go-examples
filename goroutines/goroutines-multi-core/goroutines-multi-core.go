package main

import (
	"fmt"

	"time"
)

// WORKERS AND CORES
const numberWorkers = 20
const numberCores = 3

// OTHER
var channelBufferSize = 1 // How many channel buffers

// struct to pass in channel
type workerStats struct {
	id       int
	timeTook time.Duration
}

func doWork(msgCh chan *workerStats, id int) {

	timeStart := time.Now().UTC()
	time.Sleep(3 * time.Second)
	timeEnd := time.Now().UTC()
	timeTook := timeEnd.Sub(timeStart)

	msgCh <- &workerStats{
		id:       id,
		timeTook: timeTook,
	}

}

func getStats(msgCh chan *workerStats) {
	// Print out stats from each worker
	for r := range msgCh {
		fmt.Printf("From Worker %v, Took %v\n", r.id, r.timeTook)
	}
}

func main() {
	// START
	fmt.Println("Main start time", time.Since(start))

	// Make Channel Buffer
	msgCh := make(chan *workerStats, channelBufferSize)

	// Create Workers on Core
	for id := 1; id < numberWorkers+1; id++ {
		fmt.Println("Starting worker", id)
		go doWork(msgCh, id)

	}

	// Gather Stats from workers
	go getStats(msgCh)

	// END
	fmt.Println("Main end time", time.Since(start))

}
