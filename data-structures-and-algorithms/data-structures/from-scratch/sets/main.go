// SETS EXAMPLE
//
// A collection of unique elements built from a map.
//

package main

import "fmt"

func main() {

	// MAKE
	// We pick struct so it saves memory, you could pick int but that takes up space
	rsvp := make(map[string]struct{})
	fmt.Println("Make map:             ", rsvp)

	// ADD (RSVP)
	// name[key] = struct{}{}  (value stores nothing) Will not take up memory
	rsvp["Jeff"] = struct{}{}
	rsvp["Larry"] = struct{}{}
	rsvp["Sam"] = struct{}{}
	rsvp["Jeff"] = struct{}{} // duplicate
	fmt.Println("RSVP (Jeff twice):    ", rsvp, "- len", len(rsvp))

	// READ - normal map use
	_, ok := rsvp["Larry"] // is Amy on the list?
	fmt.Println("Is Larry coming:      ", ok)
	_, ok = rsvp["Bob"] // bob never RSVP'd
	fmt.Println("Is Bob coming:        ", ok)

	// DELETE - normal map use
	delete(rsvp, "Sam") // delete(name, key)
	fmt.Println("DELETE Sam:           ", rsvp, "- len", len(rsvp))
}
