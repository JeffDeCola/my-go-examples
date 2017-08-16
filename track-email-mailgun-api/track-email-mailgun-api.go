package main

import (
	"fmt"
)

func SendSimpleMessage(domain, apiKey string) (string, error) {
	mg := mailgun.NewMailgun(domain, apiKey, publicApiKey)
	m := mg.NewMessage(
		"Excited User <mailgun@YOUR_DOMAIN_NAME>",
		"Hello",
		"Testing some Mailgun awesomeness!",
		"YOU@YOUR_DOMAIN_NAME",
	)
	_, id, err := mg.Send(m)
	return id, err
}

func main() {
	fmt.Println("SEnding email to mailgun")
	return, _ := SendSimpleMessage()
}
