# simple-webserver-with-REST example

_Adding REST (GET, POST, PUT, DELETE) JSON API to my
[simple-webserver](https://github.com/JeffDeCola/my-go-examples/tree/master/webserver/simple-webserver)._

Table of Contents,

* [WHAT IS REST](https://github.com/JeffDeCola/my-go-examples/tree/master/api/simple-webserver-with-REST#what-is-rest)
* [REST vs TCP/IP](https://github.com/JeffDeCola/my-go-examples/tree/master/api/simple-webserver-with-REST#rest-vs-tcpip)
* [HIGH-LEVEL VIEW OF CODE](https://github.com/JeffDeCola/my-go-examples/tree/master/api/simple-webserver-with-REST#high-level-view-of-code)
* [USING A ROUTER - GORILLA/MUX](https://github.com/JeffDeCola/my-go-examples/tree/master/api/simple-webserver-with-REST#using-a-router---gorillamux)
* [MOCKDATABASE](https://github.com/JeffDeCola/my-go-examples/tree/master/api/simple-webserver-with-REST#mockdatabase)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/api/simple-webserver-with-REST#run)
  * [NORMAL WEBPAGE (GUI)](https://github.com/JeffDeCola/my-go-examples/tree/master/api/simple-webserver-with-REST#normal-webpage-gui)
  * [GET](https://github.com/JeffDeCola/my-go-examples/tree/master/api/simple-webserver-with-REST#get)
  * [POST (add)](https://github.com/JeffDeCola/my-go-examples/tree/master/api/simple-webserver-with-REST#post-add)
  * [PUT (replace/update)](https://github.com/JeffDeCola/my-go-examples/tree/master/api/simple-webserver-with-REST#put-replaceupdate)
  * [DELETE](https://github.com/JeffDeCola/my-go-examples/tree/master/api/simple-webserver-with-REST#delete)
  
[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## WHAT IS REST

REST (REpresentational State Transfer) is just a standard way for
computers to communicate over the web.
It is stateless, meaning who cares what the computer is doing at the time.

There are 4 basic `HTTP verbs` we use in requests to
interact with resources in a REST system,

* **GET** - Retrieve a specific resource by id or a collection of resources
* **POST** - Create a new resource
* **PUT** - Update a specific resource by id
* **DELETE** - Remove a specific resource by id

## REST vs TCP/IP

REST does not have state, whereas tcp has an open connection and you
can assume a lot about the server.

## HIGH-LEVEL VIEW OF CODE

For simplicity, the code is broken into,

* **simple-webserver-with-REST.go** - Kicks off webserver
* **router.go** - The gorilla router
* **routers.go** - The list of the routes (e.g. /postdata)
* **handlers** - The functions to handle the routes
* **mockdatabase.go** - A slice of structs
* **logger.go** - A log wrapper for better output

![IMAGE - simple-webserver-with-REST - IMAGE](https://github.com/JeffDeCola/my-go-examples/blob/master/docs/pics/simple-webserver-with-REST.jpg)

## USING A ROUTER - GORILLA/MUX

You will need to get `github.com/gorilla/mux` which is
a popular package for writing web handlers.

```bash
go get -u -v github.com/gorilla/mux
```

## MOCKDATABASE

A simple mock database has been set up,

```go
TodoStruct{ID: "10", Name: "Write presentation", Completed: false},
TodoStruct{ID: "20", Name: "Eat Lunch", Completed: false},
TodoStruct{ID: "30", Name: "Pick up Milk", Completed: true},
```

## RUN

```bash
go run simple-webserver-with-REST.go \
       router.go routes.go handlers.go mockdatabase.go logger.go
```

### NORMAL WEBPAGE (GUI)

In a browser,

[127.0.0.1:1234](http://127.0.0.1:1234/)

Will also show entire database.

### GET

[127.0.0.1:1234/getdata/20](http://127.0.0.1:1234/getdata/20)

```txt
{"id":"20","name":"Eat Lunch","completed":false}
```

### POST (add)

This will add data,

```bash
curl -X POST \
     -H "Content-Type: application/json" \
     -d '{"name":"Feed Cat", "Completed": false}' \
     http://127.0.0.1:1234/postdata/40
```

Check it was added,

[127.0.0.1:1234/getdata/40](http://127.0.0.1:1234/getdata/40)

### PUT (replace/update)

This will replace/update data,

```bash
curl -X PUT \
     -H "Content-Type: application/json" \
     -d '{"name":"Feed Cat", "Completed": true}' \
     http://127.0.0.1:1234/putdata/40
```

### DELETE

```bash
curl -X DELETE \
     -H "Content-Type: application/json" \
     http://127.0.0.1:1234/deletedata/20
```
