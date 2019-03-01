package main

import (
	"fmt"
)

// Looping forever - For the testing Marathon and Mesos
func main() {

	fmt.Println("Hello everyone")

}

/*
func init() {
	config.New("./")
}

func SendEmail(to string, subject string, text string, html string) {
	mg := mailgun.NewMailgun(
		"something.com",
		viper.GetString("MAILGUN_API_KEY"),
		viper.GetString("MAILGUN_PUBLIC_API_KEY"),
	)

	mg := mailgun.NewMailgun(mydomain, mailgunApiKey, mailgungpublicApiKey)

	message := mg.NewMessage("Pex <no-reply@something.com>", subject, text, to)
	message.SetHtml(html)
	if _, _, err := mg.Send(message); err != nil {
		log.SentryError(err)
	}
}

func main() {
	fmt.Println("RESTful")
}
*/
/* package main - edit this at some point
import "gopkg.in/mailgun/mailgun-go.v1"

message := mailgun.NewMessage(
    "sender@example.com",
    "Fancy subject!",
    "Hello from Mailgun Go!",
    "recipient@example.com")
resp, id, err := mg.Send(message)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("ID: %s Resp: %s\n", id, resp)

*/
