package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	url := "https://httpbin.org/get"

	response, err := http.Get(url)
	checkErr(err)

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	checkErr(err)

	fmt.Printf("\nGET is:\n\n%s\n", string(contents))

}
