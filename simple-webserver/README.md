# simple-webserver example

`simple-webserver` _uses the http package to build a simple webserver._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## The HTTP PACKAGE /net/http

The http package lets us map request paths to functions.

```go
//reponse is the server
func response(w http.ResponseWriter, request *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {

	// Call the response function when you get a http request
	http.HandleFunc("/", response)

	// Starts listening on localhost (127.0.0.1:1234)
	log.Fatal(http.ListenAndServe(":1234", nil))

}
```

## RUN

```bash
go run simple-webserver.go
```

In another terminal, use a CLI http client like httpie and you can do the
following commands:

```bash
http localhost:1234
```

or

```bash
http 127.0.0.1:1234
```