package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {

	form := map[string][]string{"q": {"github"}}

	response, err := http.PostForm("http://duckduckgo.com",
		url.Values(form))
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	fmt.Printf("post:\n %s", string(body))
}
