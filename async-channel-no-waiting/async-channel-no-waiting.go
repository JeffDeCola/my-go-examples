package main

import (
	"fmt"
	"strconv"
	"time"
)

var getInstanceSpeedSeconds = 15
var usingInstancesSpeedSeconds = 5

// getInstances - Returns a new instance list every 5 seconds
func getInstances(chRes chan []string) {

	var instanceList []string
	counter := 1

	// Loop forever - Long Running (Will start immediately then wait for tick)
	for c := time.Tick(time.Duration(getInstanceSpeedSeconds) * time.Second); ; <-c {

		// Get Instances
		fmt.Printf("    #%d - Getting Instances\n", counter)
		instanceList = append(instanceList, "I"+strconv.Itoa(counter))

		// Send Instances to Channel
		chRes <- instanceList
		fmt.Printf("    #%d - Sent Instances to Channel: %s\n", counter, instanceList)

		counter++
	}
}

// usingInstances
func usingInstances(instanceListCh chan []string) {

	var instanceList []string
	counter := 1

	// Loop forever - Long Running (Will start immediately then wait for tick)
	for c := time.Tick(time.Duration(usingInstancesSpeedSeconds) * time.Second); ; <-c {

		fmt.Printf("*** #%d - START USING INSTANCES\n", counter)

		// Get the current instance list (from channel)
		// If there is nothing in channel:
		//       - default and break out of loop
		// If there is something in channel:
		//       - read and continue until empty channel (hence, get latest)
		//       - break out of loop like above
		for {
			select {
			case newVal := <-instanceListCh:
				instanceList = newVal
				continue
			default:
			}
			break
		}

		fmt.Printf("*** #%d - Using Instances:           %s\n", counter, instanceList)
		counter++
	}
}

func main() {
	var ch = make(chan []string, 100)

	// kick off gorountine
	go getInstances(ch)

	usingInstances(ch)

}
