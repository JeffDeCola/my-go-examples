package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getJSON1(url string, target *Foo) {

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	target.Bar = string(contents)
	// fmt.Printf("Content is:\n\n%s\n", string(contents))

	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(target)
}

type Foo struct {
	Bar string
}

func main() {

	// url := "http://www.jeffdecola.com/robots.txt"
	url := "https://www.googleapis.com/storage/v1/b"

	foo1 := Foo{}
	getJSON1(url, &foo1)
	fmt.Printf("Content is:\n\n%s\n", foo1.Bar)

}
