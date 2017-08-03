package main

import (
	"log"

	"github.com/golang/protobuf/proto"
	// "github.com/nats-io/nats"
)

func main() {

	// SERVER - RECEIVE - READ/UNMARSHAL
	rcvToken := &Token{}
	err = proto.Unmarshal(msg, rcvToken)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	log.Printf("token received: %+v", rcvToken)

}
