package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/protobuf/proto"
)

func testSendReceive() {

	person := &Person{
		Name:  "Jeff",
		Age:   20,
		Email: "blah@blah.com",
		Phone: "555-555-5555",
	}

	// CLIENT - MARSHAL - WRITE/SEND
	msg, err := proto.Marshal(person)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	// SERVER - RECEIVE - READ/UNMARSHAL
	rcvPerson := &Person{}
	err = proto.Unmarshal(msg, rcvPerson)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	log.Printf("person sent :    %+v", person)
	log.Printf("person received: %+v", rcvPerson)

	// TEST SAME DATA
	if person.Name != rcvPerson.Name {
		log.Fatalf("data mismatch %q != %q", person, rcvPerson)
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

	println("Lets test sending a message and receiving a message")

	testSendReceive()

	// testData()
}
