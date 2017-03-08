# gotests-complex-function example

`gotests-complex-function` _is an example of testing a function with
complex inputs and outputs._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## THE FUNCTION

The function `Check()` has the following input/ouputs.

`func Check(input InputJSON, logger *log.Logger) (checkOutputJSON, error) {`

where,

```go
version struct {
    Ref string `json:"ref"`
}

InputJSON struct {
    Params  map[string]string `json:"params"`
    Source  map[string]string `json:"source"`
    Version version           `json:"version"`
}

checkOutputJSON []version

```

`TestCheck()` would then contain,

```go
{
    "test1",
    args{
        input: InputJSON{
            Params: map[string]string{"param1": "Hello Jeff", "param2":
            "How are you?"},
            Source: map[string]string{"source1": "sourcefoo1", "source2":
            "sourcefoo2"},
            Version: version{
                Ref: "456",
            },
        },
        logger: log.New(os.Stderr, "resource:", log.Lshortfile),
    },
    checkOutputJSON{
        version{Ref: "123"},
        version{Ref: "3de"},
        version{Ref: "456"},
        version{Ref: "777"},
    },
    false,
},
```

## GENERATE TEST FILE

```bash
gotests -w -all complex_function.go
```

## RUN TEST

```bash
go test -v -cover .
```