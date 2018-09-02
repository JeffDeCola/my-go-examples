package main

import (
	"fmt"
	"sync"

	"time"
)

func doWork(i int) {
	time.Sleep(3 * time.Second)
	fmt.Println(i, "done")
}

func main() {

	// Step 1 - create new instance of wait group
	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		// Step 2 - The number of  goroutines to wait for
		wg.Add(1)
		go func(i int) {

			// Step 3 - Execute in the goroutine to indicate that the function finsihed
			defer wg.Done()
			time.Sleep(3 * time.Second)

			doWork(i)

		}(i)

	}

	// Step 4 Call wait when we want to block
	wg.Wait()
	fmt.Println("All done!")
}
