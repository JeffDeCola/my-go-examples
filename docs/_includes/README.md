
# GO EXAMPLES

* GOROUTINES

  * [goroutines-channels-select](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines-channels-select)

    _An example of concurrency and message passing via channels in go._

  * [goroutines-waitgroup](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines-waitgroup)

    _An example of concurrency using a waitgroup._

  * [goroutines-worker-pools](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines-worker-pools)

    _An example of concurrency using a worker pool with goroutines and channels._

  * [async-channel-no-waiting](https://github.com/JeffDeCola/my-go-examples/tree/master/async-channel-no-waiting)

    _An example of a gorouting asynchronously sending data (via a channel) to a function that
    uses the latest data (if available) and does not wait._

* MESSINGING

  * [protobuf](https://github.com/JeffDeCola/my-go-examples/tree/master/protobuf)

    _Protocol buffers serialize structured data, useful for messaging._

  * [protobuf-NATS](https://github.com/JeffDeCola/my-go-examples/tree/master/protobuf-NATS)

    _Sends a protobuf msg over NATS from a client to a server._

* API

  * [http-GET-POST](https://github.com/JeffDeCola/my-go-examples/tree/master/http-GET-POST)

    _Using http package to GET or POST (update via json)._

  * [OAuth-2.0-google-cloud-storage-api](https://github.com/JeffDeCola/my-go-examples/tree/master/OAuth-2.0-google-cloud-storage-api)

    _Using OAuth 2.0 to access a users google-cloud-storage (based on scopes) via googles api._

  * [OAuth-2.0-google-cloud-storage-api-over-NATS](https://github.com/JeffDeCola/my-go-examples/tree/master/OAuth-2.0-google-cloud-storage-api-over-NATS)

    _Using OAuth 2.0 (frontend and backend via protobuf over NATS) to
    access a users google cloud storage (based on scopes) via googles api._

  * [track-email-mailgun-api](https://github.com/JeffDeCola/my-go-examples/tree/master/track-email-mailgun-api)

    _Send and track an email using mailgun over their go client api_

  * [something-youttube-content-id-api](https://github.com/JeffDeCola/my-go-examples/tree/master/something-youtube-content-id-api)

    _TBD_

  * [something-youttube-data-api-v3](https://github.com/JeffDeCola/my-go-examples/tree/master/something-youtube-data-api-ve)

    _TBD_

  * [something-RESTful](https://github.com/JeffDeCola/my-go-examples/tree/master/something-RESTful)

    _TBD_

  * [simple-webserver-with-REST](https://github.com/JeffDeCola/my-go-examples/tree/master/simple-webserver-with-REST)

    _Adding REST to the [simple-webserver](https://github.com/JeffDeCola/my-go-examples/tree/master/simple-webserver)._

* WEBSERVER

  * [simple-webserver](https://github.com/JeffDeCola/my-go-examples/tree/master/simple-webserver)

    _Using the http package to build a simple webserver._

* DATABASE

  * [something-postgreSQL](https://github.com/JeffDeCola/my-go-examples/tree/master/something-postgreSQL)

    _TBD_

* BASIC PROGRAMMING

  * [json](https://github.com/JeffDeCola/my-go-examples/tree/master/json)

    _Encode a struct to json and decode back to a struct._

  * [gotests-complex-function](https://github.com/JeffDeCola/my-go-examples/tree/master/gotests-complex-function)

    _Testing a function with complex inputs and outputs._

  * [gomock](https://github.com/JeffDeCola/my-go-examples/tree/master/gomock)

    _A helloween theme is used for gomock on an interface for unit testing._

  * [read-file](https://github.com/JeffDeCola/my-go-examples/tree/master/read-file)

    _Reading a file a few different ways._

  * [linked-list](https://github.com/JeffDeCola/my-go-examples/tree/master/linked-list)

    _An example of a Singly Linked List (i.e. using just a head pointer)._

  * [pointers](https://github.com/JeffDeCola/my-go-examples/tree/master/pointers)

    _An example of a pointer to a struct._*

  * [structs-methods-interfaces](https://github.com/JeffDeCola/my-go-examples/tree/master/structs-methods-interfaces)

    _An example of structs, methods and interfaces._

  * [recursion](https://github.com/JeffDeCola/my-go-examples/tree/master/recursion)

    _A function calling itself to make a fibonacci series._

  * [logging-error-handling](https://github.com/JeffDeCola/my-go-examples/tree/master/logging-error-handling)

    _Logging and error handling._


## TESTED USING CONCOURSE

A Concourse CI Pipeline will automate unit testing and update the GitHub WebPage.

![IMAGE - my-go-examples concourse ci piepline - IMAGE](pics/my-go-examples-pipeline.jpg)

A _ci/.credentials.yml_ file needs to be created for your _slack_url_ and _repo_github_token_.

Use fly to upload the the pipeline file _ci/pipline.yml_ to Concourse:

```bash
fly -t ci set-pipeline -p my-go-examples -c ci/pipeline.yml --load-vars-from ci/.credentials.yml
```

## CONCOURSE RESOURCES IN PIPELINE

`my-go-examples` also contains a few extra concourse resources:

* A resource (_resource-slack-alert_) uses a [docker image](https://hub.docker.com/r/cfcommunity/slack-notification-resource)
  that will notify slack on your progress.
* A resource (_resource-repo-status_) use a [docker image](https://hub.docker.com/r/dpb587/github-status-resource)
  that will update your git status for that particular commit.

The above resources can be removed from the pipeline.
