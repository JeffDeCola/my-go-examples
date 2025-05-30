# SIMPLE C CODE USING STDIO EXAMPLE

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_A c function in go using stdio.h for printf._

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/cgo/c-code/simple-c-code-using-stdio#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/cgo/c-code/simple-c-code-using-stdio#run)

## OVERVIEW

Here is a very simple c code example,

```c
int Add(int a, int b, char * name){
    int s;
    s = a+b;
    printf("Hi %s. ", name);
    printf("I'm printing the sum from C:  %d\n", s);
    return s;
}
```

This example will use stdio.h for printf.
To implement using go use `import "C"`,

```go
package main

/*
#include <stdio.h>

int Add(int a, int b, char * name){
    int s;
    s = a+b;
    printf("Hi %s,\n", name);
    printf("I'm printing the sum from C:  %d\n", s);
    return s;
}
*/
import "C"
import "fmt"

func main() {

    a := C.int(10)
    b := C.int(20)
    name := C.CString("jeff")
    c := C.Add(a, b, name)
    fmt.Printf("I'm printing the sum from go: %v\n", c)

}
```

## RUN

```bash
go run main.go
```

The output is,

```txt
Hi jeff,
I'm printing the sum from C:  30
I'm printing the sum from go: 30
```
