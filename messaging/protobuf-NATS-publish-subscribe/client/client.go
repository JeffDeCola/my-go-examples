package main

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"
)

// Check your error
func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error is %+v\n", err)
		log.Fatal("ERROR:", err)
	}
}

func main() {

	sndPerson := &Person{
		Name:  "Jeff",
		Age:   20,
		Email: "blah@blah.com",
		Phone: "555-555-5555",
		Count: 1,
	}

	// CONNECT TO NATS (nats-server)
	nc, err := nats.Connect("nats://127.0.0.1:4222")
	checkErr(err)
	log.Println("Connected to " + nats.DefaultURL)

	var counter int32
	counter = 1

	// Loop forever - Long Running
	for c := time.Tick(time.Duration(2) * time.Second); ; <-c {

		log.Printf("Count is %d\n", counter)
		sndPerson.Count = counter

		// PROTOBUF -> MARSHAL
		msg, err := proto.Marshal(sndPerson)
		checkErr(err)

		// SEND - NATS - PUBLISH on "foo" (THE PIPE)
		log.Printf("   Publishing to subject 'foo'\n")
		nc.Publish("foo", msg)

		counter++
	}
}
