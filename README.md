# my-go-examples

[![Go Report Card](https://goreportcard.com/badge/github.com/JeffDeCola/my-go-examples)](https://goreportcard.com/report/github.com/JeffDeCola/my-go-examples)
[![GoDoc](https://godoc.org/github.com/JeffDeCola/my-go-examples?status.svg)](https://godoc.org/github.com/JeffDeCola/my-go-examples)
[![Maintainability](https://api.codeclimate.com/v1/badges/3c5477c63d77a071fdff/maintainability)](https://codeclimate.com/github/JeffDeCola/my-go-examples/maintainability)
[![Issue Count](https://codeclimate.com/github/JeffDeCola/my-go-examples/badges/issue_count.svg)](https://codeclimate.com/github/JeffDeCola/my-go-examples/issues)
[![License](http://img.shields.io/:license-mit-blue.svg)](http://jeffdecola.mit-license.org)

`my-go-examples` _is a place to keep my go code snippets and examples._

I organized everything (alphabetically) into these sections,

* [API](https://github.com/JeffDeCola/my-go-examples#api)
* [BASIC PROGRAMMING](https://github.com/JeffDeCola/my-go-examples#basic-programming)
* [BASIC SYNTAX](https://github.com/JeffDeCola/my-go-examples#basic-syntax)
  * Data Types
  * Functions
  * Interfaces
  * Methods
* [CGO](https://github.com/JeffDeCola/my-go-examples#cgo)
* [CLOUD SERVICES](https://github.com/JeffDeCola/my-go-examples#cloud-services)
* [DATABASE](https://github.com/JeffDeCola/my-go-examples#database)
* [GOROUTINES](https://github.com/JeffDeCola/my-go-examples#goroutines)
* [JEFFS GO TEMPLATES](https://github.com/JeffDeCola/my-go-examples#jeffs-go-templates)
* [MESSAGING](https://github.com/JeffDeCola/my-go-examples#messaging)
* [PACKAGES](https://github.com/JeffDeCola/my-go-examples#packages)
* [TESTING](https://github.com/JeffDeCola/my-go-examples#testing)
* [USEFUL TOOLS I CREATED](https://github.com/JeffDeCola/my-go-examples#useful-tools-i-created)
* [WEBSERVER](https://github.com/JeffDeCola/my-go-examples#webserver)

These go examples also contain info I gathered from other sources.

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## GO EXAMPLES

_All sections in alphabetical order._

### API

* [http-GET-POST](https://github.com/JeffDeCola/my-go-examples/tree/master/api/http-GET-POST)

  _Using a http package to GET or POST (update via json)._

* [OAuth-2.0-google-cloud-storage-api](https://github.com/JeffDeCola/my-go-examples/tree/master/api/OAuth-2.0-google-cloud-storage-api)

  _Using OAuth 2.0 to access a users google cloud storage
  (based on scopes) via googles api._

* [OAuth-2.0-google-cloud-storage-api-over-NATS](https://github.com/JeffDeCola/my-go-examples/tree/master/api/OAuth-2.0-google-cloud-storage-api-over-NATS)

  _Using OAuth 2.0 (frontend and backend via protobuf over NATS)
  to access a users google cloud storage (based on scopes) via googles api._

* [simple-webserver-with-REST](https://github.com/JeffDeCola/my-go-examples/tree/master/api/simple-webserver-with-REST)

  _Adding REST to my
  [simple-webserver](https://github.com/JeffDeCola/my-go-examples/tree/master/webserver/simple-webserver)._

* [something-postgreSQL](https://github.com/JeffDeCola/my-go-examples/tree/master/api/something-postgreSQL)

  _tbd._

* [something-RESTful](https://github.com/JeffDeCola/my-go-examples/tree/master/api/something-RESTful)

  _tbd._
  
* [something-youtube-content-id-api](https://github.com/JeffDeCola/my-go-examples/tree/master/api/something-youtube-content-id-api)

  _tbd._

* [something-youtube-data-api-v3](https://github.com/JeffDeCola/my-go-examples/tree/master/api/something-youtube-data-api-v3)

  _tbd._

* [track-email-mailgun-api](https://github.com/JeffDeCola/my-go-examples/tree/master/api/track-email-mailgun-api)

  _Send and track an email using mailgun over their go client api._

### BASIC PROGRAMMING

_Just cool programs that are very useful and can be used as building blocks for
other cool programs._

* [json](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-programming/json)

  _Encode a struct to json and decode back to a struct._

* [read-file](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-programming/read-file)

  _Reading a file a few different ways._

* [singly-linked-list](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-programming/singly-linked-list)

  _A singly linked list (i.e. using just a head pointer)._

### BASIC SYNTAX

* DATA TYPES

  * [closure](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/data-types/closure)

    _Assigning a variable to a function._

  * [pointers1](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/data-types/pointers1)

    _A struct passed to function by value (copy) and by address (reference)._

  * [pointers2](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/data-types/pointers2)

    _An int passed to function by address (reference) and by int pointer (reference)._

  * [slices-underlying-array](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/data-types/slices-underlying-array)

    _A slice being appended to, and shows the underlying array
    being added to the slice._

* FUNCTIONS

  * [callback](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/functions/callback)

    _Passing a function (as an argument) to a function._

  * [recursion](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/functions/recursion)

    _A function calling itself to make a fibonacci series._

* INTERFACES

  * [interface](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/interface)

    _Accepting ANYTHING (as long as anything has a method attached
    to an interface)._

  * [interfaces-card-atm-bank](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/interfaces-card-atm-bank)

    _Using interfaces to show how an atm card is
    used to interact with an atm and bank._

  * [structs-methods-interfaces](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/structs-methods-interfaces)

    _Structs, methods and interfaces._

* METHODS

  * tbd

### CGO

_Using c with go._

* [simple-c-code](https://github.com/JeffDeCola/my-go-examples/tree/master/cgo/simple-c-code)

  _A very simple example to show you how to write a c function in go._

* [simple-c-code-using-stdio](https://github.com/JeffDeCola/my-go-examples/tree/master/cgo/simple-c-code-using-stdio)

  _Same code from above, but using stdio.h for printf._

### CLOUD SERVICES

* [amazon-aws](https://github.com/JeffDeCola/my-go-examples/tree/master/cloud-services/amazon-aws)

  _Interacting with amazon's
  [aws](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-providers/amazon-web-services-cheat-sheet)
  using
  `aws-sdk-go`
  client libraries._

* [google-gcp](https://github.com/JeffDeCola/my-go-examples/tree/master/cloud-services/google-gcp)

  _Interacting with google's
  [gcp](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-providers/google-cloud-platform-cheat-sheet)
  using
  [google-cloud-go](https://github.com/googleapis/google-api-go-client/tree/master)
  client libraries._

### DATABASE

* [google-cloud-spanner-table](https://github.com/JeffDeCola/my-go-examples/tree/master/database/google-cloud-spanner-table)

  _tbd._

* [postgreSQL](https://github.com/JeffDeCola/my-go-examples/tree/master/database/postgreSQL)

  _A relational database - Read/Write from/to a table._

* [redis](https://github.com/JeffDeCola/my-go-examples/tree/master/database/redis)

  _A non-relational (NoSQL) database - Set/Get from/to a key/value pair._

### GOROUTINES

_Go is written for concurrency.
The go runtime schedules goroutines on threads.
The OS schedules these threads on cpus/cores._

* [goroutines-async-channel-receive-no-waiting](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines/goroutines-async-channel-receive-no-waiting)

  _A goroutine asynchronously sending data (via a channel buffer) and a goroutine
  receiving that data, using the latest (if available) and does not wait._

* [goroutines-async-channel-send-receive-waiting](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines/goroutines-async-channel-send-receive-waiting)

  _Sending data to a goroutine via an async channel.
  Both SEND and RCV will block/wait for data._

* [goroutines-multi-core](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines/goroutines-multi-core)

  _Concurrency across multiples cores.  You can play around with workers,
  threads, cpus/cores and nice to find the fastest performance.
  It will find the total amount of prime numbers up to a number._

* [goroutines-waitgroup](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines/goroutines-waitgroup)

  _Concurrency using a waitgroup (waiting for a collection of goroutines to finish)._

* [goroutines-worker-pools](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines/goroutines-worker-pools)

  _Concurrency using worker pools._

### JEFFS GO TEMPLATES

_A starting point for all my go projects._

* [jeffs-basic-go-template](https://github.com/JeffDeCola/my-go-examples/tree/master/jeffs-go-templates/jeffs-basic-go-template)

  _A simple go template with flags, logging & error handling._

### MESSAGING

* [protobuf](https://github.com/JeffDeCola/my-go-examples/tree/master/messaging/protobuf)

  _Protocol buffers serialize structured data, useful for messaging._

* [protobuf-NATS-publish-subscribe](https://github.com/JeffDeCola/my-go-examples/tree/master/messaging/protobuf-NATS-publish-subscribe)

  _Sending a protobuf msg over NATS from a client
  to a server using publish and subscribe._

* [protobuf-NATS-request-response](https://github.com/JeffDeCola/my-go-examples/tree/master/messaging/protobuf-NATS-request-response)

  _Sending a protobuf msg over NATS from a
  client to a server using request and response._

### PACKAGES

_Refer to my repo
[my-go-packages](https://github.com/JeffDeCola/my-go-packages)
for packages I created. Refer to
[golang.org/pkg](https://golang.org/pkg/)
for the standard go packages_

* [crypto/md5](https://github.com/JeffDeCola/my-go-examples/tree/master/packages/crypto-md5)
  _(standard)_

  _The crypto/md5 package is used for getting an md5 fingerprint from an ssh key._

* [errors](https://github.com/JeffDeCola/my-go-examples/tree/master/packages/errors)
  _(github.com/pkg/errors)_

  _Error Handling using non-standard github.com/pkg/errors package._

* [flag](https://github.com/JeffDeCola/my-go-examples/tree/master/packages/flag)
  _(standard)_

  _The flag package makes it easy to implement command-line flag parsing._

* [fmt-scan](https://github.com/JeffDeCola/my-go-examples/tree/master/packages/fmt-scan)
  _(standard)_

  _The fmt package makes it easy to scan user input using fmt.Scan() by reading os.Stdin._

* [logrus](https://github.com/JeffDeCola/my-go-examples/tree/master/packages/logrus)
  _(github.com/sirupsen/logrus)_

  _Logging using non-standard logrus package._

* [test-jeffshapes](https://github.com/JeffDeCola/my-go-examples/tree/master/packages/test-jeffshapes)
  _(github.com/JeffDeCola/my-go-packages/jeffshapes)_
  
  _Testing the jeffshapes package (below)._

### TESTING

* [gomock](https://github.com/JeffDeCola/my-go-examples/tree/master/testing/gomock)

  _Using gomock on an interface for unit testing._

* [gotests-complex-function](https://github.com/JeffDeCola/my-go-examples/tree/master/testing/gotests-complex-function)

  _Testing a function with complex inputs and outputs._

### USEFUL TOOLS I CREATED

* [markdown-create-table-of-contents](https://github.com/JeffDeCola/my-go-examples/tree/master/useful-tools-i-created/markdown-create-table-of-contents)

  _Parse a markdown file to find `##`, `###` to create a table of contents (TOC)
  for links at github.  I like to use a TOC at the beginning of my README.md files._

* [markdown-delimiter-doer](https://github.com/JeffDeCola/my-go-examples/tree/master/useful-tools-i-created/markdown-delimiter-doer)

  _Take a markdown file and "do whatever you want" between the delimiters
  and output new markdown file._

### WEBSERVER

* [simple-webserver](https://github.com/JeffDeCola/my-go-examples/tree/master/webserver/simple-webserver)

  _Using the http package to build a simple webserver._

## UPDATE GITHUB WEBPAGE & UNIT TESTS USING CONCOURSE (OPTIONAL)

For fun, I use concourse to  update
[my-go-examples GitHub Webpage](https://jeffdecola.github.io/my-go-examples/),
run my unit-tests and alert me of the changes via repo status and slack.

A pipeline file [pipeline.yml](https://github.com/JeffDeCola/my-go-examples/tree/master/ci/pipeline.yml)
shows the entire ci flow. Visually, it looks like,

![IMAGE - my-go-examples concourse ci pipeline - IMAGE](docs/pics/my-go-examples-pipeline.jpg)

The `jobs` and `tasks` are,

* `job-readme-github-pages` runs task
  [readme-github-pages.sh](https://github.com/JeffDeCola/my-go-examples/tree/master/ci/scripts/readme-github-pages.sh).
* `job-unit-tests` runs task
  [unit-tests.sh](https://github.com/JeffDeCola/my-go-examples/tree/master/ci/scripts/unit-tests.sh).

The concourse `resources types` are,

* `my-go-examples` uses a resource type
  [docker-image](https://hub.docker.com/r/concourse/git-resource/)
  to PULL a repo from github.
* `resource-slack-alert` uses a resource type
  [docker image](https://hub.docker.com/r/cfcommunity/slack-notification-resource)
  that will notify slack on your progress.
* `resource-repo-status` uses a resource type
  [docker image](https://hub.docker.com/r/dpb587/github-status-resource)
  that will update your git status for that particular commit.

For more information on using concourse for continuous integration,
refer to my cheat sheet on [concourse](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations-tools/continuous-integration-continuous-deployment/concourse-cheat-sheet).
