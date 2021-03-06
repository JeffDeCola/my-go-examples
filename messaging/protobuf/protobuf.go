package main

import (
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
)

// Check your error
func checkErr(err error) {
	if err != nil {
		log.Fatal("ERROR:", err)
	}
}

func main() {

	// DATA
	sndPerson := &Person{
		Name:  "Jeff",
		Age:   20,
		Email: "blah@blah.com",
		Phone: "555-555-5555",
		Count: 1,
	}

	// MARSHAL
	sndMsg, err := proto.Marshal(sndPerson)
	checkErr(err)

	// SEND
	pipe := sndMsg

	// RECEIVE
	rcvMsg := pipe

	// UNMARSHAL -> DATA
	rcvPerson := &Person{}
	err = proto.Unmarshal(rcvMsg, rcvPerson)
	checkErr(err)

	// CHECK EVERYTHING WORKED
	log.Printf("msg sent:     %+v", sndPerson)
	log.Printf("msg received: %+v", rcvPerson)
	if sndPerson.Name != rcvPerson.Name {
		log.Fatalf("data mismatch %q != %q", sndPerson, rcvPerson)
	}

}
