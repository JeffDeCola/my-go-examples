package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	//"io/ioutil"
)

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://127.0.0.1:3000/GoogleCallback",
		ClientID:     os.Getenv("googlekey"),
		ClientSecret: os.Getenv("googlesecret"),
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

const htmlIndex = `
<html>
<body>
<a href="/GoogleLogin">Jeff Test - Log in with Google /GoogleLogin</a>
<br />
<br />
<a href="/useToken">Use the Token /useToken</a>
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

	url := googleOauthConfig.AuthCodeURL(oauthStateString)
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

	response, err := http.Get("https://www.googleapis.com/storage/v1/b/images-na?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Printf("    Trying to use token '%s'\n", err)
		http.Redirect(res, req, "/", http.StatusTemporaryRedirect)
		return
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	fmt.Fprintf(res, "Content: %s\n", contents)
}

func main() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/GoogleLogin", handleGoogleLogin)
	http.HandleFunc("/GoogleCallback", handleGoogleCallback)
	http.HandleFunc("/useToken", handleUseToken)
	fmt.Println(http.ListenAndServe(":3000", nil))
}
