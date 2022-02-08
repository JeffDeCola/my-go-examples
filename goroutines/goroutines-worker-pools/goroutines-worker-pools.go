// my-go-examples goroutines-worker-pools.go

package main

import (
	"fmt"

	"time"
)

// FOR doSomething()
var jobTaskList = []string{"taskA", "taskB", "taskC", "taskD", "taskE"} // 5 tasks
const SendNewJobTime = 10                                               // Send a job every x seconds

// FOR WORKERS - doTask()
const numberWorkers = 2 // How many workers in pool
const taskTime = 7      // How long it takes a worker to complete a task

// CHANNEL
var channelBufferSize = 2

type jobInfo struct {
	jobNumber int
	taskName  string
	sentTime  time.Time
}

func doTask(id int, msgCh chan jobInfo) {

	fmt.Printf("                                    WORKER%3d  - STARTED\n", id)

	// GRAB JOB FROM CHANNEL
	for job := range msgCh {

		fmt.Printf("                                    WORKER%2d - STARTED  JOB%2d (%s) at %s\n", id, job.jobNumber, job.taskName, time.Now().UTC().Format("03:04:05"))
		startTime := time.Now().UTC()

		// DO WORK ON TASK
		time.Sleep(taskTime * time.Second)

		// FINISHED
		diff := time.Now().UTC().Sub(startTime)
		fmt.Printf("                                    WORKER%2d - FINISHED JOB%2d (%s) at %s (TOOK %s)\n", id, job.jobNumber, job.taskName, time.Now().UTC().Format("03:04:05"), diff)

	}

}

func doSomething(msgCh chan jobInfo) {

	var jobNumber = 1 // The start job number

	endTime := time.Now().UTC().Add(-(time.Duration(SendNewJobTime) * time.Second)) // Substract (SendNewJobTime) from endTime

	for c := time.Tick(time.Duration(SendNewJobTime) * time.Second); ; <-c {

		oldendTime := endTime
		endTime = time.Now().UTC()
		startTime := endTime.Add(-(time.Duration(SendNewJobTime) * time.Second)) // Substract (SendNewJobTime) from endTime

		// CHECK IF YOU MADE IT BACK IN TIME (WITHIN A SECOND)
		diff := startTime.Sub(oldendTime)
		if diff.Seconds() > 1 {
			fmt.Printf("\n*** WARNING - JOBS BACKING UP *** \n")
			fmt.Printf("*** You are late sending a new job by %s seconds ***\n", diff)
			fmt.Printf("*** Increase workers, increase job time, reduce tasks or speed up workers ***\n\n")
		} else {
			fmt.Printf("\nPERFECT - Still sending jobs\n\n")
		}

		// SEND TASK TO BUFFER CHANNEL
		for _, taskName := range jobTaskList {

			t := time.Now().UTC()
			fmt.Printf("READY TO SEND TASK  JOB%2d (%s) %s\n", jobNumber, taskName, t.Format("03:04:05"))

			// WAITS/BLOCKS IF BUFFER IS FULL
			msgCh <- jobInfo{
				jobNumber: jobNumber,
				taskName:  taskName,
				sentTime:  t,
			}

			fmt.Printf("SENT TASK           JOB%2d (%s) %s\n", jobNumber, taskName, time.Now().UTC().Format("03:04:05"))

		}

		jobNumber++

	}
}

func main() {

	// MAKE CHANNEL BUFFER
	msgCh := make(chan jobInfo, channelBufferSize)
	fmt.Println(" ")

	// CREATE WORKER POOLS
	for id := 1; id < numberWorkers+1; id++ {
		go doTask(id, msgCh)
	}

	// DO SOMETHING (LOOP FOREVER)
	go doSomething(msgCh)

	// PRESS RETURN TO EXIT
	fmt.Scanln()
	fmt.Println("done")

}
