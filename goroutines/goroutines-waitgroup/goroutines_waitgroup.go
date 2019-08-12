package main

import (
	"fmt"
	"sync"

	"time"
)

const numberWorkers = 20

func doWork(wg *sync.WaitGroup, id int) {
	// Call Done when finished
	//defer wg.Done()
	fmt.Println(id, "Started")
	time.Sleep(3 * time.Second)
	fmt.Println(id, "Done")
	wg.Done()
}

func main() {

	// Create wait group
	var wg sync.WaitGroup

	for i := 0; i < numberWorkers; i++ {
		// Add goroutines to wait for
		wg.Add(1)
		fmt.Println("Starting worker", i)
		go doWork(&wg, i)

	}

	// Waits for all the goroutines to finish
	fmt.Println("Waiting for all the workers to finish")
	wg.Wait()
	fmt.Println("All workers done!!!!!!!!!!!!!!!!!!")
}
