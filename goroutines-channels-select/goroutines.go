package main

import (
	"fmt"
	"time"
)

type Hello struct {
}

func (h Hello) Say(ch chan string, ticktime int) error {

	ticker := time.Tick(time.Millisecond * time.Duration(ticktime))

	// Loop forever - Long running
	for {
		select {

		case now := <-ticker:

			saythis := <-ch

			fmt.Printf("%v %s\n", now, saythis)

			/* case <-quit:

			fmt.Println("Bye")
			return nil */
		}
	}
	return nil
}

func main() {

	hi := Hello{}
	ch1 := make(chan string)
	ch2 := make(chan string)

	go hi.Say(ch1, 1000)
	go hi.Say(ch2, 500)

	time.Sleep(5 * time.Second)
	ch1 <- "Jeff"
	ch2 <- "Clif"

	time.Sleep(15 * time.Second)
	ch1 <- "Jack"
	ch2 <- "Jill"

	time.Sleep(29 * time.Second)

}
