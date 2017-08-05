package main

import (
	"log"
	"time"

	nats "github.com/go-nats"
	"github.com/protobuf/proto"
)

func main() {

	// NATS - SUBSCRIBE ON "foo"
	nc, _ := nats.Connect("nats://127.0.0.1:4222")
	log.Println("Connected to " + nats.DefaultURL)
	log.Printf("Subscribing to subject 'foo'\n")
	// Synchronous way - When i want to check for one msg at a time
	sub, err := nc.SubscribeSync("foo")
	if err != nil {
		log.Fatal("SubscribeSync error: ", err)
	}

	// Loop forever - Long Running
	for {

		msg, err := sub.NextMsg(time.Duration(5) * time.Second)
		if err != nil {
			log.Fatal("next message error: ", err)
		}
		// fmt.Printf("Received a NAT message: %v\n", msg)

		// PROTOBUF - SERVER - RECEIVE - READ/UNMARSHAL
		rcvToken := &Token{}
		err = proto.Unmarshal(msg.Data, rcvToken)
		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}

		log.Printf("Token received: %+v", rcvToken)
		log.Printf("    AccessToken: %+v", rcvToken.AccessToken)
		log.Printf("    TokenType: %+v", rcvToken.TokenType)
		log.Printf("    RefreshToken: %+v", rcvToken.RefreshToken)
		log.Printf("    ExpiresAt: %+v", rcvToken.ExpiresAt)
		log.Printf("    Counter: %+v", rcvToken.Counter)

	}
}
