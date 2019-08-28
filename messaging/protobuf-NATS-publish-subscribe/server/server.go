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

	// NATS - SUBSCRIBE ON "foo"
	nc, _ := nats.Connect("nats://127.0.0.1:4222")
	log.Println("Connected to " + nats.DefaultURL)
	log.Printf("Subscribing to subject 'foo'\n")
	// Synchronous way - When i want to check for one msg at a time
	sub, err := nc.SubscribeSync("foo")
	checkErr(err)

	// Loop forever - Long Running
	for {

		msg, err := sub.NextMsg(time.Duration(5) * time.Second)
		checkErr(err)

		// PROTOBUF - SERVER - RECEIVE - READ/UNMARSHAL
		rcvToken := &Person{}
		err = proto.Unmarshal(msg.Data, rcvToken)
		checkErr(err)

		log.Printf("Person received: %+v", rcvToken)
		//log.Printf("    AccessToken: %+v", rcvToken.AccessToken)
		//log.Printf("    TokenType: %+v", rcvToken.TokenType)
		//log.Printf("    RefreshToken: %+v", rcvToken.RefreshToken)
		//log.Printf("    ExpiresAt: %+v", rcvToken.ExpiresAt)
		//log.Printf("    Counter: %+v", rcvToken.Counter)

	}
}
