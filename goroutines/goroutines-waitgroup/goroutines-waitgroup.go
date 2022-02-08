// goroutines-waitgroup.go

package main

import (
	"fmt"
	"sync"

	"time"
)

const numberWorkers = 10

func doWork(wg *sync.WaitGroup, id int) {

	// START
	fmt.Println(id, "Started")

	// DO WORK
	time.Sleep(3 * time.Second)

	// DONE
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

	// BLOCK TILL ALL WAITGROUP DONE
	fmt.Println("Waiting for all the workers to finish")
	wg.Wait()

	fmt.Println("All workers done!")

}
