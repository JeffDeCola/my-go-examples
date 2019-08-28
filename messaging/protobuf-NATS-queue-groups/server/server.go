package main

import (
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"
)

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
	log.Println("Connected to " + nats.DefaultURL)

	// SUBSCRIBE TO "foo"
	sub, err := nc.SubscribeSync("foo")
	checkErr(err)

	// Loop forever - Long Running
	for {

		// RECEIVE
		log.Printf("   Receiving msg from subject 'foo'\n")
		msg, err := sub.NextMsg(time.Duration(5) * time.Second)
		checkErr(err)

		// UNMARSHAL -> DATA
		rcvPerson := &Person{}
		err = proto.Unmarshal(msg.Data, rcvPerson)
		checkErr(err)

		log.Printf("Person received: %+v", rcvPerson)

	}
}
