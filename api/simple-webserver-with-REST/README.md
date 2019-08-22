# simple-webserver-with-REST example

`simple-webserver-with-REST` _is an example of
adding REST to my
[simple-webserver](https://github.com/JeffDeCola/my-go-examples/tree/master/simple-webserver)._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## MOCK DATABASE

In this example a mock database will be used.

## HIGH-LEVEL VIEW OF CODE

![IMAGE - simple-webserver-with-REST - IMAGE](https://github.com/JeffDeCola/my-go-examples/blob/master/docs/pics/simple-webserver-with-REST.jpg)

## SETUP

You will need to get `github.com/gorilla/mux`

```bash
go get -u -v github.com/gorilla/mux
```

## RUN

```bash
go run simple-webserver-with-REST.go router.go routes.go handlers.go logger.go mockdatabase.go
```

In a browser,

[http://127.0.0.1:1234/](http://127.0.0.1:1234/)
