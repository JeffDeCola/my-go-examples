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
