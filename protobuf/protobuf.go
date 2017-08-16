package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/protobuf/proto"
)

func testToken() {

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

// Read Large data file - Is there a limit
func testData() {

	// Read the secrets file from google
	raw, err := ioutil.ReadFile("in.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	data := &Data{
		Body: string(raw),
	}

	// CLIENT - MARSHAL - WRITE/SEND
	msg, err := proto.Marshal(data)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	// SERVER - RECEIVE - READ/UNMARSHAL
	rcvData := &Data{}
	err = proto.Unmarshal(msg, rcvData)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	// Write file
	err = ioutil.WriteFile("out.txt", []byte(rcvData.Body), 0644)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

}

func main() {
	testToken()
	testData()
}
