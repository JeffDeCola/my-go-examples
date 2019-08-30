package main

import (
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"
)

const tickTime = 1       // In seconds
const replyWaitTime = 20 // In seconds

// Check your error
func checkErr(err error) {
	if err != nil {
		log.Fatal("ERROR:", err)
	}
}

func main() {

	// CONNECT TO NATS (nats-server)
	nc, err := nats.Connect("nats://127.0.0.1:4222")
	checkErr(err)
	defer nc.Close()
	log.Println("Connected to " + nats.DefaultURL)

	// DATA
	sndPerson := &Person{
		Name:  "Jeff",
		Age:   20,
		Email: "blah@blah.com",
		Phone: "555-555-5555",
		Count: 1,
	}

	var counter uint32
	counter = 1

	// Loop forever - Long Running
	for c := time.Tick(time.Duration(tickTime) * time.Second); ; <-c {

		log.Printf("Count is %d\n", counter)
		sndPerson.Count = counter

		// MARSHAL
		msg, err := proto.Marshal(sndPerson)
		checkErr(err)

		// SEND
		// NATS - REQUEST & REPLY on "foo" (THE PIPE)
		log.Printf("   Send request msg to subject 'foo'\n")
		reply, err := nc.Request("foo", msg, replyWaitTime*time.Second)
		checkErr(err)

		// UNMARSHAL -> DATA
		myReply := &MyReply{}
		err = proto.Unmarshal(reply.Data, myReply)
		checkErr(err)
		log.Printf("   Reply received: %v\n", myReply.Thereply)

		counter++
	}
}
