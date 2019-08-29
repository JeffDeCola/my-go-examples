package main

import (
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
	defer nc.Close()
	log.Println("Connected to " + nats.DefaultURL)

	// Loop forever - Long Running
	for {

		// RECEIVE
		nc.QueueSubscribe("foo", "jeffsQueue", func(msg *nats.Msg) {

			// UNMARSHAL -> DATA
			rcvPerson := &Person{}
			err = proto.Unmarshal(msg.Data, rcvPerson)
			checkErr(err)

			log.Printf("Person received: %+v", rcvPerson)
		})

		// wait - empty select
		select {}

	}
}
