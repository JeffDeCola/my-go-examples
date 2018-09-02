# something-youtube-content-id

`something-youtube-content-id` _???????????? is an api for YouTube's Rights Management Systems._

## AUTHORIZATION OAuth 2.0

Refer to
[Web Server-Side Flow](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/OAuth-2.0-authorization-cheat-sheet/OAuth-2.0-authorization-web-server-app-flow.md)
for more information about OAuth 2.0.

## STEP 1 - CREATE OAuth 2.0 CLIENT ID & SECRET

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

## STEP 2 - APP LOGIN PAGE

Create a link the user may click on to get redirected
to the google login page.

In this example it is [http://127.0.0.1:3000](http://127.0.0.1:3000).

The `golang/oauth2` client libraries
[here](https://github.com/golang/oauth2)
to implement OAuth 2.0 in your application.

## STEP 3 - GOOGLE LOGIN PAGE

The user gets redirected to the google login handler page via a url similiar to:

```bash
https://accounts.google.com/o/oauth2/auth?
    client_id={YOUR_SECRET}&
    redirect_uri=http%3A%2F%2F127.0.0.1%3A3000%2FGoogleCallback&
    response_type=code&
    scope={THE SCOPE YOU CHOOSE}&
    state=jeffrandom
```

## STEP 4 - USER LOGS IN TO GOOGLE ACCOUNT AND IS DIRECTED BACK

The call back has the state and an authorization code.

```bash
/GoogleCallback?
    state=jeffrandom&
    code={SECRET AUTH CODE}
```

## STEP 5 - VERIFY SAME STRING VIA STATE

We verify if it's the same state string.

## STEP 6 EXCHANGE AUTH CODE FOR TOKEN

If it is then we use the `code` to ask google for a
short-lived access `token`. We can save the code for future
use to get another token later.

```go
token, err = googleOauthConfig.Exchange(oauth2.NoContext, code)
```

## PROFIT - USE ACCESS TOKEN FOR API

You can use the `google/google-api-go-client` client libraries
[here](https://github.com/google/google-api-go-client)
to use APIs in your application.

For example,

```http
http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
```

### REFRESH ACCESS TOKEN

TBD
