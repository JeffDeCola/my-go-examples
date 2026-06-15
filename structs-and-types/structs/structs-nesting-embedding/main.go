// structs-nesting-embedding
//
// Nesting and embedding a struct inside another struct, and how field
// access differs between them.
//

package main

import "fmt"

type thing struct {
	a int
	s string
}

type bigNesting struct {
	b int
	t thing // This is nesting
}

type bigEmbedding struct {
	b     int
	thing // This is embedding
}

func main() {

	// Going to use field name literals
	t := thing{
		a: 5,
		s: "smile",
	}

	n := bigNesting{
		b: 10,
		t: t,
	}

	e := bigEmbedding{
		b:     10,
		thing: t,
	}

	// Print structs
	fmt.Println("PRINT BOTH STRUCTS - Same data, but note the field name")
	fmt.Printf("    Nesting:       n     is  %+v\n", n)
	fmt.Printf("    Embedding:     e     is  %+v\n", e)

	// THE DIFFERENCE - reaching field 'a'
	fmt.Println("ACCESS FIELD a - It's simpler with embedding")
	fmt.Printf("    Nesting:       n.t.a is  %d\n", n.t.a) // must go through field name t
	fmt.Printf("    Embedding:     e.a   is  %d\n", e.a)   // PROMOTED - direct!
}
