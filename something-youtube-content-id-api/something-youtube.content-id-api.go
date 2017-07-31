package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v2"
	//"io/ioutil"
	"golang.org/x/net/context"
)

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://127.0.0.1:3000/GoogleCallback",
		ClientID:     os.Getenv("googlekey"),
		ClientSecret: os.Getenv("googlesecret"),
		Scopes:       []string{"https://www.googleapis.com/auth/drive", "https://www.googleapis.com/auth/drive.file"},
		Endpoint:     google.Endpoint,
	}
	// Some random string, random for each request
	oauthStateString = "jeffrandom"
)

const htmlIndex = `
<html>
<body>
<a href="/GoogleLogin">Jeff Test - Log in with Google</a>
</body>
</html>
`

func handleMain(res http.ResponseWriter, req *http.Request) {
	// print out the req struct for fun
	fmt.Printf("req is %+v\n\n", req.URL)
	fmt.Fprintf(res, htmlIndex)
}

func handleGoogleLogin(res http.ResponseWriter, req *http.Request) {
	// print out the req struct for fun
	fmt.Printf("req is %+v\n\n", req.URL)
	fmt.Printf("oauthStateString is %+v\n\n", oauthStateString)

	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(res, req, url, http.StatusTemporaryRedirect)
}

func handleGoogleCallback(res http.ResponseWriter, req *http.Request) {
	// print out the req struct for fun
	fmt.Printf("req is %+v\n\n", req.URL)

	state := req.FormValue("state")
	fmt.Printf("state is %s.  Should match %s. \n\n", state, oauthStateString)

	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(res, req, "/", http.StatusTemporaryRedirect)
		return
	}

	code := req.FormValue("code")
	fmt.Printf("code is %s \n\n", code)

	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	fmt.Printf("token is %+v \n\n", token)

	if err != nil {
		fmt.Printf("Toekn Issue - oauthConf.Exchange() failed with '%s'\n\n", err)
		http.Redirect(res, req, "/", http.StatusTemporaryRedirect)
		return
	}
	client := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(token))

	//response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	driveService, err := drive.New(client)

	myR, err := driveService.Files.List().MaxResults(10).Do()
	if err != nil {
		fmt.Fprintf(res, "Couldn't retrieve files %+v", err)
	}
	if len(myR.Items) > 0 {
		for _, i := range myR.Items {
			fmt.Fprintf(res, i.Title, " ", i.Id)
		}
	} else {
		fmt.Fprintf(res, "No files found.")
	}

	//defer response.Body.Close()
	//contents, err := ioutil.ReadAll(response.Body)
	//fmt.Fprintf(w, "Content: %s\n", contents)
}

func main() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/GoogleLogin", handleGoogleLogin)
	http.HandleFunc("/GoogleCallback", handleGoogleCallback)
	fmt.Println(http.ListenAndServe(":3000", nil))
}
