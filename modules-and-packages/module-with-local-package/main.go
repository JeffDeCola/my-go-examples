package main

import (
	"fmt"
	arithmetic "module-with-local-package/mypackage"
)

func main() {

	sum := arithmetic.Add(2, 2)
	fmt.Println(sum)

}
