# simple-webserver-with-REST example

`simple-webserver-with-REST` _is an example of
adding REST to my
[simple-webserver](https://github.com/JeffDeCola/my-go-examples/tree/master/webserver/simple-webserver)._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## HIGH-LEVEL VIEW OF CODE

![IMAGE - simple-webserver-with-REST - IMAGE](https://github.com/JeffDeCola/my-go-examples/blob/master/docs/pics/simple-webserver-with-REST.jpg)

Notice that a mock database is used.  Just to keep it simple.

## PREREQUISITES

You will need to get `github.com/gorilla/mux` which is
a popular package for writing web handlers.

```bash
go get -u -v github.com/gorilla/mux
```

## RUN

```bash
go run simple-webserver-with-REST.go router.go routes.go handlers.go logger.go mockdatabase.go
```

In a browser,

[http://127.0.0.1:1234/](http://127.0.0.1:1234/)
