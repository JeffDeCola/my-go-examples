# my-go-examples

[![Code Climate](https://codeclimate.com/github/JeffDeCola/my-go-examples/badges/gpa.svg)](https://codeclimate.com/github/JeffDeCola/my-go-examples)
[![Issue Count](https://codeclimate.com/github/JeffDeCola/my-go-examples/badges/issue_count.svg)](https://codeclimate.com/github/JeffDeCola/my-go-examples/issues)
[![Go Report Card](https://goreportcard.com/badge/jeffdecola/my-go-examples)](https://goreportcard.com/report/jeffdecola/my-go-examples)
[![License](http://img.shields.io/:license-mit-blue.svg)](http://jeffdecola.mit-license.org)

`my-go-examples` _is a place to keep my go code snippets and examples._

I organized everything into these sections,

* API
* BASIC PROGRAMMING
* DATABASE
* GOROUTINES
* MESSAGING
* WEBSERVER

These go examples also contain info I gathered from other sources.

View my entire list of go examples on
[my-go-examples GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## GO EXAMPLES

_All sections in alphabetical order._

* API

  * [http-GET-POST](https://github.com/JeffDeCola/my-go-examples/tree/master/api/http-GET-POST)

    _Using http package to GET or POST (update via json)._

  * [OAuth-2.0-google-cloud-storage-api](https://github.com/JeffDeCola/my-go-examples/tree/master/api/OAuth-2.0-google-cloud-storage-api)

    _Using OAuth 2.0 to access a users google-cloud-storage (based on scopes) via googles api._

  * [OAuth-2.0-google-cloud-storage-api-over-NATS](https://github.com/JeffDeCola/my-go-examples/tree/master/api/OAuth-2.0-google-cloud-storage-api-over-NATS)

    _Using OAuth 2.0 (frontend and backend via protobuf over NATS) to
    access a users google cloud storage (based on scopes) via googles api._

  * [simple-webserver-with-REST](https://github.com/JeffDeCola/my-go-examples/tree/master/api/simple-webserver-with-REST)

    _Adding REST to the [simple-webserver](https://github.com/JeffDeCola/my-go-examples/tree/master/webserver/simple-webserver)._

  * [something-postgreSQL](https://github.com/JeffDeCola/my-go-examples/tree/master/api/something-postgreSQL)

    _tbd_

  * [something-RESTful](https://github.com/JeffDeCola/my-go-examples/tree/master/api/something-RESTful)

    _tbd_
  
  * [something-youtube-content-id-api](https://github.com/JeffDeCola/my-go-examples/tree/master/api/something-youtube-content-id-api)

    _tbd_

  * [something-youtube-data-api-v3](https://github.com/JeffDeCola/my-go-examples/tree/master/api/something-youtube-data-api-ve)

    _tbd_

  * [track-email-mailgun-api](https://github.com/JeffDeCola/my-go-examples/tree/master/api/track-email-mailgun-api)

    _Send and track an email using mailgun over their go client api_

* BASIC PROGRAMMING

  * [gomock](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-programming/gomock)

    _A halloween theme is used for gomock on an interface for unit testing._

  * [gotests-complex-function](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-programming/gotests-complex-function)

    _Testing a function with complex inputs and outputs._

  * [json](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-programming/json)

    _Encode a struct to json and decode back to a struct._

  * [linked-list](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-programming/linked-list)

    _An example of a Singly Linked List (i.e. using just a head pointer)._

  * [logging-error-handling](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-programming/logging-error-handling)

    _Logging and error handling._

  * [pointers](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-programming/pointers)

    _An example of a pointer to a struct._*

  * [read-file](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-programming/read-file)

    _Reading a file a few different ways._

  * [recursion](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-programming/recursion)

    _A function calling itself to make a fibonacci series._

  * [structs-methods-interfaces](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-programming/structs-methods-interfaces)

    _An example of structs, methods and interfaces._

* DATABASE

  * [google-cloud-spanner-table](https://github.com/JeffDeCola/my-go-examples/tree/master/database/google-cloud-spanner-table)

    _read/write from/to a table._

  * [postgreSQL](https://github.com/JeffDeCola/my-go-examples/tree/master/database/postgreSQL)

    _read/write from/to a table._

* GOROUTINES

  * [async-channel-no-waiting](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines/async-channel-no-waiting)

    _An example of a goroutine asynchronously sending data (via a channel) to a function that
    uses the latest data (if available) and does not wait._

  * [goroutines-channels-select](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines/goroutines-channels-select)

    _An example of concurrency and message passing via channels in go._

  * [goroutines-waitgroup](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines/goroutines-waitgroup)

    _An example of concurrency using a waitgroup._

  * [goroutines-worker-pools](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines/goroutines-worker-pools)

    _An example of concurrency using a worker pool with goroutines and channels._

* MESSAGING

  * [protobuf](https://github.com/JeffDeCola/my-go-examples/tree/master/messaging/protobuf)

    _Protocol buffers serialize structured data, useful for messaging._

  * [protobuf-NATS-request-response](https://github.com/JeffDeCola/my-go-examples/tree/master/messaging/protobuf-NATS-request-response)

    _Sends a protobuf msg over NATS from a client to a server using request and response._

  * [protobuf-NATS-publish-subscribe](https://github.com/JeffDeCola/my-go-examples/tree/master/messaging/protobuf-NATS-publish-subscribe)

    _Sends a protobuf msg over NATS from a client to a server using publish and subscribe._

* WEBSERVER

  * [simple-webserver](https://github.com/JeffDeCola/my-go-examples/tree/master/webserver/simple-webserver)

    _Using the http package to build a simple webserver._

## UNIT TESTING AND MY GITHUB WEBPAGE IS UPDATED USING CONCOURSE

For fun, I use concourse to automate unit testing, update
[my-go-examples GitHub Webpage](https://jeffdecola.github.io/my-go-examples/) and alert me of
the changes via repo status and slack.

The unit testing is accomplished by running this script this script
[here](https://github.com/JeffDeCola/my-go-examples/tree/master/ci/scripts/unit-tests.sh).

The github webpage update is accomplished this by copying and editing
this `README.md` file to `/docs/_includes/README.md`.
You can see the concourse task (a shell script)
[here](https://github.com/JeffDeCola/my-go-examples/tree/master/ci/scripts/readme-github-pages.sh).

A pipeline file [pipeline.yml](https://github.com/JeffDeCola/my-go-examples/tree/master/ci/pipeline.yml)
shows the entire ci flow. Visually, it looks like,

![IMAGE - my-go-examples concourse ci pipeline - IMAGE](docs/pics/my-go-examples-pipeline.jpg)

For more information on using concourse for continuous integration,
refer to my cheat sheet on [concourse](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/operations-tools/continuous-integration-continuous-deployment/concourse-cheat-sheet).
