# SIMPLE EXTERNAL COMMANDS EXAMPLE

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Run a few os commands using the
[os/exec](https://pkg.go.dev/os/exec)
standard package._

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/go-runtime/interact-host-os/simple-os-interactions#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/go-runtime/interact-host-os/simple-os-interactions#run)

Documentation and Reference

* [os/exec](https://pkg.go.dev/os/exec)
  standard package

## OVERVIEW

Package exec runs external commands.

For example,

```go
// ls -s
out, err := exec.Command("ls", "-l").Output()
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(out))
```

## RUN

Run,

```bash
go run main.go
```
