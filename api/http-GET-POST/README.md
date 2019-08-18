# http-GET-POST

`http-GET-POST` _is an example of
using the standard net/http package to GET or POST (update via json)._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## get

Uses `http.Get`  and `ioutil.ReadAll`:

```go
response, err := http.Get("http://www.jeffdecola.com/robots.txt")
defer response.Body.Close()
contents, err := ioutil.ReadAll(response.Body)
fmt.Printf("Content is:\n\n%s\n", string(contents))
```

## get-struct

Same as above but it writes to struct.

## post

How to update over an api.