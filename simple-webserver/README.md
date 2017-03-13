# simple-webserver example

`simple-webserver` _uses the http package to build a simple webserver._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## The HTTP PACKAGE /net/http

The http package lets us map request paths to functions.

1. Setup which port you would like to listen on

```go
log.Fatal(http.ListenAndServe(":1234", nil))
```

2. When a request is made for a particular URL kick off your function jeffHandler()

```go
	http.HandleFunc("/jeff", jeffHandler)
```

3. Create your function jeffHandler

```go
func jeffHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Printf("req is %+v\n\n", req.URL)
	io.WriteString(res, "hello, Jeff!\n")
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