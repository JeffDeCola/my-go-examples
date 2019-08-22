# simple-webserver example

`simple-webserver` _is an example of
using the standard `net/http` package to build a simple webserver.

Also checkout my example
[simple-webserver-with-REST](https://github.com/JeffDeCola/my-go-examples/tree/master/api/simple-webserver-with-REST).

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## The HTTP PACKAGE /net/http

The `/net/http` package lets us map request paths to functions.

1. Set which port you would like to listen on,

    ```go
    log.Fatal(http.ListenAndServe(":1234", nil))
    ```

1. When a request is made for a particular URL/jeff kick off your function `jeffHandler()`,

    ```go
    http.HandleFunc("/jeff", jeffHandler)
    ```

1. Create your function `jeffHandler()`,

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

In another terminal, use a CLI http client like
httpie and you can do the following commands:

```bash
http localhost:1234
http localhost:1234/jeff
http localhost:1234/monkey
```

or

```bash
http 127.0.0.1:1234
http 127.0.0.1:1234/jeff
http 127.0.0.1:1234/monkey
```

Or you can use a browser,

```go
[http://127.0.0.1:1234/](http://127.0.0.1:1234/)
[http://127.0.0.1:1234/jeff](http://127.0.0.1:1234/jeff)
[http://127.0.0.1:1234/monkey](http://127.0.0.1:1234/monkey)
```
