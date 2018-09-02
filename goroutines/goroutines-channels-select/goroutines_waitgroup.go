package main

import (
	"fmt"
	"sync"

	"time"
)

func doWork(i int) {
	time.Sleep(1 * time.Second)
	fmt.Println(i, "done")
}

func main() {

	// Step 1 - Create new instance of wait group
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		
		// Step 2 - The number of  goroutines to wait for (to be done) - Just 1
		wg.Add(1)

		go func(i int) {

			// Step 3 - Execute in the goroutine to indicate that the function finished
			defer wg.Done()

			doWork(i)
		}(i)
	}

	// Step 4 Call wait when we want to block
	wg.Wait()
	fmt.Println("All done!")
}
