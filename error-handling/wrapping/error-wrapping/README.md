# ERROR WRAPPING EXAMPLE

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Wrapping errors with %w to add context as they propagate up the call stack
using the
[errors](https://pkg.go.dev/errors)
standard package._

tl;dr

```text
# BASIC - make an error, return it, handle it
    func checkFilename(f string) error {...}      // returns nil or an error

    err := checkFilename("")                      // caller gets the error
    if err != nil {                               // ...and handles it ONCE
      // handle and return
    }

# WRAPPING - add context with %w as it travels up
    func loadConfig(f string) error {...}         // top - wraps, passes up
    return fmt.Errorf("loadConfig: %w", err)      // wrap with %w

    func readConfigFile(f string) error {...}     // middle - wraps, passes up
    return fmt.Errorf("readConfigFile: %w", err)  // wrap with %w

    func checkFilename(f string) error {...}      // bottom - ORIGINATES
    return errors.New("filename can not...")      //

    err := loadConfig("")                         // caller gets the whole chain
    if err != nil {                               // ...handles it ONCE at the top
      // error: loadConfig: readConfigFile: filename can not...
    }

# SENTINELS - a named error callers can check for by identity
    var ErrNotFound = errors.New("not found")     // named, up top

    return ErrNotFound                            // return it by identity
    return fmt.Errorf("loadConfig: %w", ...)      // ...or wrapped

    if errors.Is(err, ErrNotFound) {              // caller branches on WHICH error
        // handle THIS error differently
    }
```

Examples

* [error-simple](https://github.com/JeffDeCola/my-go-examples/tree/master/error-handling/basic/error-simple)
* [error-wrapping](https://github.com/JeffDeCola/my-go-examples/tree/master/error-handling/wrapping/error-wrapping)
  **YOU ARE HERE**
* [error-sentinels](https://github.com/JeffDeCola/my-go-examples/tree/master/error-handling/sentinels/error-sentinels)

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/error-handling/wrapping/error-wrapping#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/error-handling/wrapping/error-wrapping#run)

Documentation and Reference

* [errors](https://pkg.go.dev/errors)
  standard package
* [working with errors in go 1.13](https://go.dev/blog/go1.13-errors)

## OVERVIEW

When an error happens deep in a call stack, each function it passes through
wraps it with `%w` to add context about what that level was doing. The result
is a chain - the trail of how you got there in front, the root cause at the end.

```go
func loadConfig(f string) error {...}      // top    - wraps, passes up
func readConfigFile(f string) error {...}  // middle - wraps, passes up
func checkFilename(f string) error {...}   // bottom - ORIGINATES the error

return fmt.Errorf("loadConfig: %w", err)   // wrap with %w as it goes UP
return errors.New("filename can not...")   // originate at the bottom
```

The deepest level originates the error with a bare description of the problem.
Each level above wraps it, adding its own name as context. The caller at the
top handles it once - and because the chain already names every level, it just
prints the error.

```bash
error: loadConfig: readConfigFile: filename can not be empty
```

Read it like a trail - `loadConfig` was running, which called
`readConfigFile`, and the root cause was an empty filename.

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
error: loadConfig: readConfigFile: filename can not be empty
exit status 1
```
