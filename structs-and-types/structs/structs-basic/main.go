package main

import "fmt"

type thing struct {
	a int
	s string
}

func main() {

	// Field-name literal
	t1 := thing{
		a: 4,
		s: "whatever",
	}

	// Zero value (each field gets its type's zero: int 0, string "")
	var t2 thing

	// Pointer to a struct
	t3 := &thing{a: 7, s: "happy"}

	// Print the structs
	fmt.Println("PRINT THE STRUCTS")
	fmt.Printf("    t1 is %+v\n", t1) // {a:4 s:whatever}
	fmt.Printf("    t2 is %+v\n", t2) // {a:0 s:}
	fmt.Printf("    t3 is %+v\n", t3) // &{a:7 s:happy}

	// Access a field
	fmt.Println("ACCESS A FIELD")
	total := t1.a + 5
	fmt.Printf("    t1.a %d plus 5 is %d\n", t1.a, total)

	// Modify a field
	fmt.Println("MODIFY A FIELD")
	t1.s = "smile"
	fmt.Printf("    t1 modified is %+v\n", t1) // {a:4 s:smile}

}
