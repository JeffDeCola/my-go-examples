// my-go-examples json.go

package main

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

type myData struct {
	Name string
	Body string
	Time int64
}

func encodeJSON() (string, error) {

	m := myData{"Alice", "Hello", 123}
	log.Info("  struct:  ", m)

	// ENCODE - STRUCT TO JSON
	log.Info("  encode:   Marshal")
	b, err := json.Marshal(m)
	if err != nil {
		jsondata := "poop"
		return jsondata, err
	}
	log.Info("  byte:    ", b)
	jsondata := string(b)
	log.Info("  string:  ", jsondata)
	return jsondata, err
}

func decodeJSON(jsondata string) error {

	log.Info("  string:  ", jsondata)

	// DECODE - JSON TO STRUCT
	b := []byte(jsondata)
	log.Info("  byte:    ", b)

	foo := myData{} // empty struct
	log.Info("  decode:   Unmarshal")
	err := json.Unmarshal(b, &foo)
	if err != nil {
		return err
	}
	log.Info("  struct:  ", foo)

	return nil
}

func main() {

	// SET LOG LEVEL
	// log.SetLevel(log.InfoLevel)
	log.SetLevel(log.TraceLevel)

	// SET FORMAT
	log.SetFormatter(&log.TextFormatter{})

	// SET OUTPUT (DEFAULT stderr)
	// log.SetOutput(os.Stdout)

	var jsondata string

	log.Info("ENCODE - STRUCT TO JSON")
	jsondata, err := encodeJSON()
	if err != nil {
		log.Fatal("Failed to encode to stdout: %s", err)
	}

	log.Info("DECODE - JSON TO STRUCT")
	if err := decodeJSON(jsondata); err != nil {
		log.Fatal("Failed to decode to stdout: %s", err)
	}

}
