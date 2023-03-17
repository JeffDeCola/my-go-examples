# ERROR EXAMPLE

_Error Handling using the standard `error` package._

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/error-reporting/error-example#overview)
* [PREREQUISITES](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/error-reporting/error-example#prerequisites)
* [DEFINE YOUR OWN ERROR](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/error-reporting/error-example#define-your-own-error)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/error-reporting/error-example#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/error-reporting/error-example#test)

Documentation and References

* Refer to the
  [errors](https://pkg.go.dev/errors)
  package for more info

## OVERVIEW

Go doesn’t have exceptions, so it doesn’t have try, catch or anything similar.
So how do we handle errors? **With Multiple return values**.

## PREREQUISITES

You will need the following go packages,

```bash
go get -u github.com/cweill/gotests/...
```

## DEFINE YOUR OWN ERROR

It's very easy. For example,

```go
var ErrIncorrectAnswer = errors.New("the answer is incorrect")

if answer != 4 {
    return fmt.Errorf("%d was provided: %w", answer, ErrIncorrectAnswer)
}
```

## RUN

The programs asks you what is 2+2 and depending on your answer will
say you are incorrect, correct and/or give you an error. Notice how it **propagates**
the error up the function calls.

Run,

```bash
go run main.go
```

The output may look like,

```bash
What is 2+2 (type stop to quit)? 4
YOUR ANSWER 4 IS CORRECT!
------------------------
What is 2+2 (type stop to quit)? 5
INCORRECT!
ERRO[0003] Error with answer: error calling checkNumber: 5 was provided: the answer is incorrect
------------------------
What is 2+2 (type stop to quit)? cat
Not an integer
ERRO[0005] Error with answer: unable to convert to integer: strconv.Atoi: parsing "cat": invalid syntax
------------------------
What is 2+2 (type stop to quit)? stop
Done
```

Notice how it traces the error back to your original function.

## TEST

To create _test files,

```bash
gotests -w -all main.go
```

To unit test the code,

```bash
go test -cover ./...
```
