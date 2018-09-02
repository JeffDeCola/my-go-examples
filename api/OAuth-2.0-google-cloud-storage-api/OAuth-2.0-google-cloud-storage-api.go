package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

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
var cs ClientSecret

const htmlIndex = `
<html>
<body>
<a href="/GoogleLogin">Jeff Test - Log in with Google /GoogleLogin</a>
<br />
<br />
<a href="/useToken?call=all">Use the Token to get a list of buckets /useToken?call=all</a>
<br />
<a href="/useToken?call=images-na">Use the Token to get metadata on a particular bucker /useToken?call=images-na</a>
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

	// url := googleOauthConfig.AuthCodeURL(oauthStateString)
	url :=
		"https://accounts.google.com/o/oauth2/v2/auth" + "?" +
			"client_id=" + cs.Web.ClientID + "&" +
			"redirect_uri=" + "http://127.0.0.1:3000/GoogleCallback" + "&" +
			"response_type=" + "code" + "&" +
			"scope=" + "https://www.googleapis.com/auth/devstorage.read_only" + "&" +
			"access_type=" + "offline" + "&" +
			"state=" + oauthStateString
	fmt.Printf("    url for goolge login is %+v\n", url)
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

	code := req.FormValue("code")
	fmt.Printf("    returned code is %s \n", code)

	// Get the Token - Exchange Auth Code for Token.
	var err error
	token, err = googleOauthConfig.Exchange(oauth2.NoContext, code)
	fmt.Printf("    token is %+v \n", token)
	fmt.Printf("      - token.AccessToken is %+v \n", token.AccessToken)
	fmt.Printf("      - token.RefreshToken is %+v \n", token.RefreshToken)
	fmt.Printf("      - token.TokenType is %+v \n", token.TokenType)
	fmt.Printf("      - token.Expiry is %+v \n", token.Expiry.Format(time.RFC3339))

	if err != nil {
		fmt.Printf("    Token Issue - oauthConf.Exchange() failed with '%s'\n", err)
		http.Redirect(res, req, "/", http.StatusTemporaryRedirect)
		return
	}

	http.Redirect(res, req, "/", http.StatusTemporaryRedirect)

}

func handleUseToken(res http.ResponseWriter, req *http.Request) {
	// print out the req struct for fun
	fmt.Printf("\nhandleUseToken - req is %+v\n", req.URL)

	// LETS USE THE TOKEN
	// response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	fmt.Printf("    token is %+v \n", token)
	fmt.Printf("      - token.AccessToken is %+v \n", token.AccessToken)
	fmt.Printf("      - token.RefreshToken is %+v \n", token.RefreshToken)
	fmt.Printf("      - token.TokenType is %+v \n", token.TokenType)
	fmt.Printf("      - token.Expiry is %+v \n", token.Expiry.Format(time.RFC3339))

	call := req.FormValue("call")

	var err error
	var response *http.Response

	switch call {
	case "all":
		fmt.Printf("      Call is 'all' \n")
		response, err = http.Get("https://www.googleapis.com/storage/v1/b?access_token=" + token.AccessToken + "&project=lofty-outcome-860")
	case "images-na":
		fmt.Printf("      Call is 'images-na' \n")
		response, err = http.Get("https://www.googleapis.com/storage/v1/b/images-na?access_token=" + token.AccessToken)
	default:
		fmt.Printf("      Call is 'default' \n")
		response, err = http.Get("https://www.googleapis.com/storage/v1/b/images-na?access_token=" + token.AccessToken)
	}

	if err != nil {
		fmt.Printf("    Trying to use token '%s'\n", err)
		http.Redirect(res, req, "/", http.StatusTemporaryRedirect)
		return
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	fmt.Fprintf(res, "    Content: %s\n", contents)
}

func unmarshalJSONFile() {

	// Read the secrets file from google
	// Generate from https://console.developers.google.com/projectselector/apis/credentials
	raw, err := ioutil.ReadFile(os.Getenv("HOME") + "/secrets/client-secrets-stepan.json")
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
	http.HandleFunc("/useToken", handleUseToken)
	fmt.Println(http.ListenAndServe(":3000", nil))
}
