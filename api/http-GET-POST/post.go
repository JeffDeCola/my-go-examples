package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	formData := url.Values{
		"name": {"jeff"},
	}
	url := "https://httpbin.org/post"

	response, err := http.PostForm(url, formData)
	checkErr(err)

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	checkErr(err)

	fmt.Printf("\nPOST is:\n\n%s\n", string(contents))
}
