// my-go-examples simple-c-code.go

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
	fmt.Printf("Print returned sum from go: %v\n", c) // 30
}
