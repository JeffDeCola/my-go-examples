// my-go-examples goroutines-waitgroup.go

package main

import (
	"fmt"
	"sync"

	"time"
)

const numberWorkers = 20

func doWork(wg *sync.WaitGroup, id int) {

	fmt.Println(id, "Started")
	time.Sleep(3 * time.Second)
	fmt.Println(id, "Done")
	wg.Done()

}

func main() {

	// CREATE WAITGROUP
	var wg sync.WaitGroup

	// KICK OFF WORKERS
	for i := 0; i < numberWorkers; i++ {
		// ADD WAITGROUP
		wg.Add(1)
		fmt.Println("Starting worker", i)
		go doWork(&wg, i)

	}

	// Waits for all the goroutines to finish
	fmt.Println("Waiting for all the workers to finish")
	wg.Wait()
	fmt.Println("All workers done!")

}
