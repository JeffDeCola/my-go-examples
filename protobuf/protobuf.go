package main

import (
	"log"

	"github.com/golang/protobuf/proto"
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

	// SERVER - RECEIVE - READ/UNMARSHAL
	rcvToken := &Token{}
	err = proto.Unmarshal(msg, rcvToken)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	log.Printf("token sent :    %+v", token)
	log.Printf("token received: %+v", rcvToken)

	// TEST SAME DATA
	if token.AccessToken != rcvToken.AccessToken {
		log.Fatalf("data mismatch %q != %q", token, rcvToken)
	}

}
