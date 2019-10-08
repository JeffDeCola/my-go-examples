// my-go-examples simple-tcp-server handlers.go

package main

import (
	"bufio"
	"strconv"
	"strings"
)

func handleAdd(rw *bufio.ReadWriter) {

	s := "Please enter two integers"
	returnMessage(s, rw)

	// WAITING FOR INT 1
	s1, err := rw.ReadString('\n')
	checkErr(err)
	s1 = strings.Trim(s1, "\n ")
	i1, _ := strconv.Atoi(s1)
	s = "Received first integer: " + s1
	returnMessage(s, rw)

	// WAITING FOR INT 2
	s2, err := rw.ReadString('\n')
	checkErr(err)
	s2 = strings.Trim(s2, "\n ")
	i2, _ := strconv.Atoi(s2)
	s = "Received second integer: " + s2
	returnMessage(s, rw)

	// THE SUM
	sum := i1 + i2
	s = "The sum of " + s1 + " + " + s2 + " = " + strconv.Itoa(sum)
	returnMessage(s, rw)

}

func handleSubstract(rw *bufio.ReadWriter) {

	s := "Please enter two integers"
	returnMessage(s, rw)

	// WAITING FOR INT 1
	s1, err := rw.ReadString('\n')
	checkErr(err)
	s1 = strings.Trim(s1, "\n ")
	i1, _ := strconv.Atoi(s1)
	s = "Received first integer: " + s1
	returnMessage(s, rw)

	// WAITING FOR INT 2
	s2, err := rw.ReadString('\n')
	checkErr(err)
	s2 = strings.Trim(s2, "\n ")
	i2, _ := strconv.Atoi(s2)
	s = "Received second integer: " + s2
	returnMessage(s, rw)

	// SUBTRACT
	sum := i1 - i2
	s = "Subtracting " + s1 + " - " + s2 + " = " + strconv.Itoa(sum)
	returnMessage(s, rw)

}
