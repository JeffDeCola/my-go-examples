
# GO EXAMPLES

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

* [methods-and-interfaces](https://github.com/JeffDeCola/my-go-examples/tree/master/methods-and-interfaces)

   _An example of methods and interfaces._

* [recursion](https://github.com/JeffDeCola/my-go-examples/tree/master/recursion)

   _A function calling itself to make a fibonacci series._

* [logging-error-handling](https://github.com/JeffDeCola/my-go-examples/tree/master/logging-error-handling)

   _Logging and error handling._

* [simple-webserver](https://github.com/JeffDeCola/my-go-examples/tree/master/simple-webserver)

   _Using the http package to build a simple webserver._

* [simple-webserver-with-REST](https://github.com/JeffDeCola/my-go-examples/tree/master/simple-webserver-with-REST)

   _Adding REST to the simple-webserver._

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
