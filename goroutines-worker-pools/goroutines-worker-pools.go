package main

import (
	"fmt"

	"time"
)

var instanceList = []string{"i1", "i2", "i3", "i4", "i5"}
var metricList = []string{"m1", "m2", "m3", "m4", "m5"}
var ticktimeSeconds = 10
var sampletime = 11
var numWorkers = 3
var ticker = time.Tick(time.Duration(ticktimeSeconds) * time.Second)
var jobNumber = 1
var jobTime time.Duration = 7
var completedJobs = 0

type request struct {
	instance  string
	metric    string
	callTime  time.Time
	jobNumber int
}

func doWorkers(id int, ch chan *request) {

	fmt.Printf("W%d - START Worker\n", id)

	// Monitor channel for request struct
	for r := range ch {
		time.Sleep(jobTime * time.Second)
		fmt.Printf("W%d - RESPONSE #%d - %s, %s - %s (TOOK %s)\n", id, r.jobNumber, r.instance, r.metric, time.Now().UTC().Format(time.ANSIC), time.Now().UTC().Sub(r.callTime))
		completedJobs++
	}

}

func doSomething(ch chan *request) {

	endTime := time.Now().UTC()

	// Loop forever - Long Running
	for {
		select {
		case now := <-ticker:

			fmt.Printf("\n\n********************** TICK %s *****************************\n\n", now.UTC().Format(time.ANSIC))

			oldendTime := endTime
			endTime = time.Now().UTC()
			startTime := endTime.Add(-(time.Duration(sampletime) * time.Second))

			// CHECK TICKER IF YOU MADE IT BACK IN TIME

			diff := startTime.Sub(oldendTime)
			fmt.Printf("DIFF: %s (oldendTime: %s - starttime %s)\n", diff, startTime.Format(time.ANSIC), oldendTime.Format(time.ANSIC))
			if diff.Seconds() > 1 {
				fmt.Printf("\n\n\n************* ERROR - YOU MISSED SOME TIME - Not enough time to process *****************\n")
				fmt.Printf("    JOBS NOT FINISHED: %d\n", (jobNumber-1)-completedJobs)
				fmt.Printf("    SKIP A TICK TO CATCH UP\n\n")
			} else {
				fmt.Printf("Perfect - You did not miss any time\n")
				fmt.Printf("JOBS NOT FINISHED YET: %d\n", (jobNumber-1)-completedJobs)

				// Interate over
				for _, instance := range instanceList {

					for _, metric := range metricList {

						// BLOCKS - WAITS FOR WORKER
						ch <- &request{
							instance:  instance,
							metric:    metric,
							callTime:  time.Now().UTC(),
							jobNumber: jobNumber,
						}
						fmt.Printf("CALL #%d - %s %s %s\n", jobNumber, instance, metric, time.Now().UTC().Format(time.ANSIC))
						jobNumber++
					}
				}
			}
		}
	}
}

func main() {

	ch := make(chan *request)

	// CREATE WORKERS
	for i := 0; i < numWorkers; i++ {
		go doWorkers(i, ch)
	}

	doSomething(ch)

}
