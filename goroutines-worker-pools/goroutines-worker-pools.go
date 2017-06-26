package main

import (
	"fmt"

	"time"
)

type request struct {
	instance string
	metric   string
	callTime time.Time
	count    int
}

func doWorkers(id int, ch chan *request) {

	fmt.Printf("W%d - START Worker\n", id)

	for r := range ch {
		time.Sleep(18 * time.Second)
		fmt.Printf("W%d - RESPONSE #%d - %s, %s - TOOK %s\n", id, r.count, r.instance, r.metric, time.Now().UTC().Sub(r.callTime))
	}

	// fmt.Printf("YOU WIL NEVER SEE THIS\n")

}

func main() {

	var instanceList = []string{"a"}
	var metricList = []string{"m1", "m2", "m3", "m4"}
	var ticktimeSeconds = 10

	ticker := time.Tick(time.Duration(ticktimeSeconds) * time.Second)
	sampletime := 20
	ch := make(chan *request)
	count := 1

	// CREATE WORKERS
	for i := 0; i < 3; i++ {
		go doWorkers(i, ch)
	}

	time.Sleep(3 * time.Second)

	ranThrough := false
	endTime := time.Now().UTC()

	// Loop forever - Long Running
	for {
		select {
		case now := <-ticker:

			fmt.Printf("\n\n********************** TICK %s *****************************\n\n", now.UTC().Format(time.ANSIC))

			oldendTime := endTime
			endTime = now.UTC()
			startTime := endTime.Add(-(time.Duration(sampletime) * time.Second))

			// CHECK TICKER IF YOU MADE IT BACK IN TIME
			if ranThrough {

				diff := startTime.Sub(oldendTime)
				fmt.Printf("DIFF: %s (oldendTime: %s - starttime %s)\n", diff, startTime.Format(time.ANSIC), oldendTime.Format(time.ANSIC))
				if diff.Seconds() > 1 {
					fmt.Printf("ERROR - YOU MISSED SOME TIME - Not enough time to process\n")
				} else {
					fmt.Printf("Perfect - You did not miss any time\n")
				}

			}

			// Interate over
			for _, instance := range instanceList {

				for _, metric := range metricList {

					fmt.Printf("CALL #%d - %s %s %s\n", count, instance, metric, now.Format(time.ANSIC))

					ch <- &request{
						instance: instance,
						metric:   metric,
						callTime: time.Now().UTC(),
						count:    count,
					}
					count++
				}
			}
			ranThrough = true
		}
	}
	// fmt.Println("YOU WILL NEVER SEE THIS - All done!")
}
