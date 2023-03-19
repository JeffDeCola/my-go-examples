package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {

	// ls -s
	out, err := exec.Command("ls", "-l").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))

}
