// my-go-examples goroutines-worker-pools.go

package main

import (
	"fmt"

	"time"
)

// SET CONSTANTS FOR WORKER doWork()
const numberWorkers = 2 // How many workers you want
const workerTime = 7    // How long it takes a worker to work

// SET CONSTANTS FOR doSomething()
var jobList = []string{"jobA", "jobB", "jobC", "jobD", "jobE"} // 5 jobs with jobNames
var ticktimeSeconds = 10                                       // Tick time to send a bunch of jobs workers
var jobNumber = 1                                              // The job number

// OTHER
var channelBufferSize = 1 // How many channel buffers

type request struct {
	jobNumber   int
	jobName     string
	requestTime time.Time
}

// doWork grabs jobs from the channel
func doWork(id int, msgCh chan *request) {

	fmt.Printf("                                    WORKER%3d  - STARTED\n", id)

	// Monitor channel for job
	for r := range msgCh {
		fmt.Printf("                                    WORKER%3d - STARTED  %4d - %s (%s)\n", id, r.jobNumber, r.jobName, time.Now().UTC().Format("03:04:05.000"))
		time.Sleep(workerTime * time.Second)
		diff := time.Now().UTC().Sub(r.requestTime)
		fmt.Printf("                                    WORKER%3d - FINISHED %4d - %s (TOOK %s)", id, r.jobNumber, r.jobName, diff)
		if diff.Seconds() > workerTime+1 {
			fmt.Printf(" Taking a bit longer (pulled from buffer)")
		}
		fmt.Printf("\n")
	}

}

func doSomething(msgCh chan *request) {

	endTime := time.Now().UTC().Add(-(time.Duration(ticktimeSeconds) * time.Second)) // Substract (ticktimeSeconds) from endTime

	// Loop forever - Long Running (Will start immediately then wait for tick)
	for c := time.Tick(time.Duration(ticktimeSeconds) * time.Second); ; <-c {

		fmt.Printf("\n\n*** TICK (%s) ***\n", time.Now().UTC().Format("03:04:05.000"))

		oldendTime := endTime
		endTime = time.Now().UTC()
		startTime := endTime.Add(-(time.Duration(ticktimeSeconds) * time.Second)) // Substract (ticktimeSeconds) from endTime

		// CHECK TICKER IF YOU MADE IT BACK IN TIME (WITHIN A SECOND)
		diff := startTime.Sub(oldendTime)
		fmt.Printf("DIFF: %s\n", diff)
		// fmt.Printf("DIFF: %s (starttime: %s - oldendTime %s)\n", diff, startTime.Format("03:04:05.000000"), oldendTime.Format("03:04:05.000000"))
		if diff.Seconds() > 1 {
			fmt.Printf("*** ERROR - YOU MISSED SOME TICK TIME *** \n")
			fmt.Printf("*** Not enough time to process all jobs within tick ***\n")
			fmt.Printf("*** Increase workers, reduce number of jobs ***\n\n")
		} else {
			fmt.Printf("PERFECT - You did not miss a tick\n\n")
		}

		// Interate over jobList and send to worker
		for _, jobName := range jobList {

			t := time.Now().UTC()
			fmt.Printf("SENDING %4d - %s (%s)\n", jobNumber, jobName, t.Format("03:04:05.000"))

			// WAITS FOR WORKER
			msgCh <- &request{
				jobNumber:   jobNumber,
				jobName:     jobName,
				requestTime: t,
			}

			fmt.Printf("SENT    %4d - %s (%s)\n", jobNumber, jobName, time.Now().UTC().Format("03:04:05.000"))
			jobNumber++

		}
	}
}

func main() {

	// Make Channel Buffer
	msgCh := make(chan *request, channelBufferSize)

	fmt.Println(" ")

	// Create Workers
	for id := 1; id < numberWorkers+1; id++ {
		go doWork(id, msgCh)
	}

	// Will loop forever
	go doSomething(msgCh)

	// Press return to exit
	fmt.Scanln()
	fmt.Println("done")
}
