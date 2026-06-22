// module-with-local-package
//
// A go module with a local package.
//

package main

import (
	"fmt"
	"module-with-local-package/mypackage"
)

func main() {

	sum := mypackage.Add(2, 2)
	fmt.Println(sum)

}
