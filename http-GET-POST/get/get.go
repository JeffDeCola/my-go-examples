package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	// url := "http://www.jeffdecola.com/robots.txt"
	url := "https://www.googleapis.com/storage/v1/b"

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	fmt.Printf("Content is:\n\n%s\n", string(contents))

}
