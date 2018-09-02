package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	nats "github.com/go-nats"
	"github.com/protobuf/proto"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	//"io/ioutil"
)

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
	// Some random string, random for each request
	oauthStateString = "jeffrandom"
)

// Other scopes
// "https://www.googleapis.com/auth/userinfo.profile",
// "https://www.googleapis.com/auth/userinfo.email",

var token *oauth2.Token
var code string
var cs ClientSecret

const htmlIndex = `
<html>
<body>
<a href="/GoogleLogin">Jeff Test - Log in with Google /GoogleLogin</a>
<br />
<a href="/SendAction?call=refresh">Refresh Token /SendAction?call=refresh</a>
<br />
<a href="/SendAction?call=tokeninfo">Token Info /SendAction?call=tokeninfo</a>
<br />
<br />
<a href="/SendAction?call=all">Get a list of buckets /SendAction?call=all</a>
<br />
<a href="/SendAction?call=images-na">Get metadata on a particular bucket /SendAction?call=images-na</a>
</body>
</html>
`

func handleMain(res http.ResponseWriter, req *http.Request) {
	// print out the req struct for fun
	fmt.Printf("\nhandleMain - req is %+v\n", req.URL)
	fmt.Fprintf(res, htmlIndex)
}

func handleGoogleLogin(res http.ResponseWriter, req *http.Request) {
	// print out the req struct for fun
	fmt.Printf("\nhandleGoogleLogin - req is %+v\n", req.URL)
	fmt.Printf("    oauthStateString is %+v\n", oauthStateString)
	// fmt.Printf("    googleOauthConfig is %+v\n", googleOauthConfig)

	// Note, make this offline to get refresh token
	// url := googleOauthConfig.AuthCodeURL(oauthStateString, oauth2.AccessTypeOffline)
	url :=
		"https://accounts.google.com/o/oauth2/v2/auth" + "?" +
			"client_id=" + cs.Web.ClientID + "&" +
			"redirect_uri=" + "http://127.0.0.1:3000/GoogleCallback" + "&" +
			"response_type=" + "code" + "&" +
			"scope=" + "https://www.googleapis.com/auth/devstorage.read_only" + "&" +
			"access_type=" + "offline" + "&" +
			"state=" + oauthStateString
	fmt.Printf("    URL for goolge login is %+v\n", url)
	http.Redirect(res, req, url, http.StatusTemporaryRedirect)
}

func handleGoogleCallback(res http.ResponseWriter, req *http.Request) {
	// print out the req struct for fun
	fmt.Printf("\nhandleGoogleCallback - req is %+v\n", req.URL)

	state := req.FormValue("state")
	fmt.Printf("    state is %s.  Should match %s. \n", state, oauthStateString)

	if state != oauthStateString {
		fmt.Printf("    invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(res, req, "/", http.StatusTemporaryRedirect)
		return
	}

	code = req.FormValue("code")
	fmt.Printf("    returned code is %s \n", code)

	// Get the Token - Exchange Auth Code for Token.
	//token, _ := googleOauthConfig.Exchange(oauth2.NoContext, code)
	//fmt.Printf("    token is %+v \n", token)
	//fmt.Printf("      - token.AccessToken is %+v \n", token.AccessToken)
	//fmt.Printf("      - token.RefreshToken is %+v \n", token.RefreshToken)
	//fmt.Printf("      - token.TokenType is %+v \n", token.TokenType)
	//fmt.Printf("      - token.Expiry is %+v \n", token.Expiry.Format(time.RFC3339))

	sendAuthCode()

	http.Redirect(res, req, "/", http.StatusTemporaryRedirect)

}

func sendAuthCode() {

	// PASS CODE TO BACK-END protobuf over NATS
	authCode := &AuthCode{
		AuthCode: code,
	}

	// Connect to NATS SERVER
	nc, _ := nats.Connect("nats://127.0.0.1:4222")
	log.Println("Connected to " + nats.DefaultURL)

	// PROTOBUF - CLIENT - MARSHAL - WRITE/SEND
	msg, err := proto.Marshal(authCode)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	log.Printf("   auth code sending :    %+v", authCode)

	// NATS - PUBLISH on "code"
	log.Printf("   Publishing to subject 'code'\n")
	nc.Publish("code", msg)

}

func handleSendAction(res http.ResponseWriter, req *http.Request) {
	// print out the req struct for fun
	fmt.Printf("\nhandleSendAction - req is %+v\n", req.URL)
	call := req.FormValue("call")

	sendAction(call)

	http.Redirect(res, req, "/", http.StatusTemporaryRedirect)

}

func sendAction(call string) {

	// PASS call TO BACK-END protobuf over NATS
	action := &Action{
		Action: call,
	}

	// Connect to NATS SERVER
	nc, _ := nats.Connect("nats://127.0.0.1:4222")
	log.Println("Connected to " + nats.DefaultURL)

	// PROTOBUF - CLIENT - MARSHAL - WRITE/SEND
	msg, err := proto.Marshal(action)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	log.Printf("   Action sending :    %+v", action)

	// NATS - PUBLISH on "action"
	log.Printf("   Publishing to subject 'action'\n")
	nc.Publish("action", msg)

}

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

func main() {

	unmarshalJSONFile()

	http.HandleFunc("/", handleMain)
	http.HandleFunc("/GoogleLogin", handleGoogleLogin)
	http.HandleFunc("/GoogleCallback", handleGoogleCallback)
	http.HandleFunc("/SendAction", handleSendAction)
	fmt.Println(http.ListenAndServe(":3000", nil))
}
