# simple-c-code-using-stdio example

`simple-c-code-using-stdio` _is an example of
how to write a c function in go and using stdio.h._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## THE C CODE

Here is a very simple c code example,

```c
int Add(int a, int b, char * name){
    int s;
    s = a+b;
    printf("Hi %s. ", name);
    printf("Print the sum from C: %d\n", s);
    return s;
}
```

## IMPLEMENT WITH GO

This example will use stdio.h for printf.
To implement using go use `import "C"`,

```go
package main

/*
#include <stdio.h>

int Add(int a, int b, char * name){
    int s;
    s = a+b;
    printf("Hi %s. ", name);
    printf("Print the sum from C: %d\n", s);
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
    fmt.Printf("Print returned sum from go: %v\n", c) // 30
}
```

## RUN

```bash
go run simple-c-code-using-stdio.go
```
