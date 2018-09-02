# simple-webserver example

`simple-webserver` _uses the http package to build a simple webserver._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

Also checkout [simple-webserver-with-REST](https://github.com/JeffDeCola/my-go-examples/tree/master/simple-webserver-with-REST).

## The HTTP PACKAGE /net/http

The http package lets us map request paths to functions.

1. Setup which port you would like to listen on

    ```go
    log.Fatal(http.ListenAndServe(":1234", nil))
    ```

1. When a request is made for a particular URL/jeff kick off your function jeffHandler()

    ```go
    http.HandleFunc("/jeff", jeffHandler)
    ```

1. Create your function jeffHandler()

    req is the request structure. res is your reponse.

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

## NO ROUTER

Hence, you can do `http://127.0.0.1:1234/foo` and it will work.