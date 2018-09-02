package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	mailgun "github.com/mailgun/mailgun-go"
)

const (
	replyAPIURL = "https://sw.api.mailgun.net/v3"
)

// Secrets is the client secrets .json file
type Secrets struct {
	APIKey             string `json:"apiKey"`
	PublicAPIKey       string `json:"publicApiKey"`
	FromEmailSubDomain string `json:"fromEmailSubDomain"`
	SendtoEmail        string `json:"sendtoEmail"`
}

func getSecrets() (string, string, string, string) {

	var s Secrets

	// Read the secrets file
	raw, err := ioutil.ReadFile(os.Getenv("HOME") + "/secrets/mailgun-secrets.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// unmarshal json
	err = json.Unmarshal(raw, &s)

	return s.APIKey, s.PublicAPIKey, s.FromEmailSubDomain, s.SendtoEmail
}

// MESSAGE API
func sendMessage(domain, APIKey, publicAPIKey, fromEmail, sendtoEmail, subject, body string) (string, string) {
	fmt.Printf("In sendMessage()\n")
	mg := mailgun.NewMailgun(domain, APIKey, publicAPIKey)
	message := mg.NewMessage(
		fromEmail,
		subject,
		body,
		sendtoEmail,
	)
	// message.AddCC("baz@example.com")
	// message.AddBCC("bar@example.com")
	// message.SetTracking(true)
	resp, id, err := mg.Send(message)
	if err != nil {
		log.Fatal(err)
	}
	return id, resp
}

// MESSAGE API
func getStoredMessage(domain, APIKey, publicAPIKey, id string) mailgun.StoredMessage {

	mg := mailgun.NewMailgun(domain, APIKey, publicAPIKey)
	mg.SetAPIBase(replyAPIURL) // Must do this since mailgun has wrong url

	resp, err := mg.GetStoredMessage(id)
	if err != nil {
		log.Fatal(err)
	}

	/*
		fields := map[string]string{
			"       From": resp.From,
			"     Sender": resp.Sender,
			"    Subject": resp.Subject,
			"Attachments": fmt.Sprintf("%d", len(resp.Attachments)),
			"    Headers": fmt.Sprintf("%d", len(resp.MessageHeaders)),
			"  BodyPlain": resp.BodyPlain,
		}
		for k, v := range fields {
			fmt.Printf("%13s: %s\n", k, v)
		}
	*/
	return resp
}

func testGetStoredMessage(domain string, APIKey string, publicAPIKey string, id string) {

	os.Setenv("MG_API_KEY", APIKey)
	os.Setenv("MG_DOMAIN", domain)
	os.Setenv("MG_PUBLIC_API_KEY", publicAPIKey)
	os.Setenv("MG_URL", replyAPIURL)

	mg, err := mailgun.NewMailgunFromEnv()
	if err != nil {
		log.Fatal(err)
		return
	}
	// id, err = findStoredMessageID(mg)
	// if err != nil {
	//	log.Fatal(err)
	//	return
	//}

	// First, get our stored message.
	msg, err := mg.GetStoredMessage("asdf")
	fmt.Printf("msg: %+v\n", msg)

	/*
		fields := map[string]string{
			"       From": msg.From,
			"     Sender": msg.Sender,
			"    Subject": msg.Subject,
			"  BodyPlain": msg.BodyPlain,
			"Attachments": fmt.Sprintf("%d", len(msg.Attachments)),
			"    Headers": fmt.Sprintf("%d", len(msg.MessageHeaders)),
		}
		for k, v := range fields {
			fmt.Printf("%13s: %s\n", k, v)
		}
	*/

	// We're done with it; now delete it.
	mg.DeleteStoredMessage(id)

}

// Tries to locate the first stored event type, returning the associated stored message key.
/*func findStoredMessageID(mg mailgun.Mailgun) (string, error) {
	ei := mg.NewEventIterator()
	err := ei.GetFirstPage(mailgun.GetEventsOptions{})
	for {
		if err != nil {
			fmt.Println("ugggggg")
			return "", err
		}
		if len(ei.Events()) == 0 {
			break
		}
		for _, event := range ei.Events() {
			if event["event"] == "stored" {
				s := event["storage"].(map[string]interface{})
				k := s["key"]
				return k.(string), nil
			}
		}
		err = ei.GetNext()
	}
	return "", fmt.Errorf("No stored messages found.  Try changing MG_EMAIL_TO to an address that stores messages and try again.")
}*/

// EVENTS API
func getMailLog(domain string, apiKey string, filter map[string]string) ([]mailgun.Event, error) {
	mg := mailgun.NewMailgun(domain, apiKey, "")
	ei := mg.NewEventIterator()
	begin := time.Now().Add(-3600 * time.Minute)
	err := ei.GetFirstPage(mailgun.GetEventsOptions{
		Begin:          &begin,
		ForceAscending: true,
		Limit:          1,
		Filter:         filter,
	})
	if err != nil {
		return nil, err
	}
	return ei.Events, nil
}

func createRoute(domain, fromEmailSubDomain, apiKey string) (mailgun.Route, error) {
	mg := mailgun.NewMailgun(domain, apiKey, "")
	return mg.CreateRoute(mailgun.Route{
		Priority:    1,
		Description: "Jeffs Route - store() all pex.YOURDOMAIN.com",
		Expression:  "match_recipient(\".*@" + fromEmailSubDomain + "\")",
		Actions: []string{
			//"forward(\"http://example.com/messages/\")",
			"store()",
			// "store(notify=\"http://pex.com/callback\")",
		},
	})
}

func listRoutes(domain, apiKey string) (int, []mailgun.Route, error) {
	mg := mailgun.NewMailgun(domain, apiKey, "")
	return mg.GetRoutes(-1, -1)
}

func main() {

	APIKey, publicAPIKey, fromEmailSubDomain, sendtoEmail := getSecrets()
	fmt.Printf("Secrets are %s, %s, %s %s\n", APIKey, publicAPIKey, fromEmailSubDomain, sendtoEmail)
	// uniqueID := "777"

	// SENDING EMAIL **********************************************************************************
	/*
		fmt.Printf("Sending email to mailgun using unique ID %s\n", uniqueID)
		subject := "Test with unique ID " + uniqueID
		body := "How's it going, unique ID is: " + uniqueID
		fromEmail := uniqueID + "@" + fromEmailSubDomain
		id, resp := sendMessage(fromEmailSubDomain, APIKey, publicAPIKey, fromEmail, sendtoEmail, subject, body)
		fmt.Printf("Response:\nID: %s\nRESP: %s\n", id, resp)
	*/

	// CHECK LOG **************************************************************************************

	fmt.Printf("Check Log\n")
	// id := ""
	filter := map[string]string{
		// "message-id": id,
		// "event": "rejected OR failed",
		"event": "stored",
		//"recipient": "NAME@gmail.com",
	}
	respLog, err := getMailLog(fromEmailSubDomain, APIKey, filter)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("Response:\n%+v\n", respLog)
	// b, err := json.MarshalIndent(respLog, "", "  ")
	// Put in a json serial string
	b, err := json.Marshal(respLog)
	// fmt.Printf("Response:\n%s\n", b)

	// Take json and put in map
	var data map[string]interface{}
	_ = json.Unmarshal([]byte(b), &data)
	fmt.Println(data)
	fmt.Println(data["id"])

	// Get the key to read from memory

	//os.Stdout.Write(b)

	// GET AND ERASE STORED MESSAGE BASED ON ID **************************************************************************************
	//fmt.Printf("Get Stored Message\n")
	//id = "eyJwIjpmYWxzZSwiayI6ImUwYzMyYTA1LTI0MGEtNDZiNS04YTExLWYyODcwNGViZWVkZiIsInMiOiJhZjllZGM3OWQwIiwiYyI6InRhbmtiIn0="
	//respStored := getStoredMessage(fromEmailSubDomain, APIKey, publicAPIKey, id)
	//fmt.Printf("Response Body:\n%v\n", respStored.BodyPlain)
	//testGetStoredMessage(fromEmailSubDomain, APIKey, publicAPIKey, id)

	// CREATE ROUTE **************************************************************************************
	// This code not needed - Easier to do this online
	// route, err1 := createRoute(fromEmailSubDomain, fromEmailSubDomain, APIKey)
	// if err1 != nil {
	//	log.Fatal(err1)
	// }
	// b, err := json.MarshalIndent(route, "", "  ")
	// fmt.Printf("Created Route is:\n%s\n", b)

	// LIST ROUTE **************************************************************************************
	// routeid, route2, err2 := listRoutes(fromEmailSubDomain, APIKey)
	// if err2 != nil {
	//		log.Fatal(err2)
	//	}
	//	b, err3 := json.MarshalIndent(route2, "", "  ")
	//	if err3 != nil {
	//		log.Fatal(err2)
	//	}
	//	fmt.Printf("List Route ID is:\n%d\n", routeid)
	//	fmt.Printf("List Route is:\n%s\n", b)

}
