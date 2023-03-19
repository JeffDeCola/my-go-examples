# SIMPLE C CODE EXAMPLE

_A very simple example to show you how to write a c function in go._

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/cgo/c-code/simple-c-code#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/cgo/c-code/simple-c-code#run)

## OVERVIEW

Here is a very simple c code example,

```c
int Add(int a, int b){
    return a+b;
```

To implement in go, use `import "C"`,

```go
package main

/*
int Add(int a, int b){
    return a+b;
}
*/
import "C"
import "fmt"

func main() {

    a := C.int(10)
    b := C.int(20)
    c := C.Add(a, b)
    fmt.Printf("I'm printing the sum from go: %v\n", c)

}
```

## RUN

```bash
go run main.go
```

The output is,

```txt
I'm printing the sum from go: 30
```
