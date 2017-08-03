package main

import (
	"log"

	"github.com/golang/protobuf/proto"
	//"github.com/nats-io/nats"
)

func main() {

	token := &Token{
		AccessToken:  "the access token",
		TokenType:    "this",
		RefreshToken: "and the refresh token",
		ExpiresAt:    5,
	}

	// CLIENT - MARSHAL - WRITE/SEND
	msg, err := proto.Marshal(token)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	log.Printf("token sent :    %+v", token)

}
