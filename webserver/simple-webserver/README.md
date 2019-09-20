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

Press return to exit.

You can interact with the web server many different ways.

### USING HTTPIE

In another terminal, use a CLI http client like
httpie and you can do the following,

```bash
http localhost:1234
http localhost:1234/jeff
http localhost:1234/monkey
```

Or you can use the IP Address,

```bash
http 127.0.0.1:1234
http 127.0.0.1:1234/jeff
http 127.0.0.1:1234/monkey
```

### USING CURL

Or you can use curl,

```bash
curl 127.0.0.1:1234
curl 127.0.0.1:1234/jeff
curl 127.0.0.1:1234/monkey
```

### USING A BROWSER

[http://127.0.0.1:1234/](http://127.0.0.1:1234/)

[http://127.0.0.1:1234/jeff](http://127.0.0.1:1234/jeff)

[http://127.0.0.1:1234/monkey](http://127.0.0.1:1234/monkey)
