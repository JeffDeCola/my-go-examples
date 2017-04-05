package main

import (
	"fmt"
	"time"
)

type Hello struct {
	defaultticktime int
	defaultsay      string
}

func (h Hello) Say(ch chan string) {

	ticktime := h.defaultticktime
	saythis := h.defaultsay

	ticker := time.Tick(time.Millisecond * time.Duration(ticktime))
	count := 0

	// Loop forever - Long running
	for {
		select {
		case now := <-ticker:
			if saythis == "quit" {
				break
			}
			fmt.Printf("%v %d %s\n", now, count, saythis)
		case bleh := <-ch:
			saythis = bleh
		}
		count++
	}
}

func main() {

	hi1 := Hello{defaultticktime: 1000, defaultsay: "Waiting1"}
	hi2 := Hello{defaultticktime: 500, defaultsay: "Waiting2"}

	ch1 := make(chan string)
	ch2 := make(chan string)

	// A goroutine calling a method with a channel for communication
	go hi1.Say(ch1)
	go hi2.Say(ch2)

	time.Sleep(5 * time.Second)
	ch1 <- "Jeff"
	ch2 <- "Clif"

	time.Sleep(15 * time.Second)
	ch1 <- "Jack"
	ch2 <- "Jill"

	time.Sleep(4 * time.Second)
	ch1 <- "quit"
	time.Sleep(4 * time.Second)
	ch2 <- "quit"
}
