package main

import (
	"encoding/json"
	"log"
	"os"
)

func encodejson(logger *log.Logger) (string, error) {
	type blah struct {
		Name string
		Body string
		Time int64
	}

	m := blah{"Alice", "Hello", 123}
	logger.Println(m)

	// ENCODE - CREATE JSON
	b, err := json.Marshal(m)
	if err != nil {
		jsondata := "poop"
		return jsondata, err
	}
	jsondata := string(b)
	logger.Println(jsondata)
	return jsondata, err
}

func decodejson(jsondata string, logger *log.Logger) error {
	type m struct {
		Name string
		Body string
		Time int64
	}

	logger.Println(jsondata)

	// DECODE - CREATE JSON
	b := []byte(jsondata)

	foo := m{}
	err := json.Unmarshal(b, &foo)
	if err != nil {
		return err
	}
	logger.Println(foo)
	return nil
}

func main() {

	var logger = log.New(os.Stderr, "jeffmain:", log.Lshortfile)
	var jsondata string

	logger.Print("ENCODE .json")
	jsondata, err := encodejson(logger)
	if err != nil {
		logger.Fatalf("Failed to encode to stdout: %s", err)
	}
	logger.Println(jsondata)

	logger.Print("DECODE .json")
	if err := decodejson(jsondata, logger); err != nil {
		logger.Fatalf("Failed to decode to stdout: %s", err)
	}

}
