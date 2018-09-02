# OAuth-2.0-google-cloud-storage-api

`OAuth-2.0-google-cloud-storage-api` _uses OAuth 2.0 to access a users google cloud storage (based on scopes) via googles api._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## AUTHORIZATION OAuth 2.0

Refer to
[Web Server-Side Flow](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/OAuth-2.0-authorization-cheat-sheet/OAuth-2.0-authorization-web-server-app-flow.md)
for a high-level view about OAuth 2.0.

## GETTING A TOKEN

The following steps allow you to get a token.

### STEP 1 - CREATE OAuth 2.0 CLIENT ID & SECRET

To create Create a `OAuth 2.0 Client ID` goto credentials page
[here](https://console.developers.google.com/projectselector/apis/credentials)
and select create credentials.

Create a OAuth 2.0 client IDs for a Web Application.

Origin URI,

[http://127.0.0.1:3000](http://127.0.0.1:3000)

Redirect URI,

[http://127.0.0.1:3000/GoogleCallback](http://127.0.0.1:3000/GoogleCallback)

You will now have a Client ID and a Secret.

The user opens the website and clicks the login button.

In the code is a great way to `unmarshalJSONFile()` the client secrets .json file.

### STEP 2 - APP LOGIN PAGE

Create a link the user may click on to get redirected
to the google login page.

In this example it is [http://127.0.0.1:3000](http://127.0.0.1:3000).

The `golang/oauth2` client libraries
[here](https://github.com/golang/oauth2)
to implement OAuth 2.0 in your application.

### STEP 3 - GOOGLE LOGIN PAGE

The user gets redirected to the google login handler page via a url similiar to:

```bash
https://accounts.google.com/o/oauth2/auth?
    client_id={YOUR_SECRET}&
    redirect_uri=http%3A%2F%2F127.0.0.1%3A3000%2FGoogleCallback&
    response_type=code&
    scope={THE SCOPE YOU CHOOSE}&
    state=jeffrandom
```

The scopes for this example are:
`https://www.googleapis.com/auth/devstorage.read_only`

The scopes availible are:

* `https://www.googleapis.com/auth/devstorage.full_control`
  Read/write and ACL management access to Google Cloud Storage.
* `https://www.googleapis.com/auth/devstorage.read_write`
  Read/write access to Google Cloud Storage.
* `https://www.googleapis.com/auth/devstorage.read_only`
  Read-only access to Google Cloud Storage.

### STEP 4 - USER LOGS IN TO GOOGLE ACCOUNT AND IS DIRECTED BACK

The call back has the state and an authorization code.

```bash
/GoogleCallback?
    state=jeffrandom&
    code={SECRET AUTH CODE}
```

### STEP 5 - VERIFY SAME STRING VIA STATE

We verify if it's the same state string.

### STEP 6 EXCHANGE AUTH CODE FOR TOKEN

IMPORTANT - Can only use the auth code once.

If it is then we use the `code` to ask google for a
short-lived access `token`. We can save the code for future
use to get another token later.

```go
token, err = googleOauthConfig.Exchange(oauth2.NoContext, code)
```

## PROFIT - USING ACCESS TOKEN FOR API (BASED ON SCOPES)

You can use the `google/google-api-go-client` client libraries
[here](https://github.com/google/google-api-go-client)
to use APIs in your application.

For example, to get meta data on YOUR_BUCKET_NAME,

```go
response, err := http.Get("https://www.googleapis.com/storage/v1/b/YOUR_BUCKET_NAME?access_token=" + token.AccessToken)
```

### REFRESH ACCESS TOKEN

TBD
