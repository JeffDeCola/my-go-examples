# http-GET-POST

`http-GET-POST` _is an example of
using the standard `net/http` package to `http.Get` and `http.PostForm`._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## http.Get

Uses `http.Get` and `ioutil.ReadAll`,

```go
url := "https://httpbin.org/post"

response, err := http.Get(url)

defer response.Body.Close()
contents, err := ioutil.ReadAll(response.Body)

fmt.Printf("\nContent is:\n\n%s\n", string(contents))
```

## http.PostForm

How to update over an http api using `http.PostForm`,

```go
formData := url.Values{
    "name": {"jeff"},
}
url := "https://httpbin.org/post"

response, err := http.PostForm(url, formData)

defer response.Body.Close()
contents, err := ioutil.ReadAll(response.Body)

fmt.Printf("\nPOST is:\n\n%s\n", string(contents))
```

## RUN

```go
go run get.go
go run postform.go
```
