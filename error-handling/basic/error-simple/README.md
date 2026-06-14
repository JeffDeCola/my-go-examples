# ERROR SIMPLE EXAMPLE

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Error handling using the
[errors](https://pkg.go.dev/errors)
standard package._

tl;dr

```go
# BASIC - make an error, return it, handle it
    func checkFilename(f string) error {...}        // returns nil or an error

    err := checkFilename("")                        // caller gets the error
    if err != nil {                                 // ...and handles it ONCE
      // handle and return
    }

# WRAPPING - add context with %w as it travels up
    func loadConfig(f string) error {...}           // top - wraps, passes up
    return fmt.Errorf("loadConfig: %w", err)        // wrap with %w

    func readConfigFile(f string) error {...}       // middle - wraps, passes up
    return fmt.Errorf("readConfigFile: %w", err)    // wrap with %w

    func checkFilename(f string) error {...}        // bottom - ORIGINATES
    return errors.New("filename can not...")        //

    err := loadConfig("")                           // caller gets the whole chain
    if err != nil {                                 // ...handles it ONCE at top
      // error: loadConfig: readConfigFile: filename can not...
    }

# SENTINELS - a named error callers can check for by identity
    var ErrNotFound = errors.New("not found")       // named, up top

    return ErrNotFound                              // return it by identity
    return fmt.Errorf("loadConfig: %w", ...)        // ...or wrapped

    if errors.Is(err, ErrNotFound) {                // caller branches WHICH error
        // handle THIS error differently
    }
```

Examples

* [error-simple](https://github.com/JeffDeCola/my-go-examples/tree/master/error-handling/basic/error-simple)
  **YOU ARE HERE**
* [error-wrapping](https://github.com/JeffDeCola/my-go-examples/tree/master/error-handling/wrapping/error-wrapping)
* [error-sentinel](https://github.com/JeffDeCola/my-go-examples/tree/master/error-handling/sentinels/error-sentinel)

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/error-handling/basic/error-simple#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/error-handling/basic/error-simple#run)

Documentation and Reference

* [errors](https://pkg.go.dev/errors)
  standard package
* [effective go - errors](https://go.dev/doc/effective_go#errors)

## OVERVIEW

The basic error pattern in go - a function returns an `error` as its
last return value, `nil` when all is well. The caller checks it
immediately and handles it before moving on.

```go
func checkFilename(filename string) error {...}   // returns nil or an error
err := checkFilename("")                          // caller gets the error
if err != nil {                                   // ...and handles it ONCE
  // handle and return
}
```

There are three ways to make an error in go. This example covers the first;
wrapping gets its own example later in this section.

```go
errors.New("filename can not be empty")              // static string  ← THIS EXAMPLE
fmt.Errorf("filename too long, got %d chars", n)     // value in the message
fmt.Errorf("opening config: %w", err)                // wrapping (see error-wrapping)
```

## RUN

To run,

```bash
go run main.go
```

The output should look like,

```bash
PART 1 - GOOD - pass filename
    Everything is good
PART 2 - BAD - pass no filename
CheckFilename failed: filename can not be empty
exit status 1
```
