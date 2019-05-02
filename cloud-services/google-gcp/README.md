# google-gcp example

`google-gcp` _is an example of
interacting with google's
[gcp](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-providers/google-cloud-platform-cheat-sheet)
using
[google-cloud-go](https://github.com/googleapis/google-api-go-client/tree/master)
client libraries._

Documentation and reference,

* My `gcp` cheat sheet is
  [here](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-providers/google-cloud-platform-cheat-sheet)
* [GCP Go Documentation](https://cloud.google.com/go/docs/)
* [Getting Started with the Google APIs for Go](https://github.com/googleapis/google-api-go-client/blob/master/GettingStarted.md)
* [Complete list of packages](https://github.com/googleapis/google-api-go-client/tree/master)
* [GoDoc Package for Compute Engine API](https://godoc.org/google.golang.org/api/compute/v1)
* [GoDoc Package for google oath2](https://godoc.org/golang.org/x/oauth2/google)

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

The `gcp` SDK for go provides APIs and utilities that developers can use
to build go applications that use `gcp` services.

## GO GET

Package `compute` provides access to the Compute Engine API.

Package `oauth2/google` provides support for making Oauth2
authorized and authenticated HTTP requests to Google APIs.

```bash
go get -u google.golang.org/api/compute/v1
go get -u golang.org/x/oauth2/google
```

And import them as,

```go
import "google.golang.org/api/compute/v1"
import "golang.org/x/oauth2/google"
```

## INSTANTIATING

Each API has a,

* `New` function taking an `*http.Client` and
* Returns an API-specific `*Service.`

For example,

```go
service, err := compute.New(httpClient)
```

The HTTP client you pass in to the service must be one
that automatically adds Google-supported Authorization
information to the requests.

## CREDENTIALS - SERVICE ACCOUNTS

To get a service for `gce` using your google service account
would be as follows,

```go
ctx := context.Background()

// Reads credentials from the specified path.
computeService, err := compute.NewService(ctx, option.WithCredentialsFile(jsonPath))
if err != nil {
    log.Fatal(err)
}
```

## LETS USE THE SERVICE

Now that we have the service, lets use it.

Use the service to get Image list,

```go
req := computeService.Images.List(project)
```

Or use the service to get a instance list,

```go
req := computeService.Instances.List(project, zone)
```
