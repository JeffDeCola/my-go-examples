# json example

`json` _is an example of a struct encoded to json and decoded
back to a struct._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## ENCODE

Encodes json to a string from a struct.

struct -> encode -> byte -> string

`json.Marshal` encodes a struct to json data.

## DECODE

Decodes json to a struct from a string.

string -> byte -> decode -> struct

`json.Unmarshal` decodes json data to a struct.

## RUN

```bash
go run json.go
```
