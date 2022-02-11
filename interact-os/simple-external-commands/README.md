# simple-external-commands

_Run a few os commands using the `exec` package._

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/interact-os/simple-os-interactions#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/interact-os/simple-os-interactions#run)

Documentation and references,

* Refer to the
  [exec](https://pkg.go.dev/exec)
  package for more info
* This repos [github webpage](https://jeffdecola.github.io/my-go-examples/)

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
go run simple-external-commands.go
```
