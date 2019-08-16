# simple-c-code example

`simple-c-code` _is an example of
of how to use c code in go._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## THE C CODE

Here is a very simple c code example,

```c
int Add(int a, int b){
    return a+b;
```

## IMPLEMENT WITH GO

To implement using go use `import "C"`,

```go
package main

//int Add(int a, int b){
//    return a+b;
//}
import "C"
import "fmt"

func main() {
    a := C.int(10)
    b := C.int(20)
    c := C.Add(a, b)
    fmt.Println(c) // 30
}
```

## RUN

```go
go run simple-c-code.go
```
