package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"
)

const numberSubscribers = 10

// Check your error
func checkErr(err error) {
	if err != nil {
		log.Fatal("ERROR:", err)
	}
}

func subscriber(nc *nats.Conn, ch chan *nats.Msg, workerID int) {

	// Pick off channel (Only one subscriber will grab the msg off the channel. This is 1-1 communication)
	for msg := range ch {

		// UNMARSHAL -> DATA
		rcvPerson := &Person{}
		err := proto.Unmarshal(msg.Data, rcvPerson)
		checkErr(err)

		log.Printf("[%v] Person received: %+v\n", workerID, rcvPerson)

		// REPLY
		myReply := &MyReply{}
		myReply.Thereply = fmt.Sprintf("[%v] This is a response for count %d", workerID, rcvPerson.Count)

		// MARSHAL
		replymsg, err := proto.Marshal(myReply)
		checkErr(err)

		// SEND
		// NATS - PUBLISH on "foo" (THE PIPE)
		log.Printf("[%v] Publishing replymsg (%v) to subject 'foo'\n", workerID, myReply.Thereply)
		err = nc.Publish(msg.Reply, replymsg)
		checkErr(err)
	}
}

func main() {

	// CONNECT TO NATS (nats-server)
	nc, err := nats.Connect("nats://127.0.0.1:4222")
	checkErr(err)
	defer nc.Close()
	log.Println("Connected to " + nats.DefaultURL)

	// Create a Channel
	ch := make(chan *nats.Msg)

	// Create Workers
	for i := 0; i < numberSubscribers; i++ {
		go subscriber(nc, ch, i)
	}

	// SUBSCRIBE TO "foo"
	nc.Subscribe("foo", func(m *nats.Msg) {
		ch <- m
	})

	// wait - empty select
	select {}

}
