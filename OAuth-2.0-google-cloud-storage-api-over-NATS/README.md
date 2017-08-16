# OAuth-2.0-google-cloud-storage-api-over-NATS

`OAuth-2.0-google-cloud-storage-api-over-NATS` _uses OAuth 2.0
(frontend and backend via protobuf over NATS) to access a users
google cloud storage (based on scopes) via googles api._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## REFRESH TOKEN

You only get refresh token once.  So you must save it.
User must revoke privlidges [here](https://myaccount.google.com/permissions)
in order to get a new refresh token.

```go
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

tokenSource := config.TokenSource(oauth2.NoContext, restoredToken)
client := oauth2.NewClient(oauth2.NoContext, tokenSource)
token, err = tokenSource.Token()

// Now use it in the api call
response, err "= client.Get("https://www.googleapis.com/storage/v1/b?project=PROJECT_NAME")
```

## HIGH-LEVEL VIEW

This example expands on [OAuth-2.0-google-cloud-storage-api](https://github.com/JeffDeCola/my-go-examples/tree/master/OAuth-2.0-google-cloud-storage-api) and adds protobuf and NATS to pass auth code from `front-end` to `back-end`.

![IMAGE - OAuth-2.0-web-server-app-authorization-flow - IMAGE](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/OAuth-2.0-authorization-cheat-sheet/OAuth-2.0-web-server-app-authorization-flow.jpg)

## RUN

Boot up a NATS server such as,

```bash
gnatsd -DV -addr localhost --port 4222
```

Run both the frontend and backend processes,

```go
cd frontend
go run client.go messages.pb.go
```

```go
cd backend
go run server.go messages.pb.go
```

Then use by going to a browser,

```bash
http://127.0.0.1:3000
```
