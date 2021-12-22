package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	nats "github.com/go-nats"
	"github.com/protobuf/proto"
)

var token *oauth2.Token
var rcvAuthCode *AuthCode
var cs ClientSecret

// ClientSecret is the client secrets .json file
type ClientSecret struct {
	Web WebType `json:"web"`
}

// WebType is the client secrets .json file
type WebType struct {
	ClientID          string   `json:"client_id"`
	ProjectID         string   `json:"project_id"`
	AuthURI           string   `json:"auth_uri"`
	TokenURI          string   `json:"token_uri"`
	CertURL           string   `json:"auth_provider_x509_cert_url"`
	ClientSecret      string   `json:"client_secret"`
	RedirectURIS      []string `json:"redirect_uris"`
	JavascriptOrigins []string `json:"javascript_origins"`
}

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://127.0.0.1:3000/GoogleCallback",
		ClientID:     "",
		ClientSecret: "",
		Scopes:       []string{"https://www.googleapis.com/auth/devstorage.read_only"},
		Endpoint:     google.Endpoint,
	}
)

func unmarshalJSONFile() {

	// Read the secrets file from google
	// Generate from https://console.developers.google.com/projectselector/apis/credentials
	raw, err := ioutil.ReadFile(os.Getenv("HOME") + "/secrets/client-secrets.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("    Raw json is %s\n\n", string(raw))

	err = json.Unmarshal(raw, &cs)
	if err != nil {
		fmt.Println("There was an error:", err)
	}
	fmt.Printf("    ProjectID is: %s\n", cs.Web.ProjectID)
	fmt.Printf("    ClientID is: %s\n", cs.Web.ClientID)

	// Put the secrets in googleOauthConfig
	googleOauthConfig.ClientSecret = cs.Web.ClientSecret
	googleOauthConfig.ClientID = cs.Web.ClientID

	fmt.Printf("    googleOauthConfig.ClientID is: %s\n", googleOauthConfig.ClientID)
	fmt.Printf("    googleOauthConfig.ClientSecret is: %s\n", googleOauthConfig.ClientSecret)

}

// This is only for a NEW auth code.  Can;t use the same once twice.
func getAuthCode() {

	// Connect to NATS SERVER
	nc, _ := nats.Connect("nats://127.0.0.1:4222")
	log.Println("Connected to " + nats.DefaultURL)

	// NATS - SUBSCRIBE ON "code" to get auth code first
	log.Printf("Subscribing to subject 'code'\n")
	// Synchronous way - When I want to check for one msg at a time
	// Wait
	subCode, err := nc.SubscribeSync("code")
	if err != nil {
		log.Fatal("SubscribeSync error: ", err)
	}

	// Loop forever - Long Running
	for {

		fmt.Printf("Waiting for next AuthCode\n")
		msg, err := subCode.NextMsg(time.Duration(25000) * time.Second)
		if err != nil {
			log.Fatal("next message error: ", err)
		}
		// fmt.Printf("Received a NAT message: %v\n", msg)

		// PROTOBUF - SERVER - RECEIVE - READ/UNMARSHAL
		rcvAuthCode := &AuthCode{}
		err = proto.Unmarshal(msg.Data, rcvAuthCode)
		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}

		log.Printf("Auth Code received: %+v", rcvAuthCode)
		log.Printf("    AuthCode: %+v", rcvAuthCode.AuthCode)

		// Exchange authcode for token
		token = getToken(rcvAuthCode)

	}
}

func tokenInfo(token *oauth2.Token) {

	fmt.Printf("\ntoken:\n")
	fmt.Printf("      - token.AccessToken is %+v \n", token.AccessToken)
	fmt.Printf("      - token.RefreshToken is %+v \n", token.RefreshToken)
	fmt.Printf("      - token.TokenType is %+v \n", token.TokenType)
	fmt.Printf("      - token.Expiry is %+v \n", token.Expiry.Format(time.RFC3339))

	currentTime := time.Now().UTC()
	expireSeconds := token.Expiry.Sub(currentTime)
	log.Printf("    ******************* Token expires in: %s", expireSeconds)
}

func getToken(rcvAuthCode *AuthCode) *oauth2.Token {

	var err error

	// Get the Token - Exchange Auth Code for Token.
	token, err = googleOauthConfig.Exchange(oauth2.NoContext, rcvAuthCode.AuthCode)

	tokenInfo(token)

	if err != nil {
		log.Fatal("Exchange Token Error: ", err)
	}

	// READ - If no refresh token, pull it out of storage.  You must already have it.
	if token.RefreshToken == "" {
		fmt.Printf("There is no refresh token, get it from storage\n")
		token.RefreshToken = readRefreshToken()
	}

	// WRITE refresh token to file
	writeRefreshToken(token.RefreshToken)

	return token

}

func readRefreshToken() string {

	fmt.Printf("readRefreshToken()\n")
	refreshToken, err := ioutil.ReadFile("/Users/jeffdecola/refreshtoken.txt")
	if err != nil {
		log.Fatal("Read File Error: ", err)
	}

	fmt.Printf("Refresh Token from storage is %s\n", string(refreshToken))
	fmt.Printf("readRefreshToken() Done\n")
	return string(refreshToken)
}

func writeRefreshToken(refreshToken string) {

	fmt.Printf("writeRefreshToken()\n")
	d1 := []byte(refreshToken)
	err := ioutil.WriteFile("/Users/jeffdecola/refreshtoken.txt", d1, 0644)
	if err != nil {
		log.Fatal("Write File Error: ", err)
	}
	fmt.Printf("writeRefreshToken() done\n")

}

func takeAction() {

	var currentTime time.Time

	// Connect to NATS SERVER
	nc, _ := nats.Connect("nats://127.0.0.1:4222")
	log.Println("Connected to " + nats.DefaultURL)

	// NATS - SUBSCRIBE ON "action" to get action
	log.Printf("Subscribing to subject 'action'\n")
	// Synchronous way - When I want to check for one msg at a time
	subAction, err := nc.SubscribeSync("action")
	if err != nil {
		log.Fatal("SubscribeSync error: ", err)
	}

	// Loop forever - Long Running
	for {

		fmt.Printf("Waiting for next Action\n")
		msg, err := subAction.NextMsg(time.Duration(50000) * time.Second)
		if err != nil {
			log.Fatal("next message error: ", err)
		}
		// fmt.Printf("Received a NAT message: %v\n", msg)

		// PROTOBUF - SERVER - RECEIVE - READ/UNMARSHAL
		rcvAction := &Action{}
		err = proto.Unmarshal(msg.Data, rcvAction)
		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}

		log.Printf("Action received: %+v", rcvAction)
		log.Printf("    Action: %+v", rcvAction.Action)

		if rcvAction.Action == "refresh" {
			refreshToken()
		} else if rcvAction.Action == "tokeninfo" {
			tokenInfo(token)
		} else {

			currentTime = time.Now().UTC()
			expireSeconds := token.Expiry.Sub(currentTime)
			log.Printf("    ******************* Token expires in: %s", expireSeconds)
			if !token.Valid() { // if user token is expired
				fmt.Printf("Token has expired\n")
				// refreshToken()
			} else {
				fmt.Printf("Token is valid\n")
			}

			config := &oauth2.Config{
				ClientID:     cs.Web.ClientID,
				ClientSecret: cs.Web.ClientSecret,
				Endpoint: oauth2.Endpoint{
					TokenURL: "https://accounts.google.com/o/oauth2/token",
				},
			}
			restoredToken := &oauth2.Token{
				AccessToken:  token.AccessToken,
				RefreshToken: token.RefreshToken,
				Expiry:       token.Expiry,
				TokenType:    token.TokenType,
			}
			tokenInfo(restoredToken)
			tokenSource := config.TokenSource(oauth2.NoContext, restoredToken)
			client := oauth2.NewClient(oauth2.NoContext, tokenSource)
			token, err = tokenSource.Token()
			if err != nil {
				fmt.Printf("    Trying to use token %s on action '%s'\n", err, rcvAction.Action)
			}

			var response *http.Response

			switch rcvAction.Action {
			case "all":
				fmt.Printf("      Call is 'all' \n")
				response, err = client.Get("https://www.googleapis.com/storage/v1/b?project=lofty-outcome-860")
			case "images-na":
				fmt.Printf("      Call is 'images-na' \n")
				response, err = client.Get("https://www.googleapis.com/storage/v1/b/images-na")
			default:
				fmt.Printf("      Call is 'default' \n")
				response, err = client.Get("https://www.googleapis.com/storage/v1/b/images-na")
			}

			if err != nil {
				fmt.Printf("    Trying to use token %s on action '%s'\n", err, rcvAction.Action)
			}

			// fmt.Printf("    Content: %v\n", response)

			defer response.Body.Close()
			contents, _ := ioutil.ReadAll(response.Body)
			fmt.Printf("Content:\n %s\n", string(contents))
		}
	}

}

func refreshToken() {

	refreshToken := readRefreshToken()

	url1 := "https://accounts.google.com/o/oauth2/token"

	fmt.Printf("    URL for refresh token is %+v\n", url1)

	form := map[string][]string{
		"client_id":     {cs.Web.ClientID},
		"client_secret": {cs.Web.ClientSecret},
		"refresh_token": {refreshToken},
		"grant_type":    {"refresh_token"}}

	response, err := http.PostForm(url1,
		url.Values(form))
	if err != nil {
		fmt.Printf("    Trying to get refresh token %s\n", err)
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	fmt.Printf("    Content: %s\n", string(contents))

	//token.AccessToken = contents
	tokenInfo(token)

}

func main() {

	unmarshalJSONFile()

	go getAuthCode()
	takeAction()

}
