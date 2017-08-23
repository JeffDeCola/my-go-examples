package main

import (
	"fmt"
	"log"

	nats "github.com/go-nats"
	"github.com/protobuf/proto"
)

func workerServer(nc *nats.Conn, ch chan *nats.Msg, workerID int) {

	// Pick off channel
	for msg := range ch {

		fmt.Printf("[%d] got request: %s\n", workerID, string(msg.Data))

		// PROTOBUF - SERVER - RECEIVE - READ/UNMARSHAL
		rcvToken := &Token{}
		err := proto.Unmarshal(msg.Data, rcvToken)
		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}

		log.Printf("%d - Count %d - Token received", workerID, rcvToken.Counter)
		// log.Printf("    AccessToken: %+v", rcvToken.AccessToken)
		//log.Printf("    TokenType: %+v", rcvToken.TokenType)
		//log.Printf("    RefreshToken: %+v", rcvToken.RefreshToken)
		//log.Printf("    ExpiresAt: %+v", rcvToken.ExpiresAt)
		//log.Printf("    Counter: %+v", rcvToken.Counter)

		tokenResponse := &TokenResponse{}
		tokenResponse.MyReply = fmt.Sprintf("This is a response, workerID=%d from count %d", workerID, rcvToken.Counter)

		// PROTOBUF - CLIENT - MARSHAL - WRITE/SEND
		msgreply, err := proto.Marshal(tokenResponse)
		if err != nil {
			log.Fatal("marshaling error: ", err)
		}

		log.Printf("    tokenResponse sending : %+v", tokenResponse)

		nc.Publish(msg.Reply, msgreply)

	}

}

func main() {

	// NATS - SUBSCRIBE ON "foo"
	nc, _ := nats.Connect("nats://127.0.0.1:4222")
	defer nc.Close()
	log.Println("Connected to " + nats.DefaultURL)

	// Create a Channel
	ch := make(chan *nats.Msg)

	// Create Workers
	for i := 0; i < 10; i++ {
		go workerServer(nc, ch, i)
	}

	// Subscribe on NATS
	nc.Subscribe("foo", func(m *nats.Msg) {
		ch <- m
	})

	// wait - empty select
	select {}

}
