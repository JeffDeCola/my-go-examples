// MAPS EXAMPLE
//
// An unordered collection of key-value pairs (a hash table).
//

package main

import "fmt"

func main() {

	// DECLARE (nil)
	var a map[string]int
	fmt.Println("Declare:              ", a, a == nil)

	// MAKE (not nil)
	a = make(map[string]int)
	fmt.Println("Make:                 ", a, a == nil)

	// ASSIGN
	a["go"] = 2009
	a["rust"] = 2010
	fmt.Println("Assign:               ", a)

	// DECLARE & INITIALIZE
	var b = map[string]int{"go": 2009, "rust": 2010}
	fmt.Println("Declare & Initialize: ", b)
	c := map[string]int{"c": 1972, "python": 1991}
	fmt.Println("Declare & Initialize: ", c)

	// READ (comma-ok)
	v, ok := a["go"] // value, present?
	fmt.Println("Read go:              ", v, ok)
	w, ok := a["java"] // absent key -> zero value, false
	fmt.Println("Read java:            ", w, ok)

	// DELETE
	delete(a, "rust") // delete(name, key)
	fmt.Println("Delete rust:          ", a)

}
