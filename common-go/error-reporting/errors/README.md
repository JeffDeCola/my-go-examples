# errors

_Error Handling using the standard `error` package._

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/error-reporting/errors#overview)
  * [DEFINE YOUR OWN ERROR](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/error-reporting/errors#define-your-own-error)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/error-reporting/errors#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/error-reporting/errors#test)

Documentation and references,

* Refer to the
  [errors](https://pkg.go.dev/errors)
  package for more info
* This repos [github webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

Go doesn’t have exceptions, so it doesn’t have try, catch or anything similar.
So how do we handle errors? **With Multiple return values**.

### DEFINE YOUR OWN ERROR

It's very easy. For example,

```go
var ErrIncorrectAnswer = errors.New("the answer is incorrect")

if answer != 4 {
    return fmt.Errorf("%d was provided: %w", answer, ErrIncorrectAnswer)
}
```

## RUN

The programs asks you what is 2+2 and depending on your answer will
say you are incorrect, correct and/or give you an error.

Run,

```bash
go run errors.go
```

The output may look like,

```bash
What is 2+2?: 4
YOUR ANSWER 4 IS CORRECT!
------------------------
What is 2+2?: 5
INCORRECT!!
ERRO[0002] Error with answer: error calling b from a: error calling c from b: 5 was provided: the answer is incorrect 
------------------------
What is 2+2?: cat
Not an int
ERRO[0004] Error with answer: error calling b from a: unable to convert in b: strconv.Atoi: parsing "cat": invalid syntax 
------------------------
What is 2+2?: stop
Done
```

Notice how it traces the error back to your original function.

## TEST

To create _test files,

```bash
gotests -w -all errors.go
```

To unit test the code,

```bash
go test -cover ./... 
```
