// module-with-remote-package
//
// A go module with a remote package.
//

package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {

	// Create a Universally Unique Identifier
	id := uuid.New()
	fmt.Println(id)

}
