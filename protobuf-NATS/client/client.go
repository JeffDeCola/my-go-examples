package main

import (
	"log"
	"time"

	nats "github.com/go-nats"
	"github.com/golang/protobuf/proto"
)

func main() {

	token := &Token{
		AccessToken:  "the access token",
		TokenType:    "this",
		RefreshToken: "and the refresh token",
		ExpiresAt:    5,
		Counter:      5,
	}

	nc, _ := nats.Connect("nats://127.0.0.1:4222")
	log.Println("Connected to " + nats.DefaultURL)

	var counter int64
	counter = 1

	// Loop forever - Long Running
	for c := time.Tick(time.Duration(2) * time.Second); ; <-c {

		log.Printf("Count %d\n", counter)
		token.Counter = counter

		// PROTOBUF - CLIENT - MARSHAL - WRITE/SEND
		msg, err := proto.Marshal(token)
		if err != nil {
			log.Fatal("marshaling error: ", err)
		}

		log.Printf("   token sending :    %+v", token)

		// NATS - PUBLISH on "foo"
		log.Printf("   Publishing to subject 'foo'\n")
		nc.Publish("foo", msg)
		// nc.Publish("foo", []byte(strconv.Itoa(counter)))

		counter++
	}
}
