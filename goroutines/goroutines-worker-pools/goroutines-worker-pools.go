package main

import (
	"fmt"

	"time"
)

// SET YOUR VARIABLES FOR THIS EXAMPLE
var numWorkers = 2
var instanceList = []string{"i1", "i2", "i3", "i4", "i5"}
var ticktimeSeconds = 10
var sampletime = 11
var jobTime time.Duration = 7
var channelBufferSize = 2

// Other Variables
var jobNumber = 1

type request struct {
	instance  string
	metric    string
	callTime  time.Time
	jobNumber int
}

// doWorkers is the goroutine Worker
func doWorkers(id int, reqch chan *request) {

	fmt.Printf("W%d - START Worker\n", id)

	// Monitor channel for request channel
	for r := range reqch {
		time.Sleep(jobTime * time.Second)
		fmt.Printf("W%d - RESPONSE #%d - %s, %s - %s (TOOK %s)\n", id, r.jobNumber, r.instance, r.metric, time.Now().UTC().Format(time.ANSIC), time.Now().UTC().Sub(r.callTime))
	}

}

func doSomething(requestch chan *request) {

	endTime := time.Now().UTC()

	// Loop forever - Long Running (Will start immediately then wait for tick)
	for c := time.Tick(time.Duration(ticktimeSeconds) * time.Second); ; <-c {

		fmt.Printf("\n\n*** TICK %s ***\n\n\n", time.Now().UTC().Format(time.ANSIC))

		oldendTime := endTime
		endTime = time.Now().UTC()
		startTime := endTime.Add(-(time.Duration(sampletime) * time.Second))

		// CHECK TICKER IF YOU MADE IT BACK IN TIME
		diff := startTime.Sub(oldendTime)
		fmt.Printf("DIFF: %s (oldendTime: %s - starttime %s)\n", diff, startTime.Format(time.ANSIC), oldendTime.Format(time.ANSIC))
		if diff.Seconds() > 1 {
			fmt.Printf("*** ERROR - YOU MISSED SOME SAMPLE TIME *** \n")
			fmt.Printf("*** Not enough time to process all instances within tick ***\n")
			fmt.Printf("*** Increase workers or reduce number of instances ***\n\n")
		} else {
			fmt.Printf("Perfect - You did not miss any time\n")
		}

		// Interate over
		for _, instance := range instanceList {

			fmt.Printf("CALLING #%d - %s %s\n", jobNumber, instance, time.Now().UTC().Format(time.ANSIC))

			// WAITS FOR WORKER
			requestch <- &request{
				instance:  instance,
				callTime:  time.Now().UTC(),
				jobNumber: jobNumber,
			}

			fmt.Printf("CALLED #%d - %s %s\n", jobNumber, instance, time.Now().UTC().Format(time.ANSIC))
			jobNumber++

		}
	}
}

func main() {

	// Make Request Channel with Buffer
	reqch := make(chan *request, channelBufferSize)

	// Create Workers
	for i := 0; i < numWorkers; i++ {
		go doWorkers(i, reqch)
	}

	doSomething(reqch)

}
