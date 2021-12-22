// my-go-examples simple-c-code-using-stdio.go

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
