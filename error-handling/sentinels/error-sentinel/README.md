# ERROR SENTINEL EXAMPLE

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_A sentinel is a named error that callers can check for by identity
using the
[errors](https://pkg.go.dev/errors)
standard package._

tl;dr

```go
// BASIC - make an error, return it, handle it
    func checkFilename(f string) error {...}          // returns nil or an error
        return errors.New("filename....")

    err := checkFilename("")                          // caller gets the error
    if err != nil {                                   // ...and handles it ONCE
        fmt.Println("CheckFilename failed:", err)
        os.Exit(1)
    }

// WRAPPING - add context with %w as it travels up
    func loadConfig(f string) error {...}             // top - wraps, passes up
        return fmt.Errorf("loadConfig: %w", err)      // wrap with %w

    func readConfigFile(f string) error {...}         // middle - wraps, passes up
        return fmt.Errorf("readConfigFile: %w", err)  // wrap with %w

    func checkFilename(f string) error {...}          // bottom - ORIGINATES
        return errors.New("filename can not...")      //

    err := loadConfig("")                             // caller gets the whole chain
    if err != nil {                                   // ...handles it ONCE at top
        fmt.Println("error:", err)
        os.Exit(1)
    }

// SENTINELS - a named error callers can check for by identity
    var ErrFileNotFound = errors.New("not found")     // named, up top

    return ErrFileNotFound                            // return it by identity
    return fmt.Errorf("loadConfig: %w", ...)          // ...or wrapped

    if errors.Is(err, ErrFileNotFound) {
        fmt.Println("no file - using default")        // Try to recover
        filename = "data.txt"
        ...
    } else if err != nil {                            // All other errors
        fmt.Println("error:", err)
        os.Exit(1)
    }
```

Examples

* [error-simple](https://github.com/JeffDeCola/my-go-examples/tree/master/error-handling/basic/error-simple)
* [error-wrapping](https://github.com/JeffDeCola/my-go-examples/tree/master/error-handling/wrapping/error-wrapping)
* [error-sentinel](https://github.com/JeffDeCola/my-go-examples/tree/master/error-handling/sentinels/error-sentinel)
  **YOU ARE HERE**

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/error-handling/sentinels/error-sentinel#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/error-handling/sentinels/error-sentinel#run)

Documentation and Reference

* [errors](https://pkg.go.dev/errors)
  standard package
* [working with errors in go 1.13](https://go.dev/blog/go1.13-errors)

## OVERVIEW

A sentinel is a named, package-level error because it has a name, a caller
can check for that _specific_ error with `errors.Is` and respond to it
differently from any other failure.

`errors.Is` digs through `%w` wrapping, so a sentinel is still recognized even
when it has been wrapped with context several layers up the call stack.

Note the `else if err != nil` branch in this small example `checkFilename`
only fails one way, so it never fires, but it shows the full shape: recover
from the error you know, bail on the ones you don't.

## RUN

To run,

```bash
go run main.go
```

The output should look like,

```bash
no file - using default data.txt
carrying on with: data.txt
carrying on with: something.txt
```
