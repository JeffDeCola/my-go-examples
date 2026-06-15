// generics-type
//
// A generic type uses a type parameter so one definition serves every element
// type, instead of writing one type per element._
//

package main

import "fmt"

// stack holds elements of ANY type T
type stack[T any] struct {
	items []T
}

// push adds an element (pointer receiver - it modifies the stack)
func (s *stack[T]) push(item T) {
	s.items = append(s.items, item)
}

// pop removes and returns the top element
func (s *stack[T]) pop() T {
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last
}

func main() {

	// A stack of ints
	fmt.Println("STACK OF INTS")
	var si stack[int]
	si.push(10)
	si.push(20)
	fmt.Printf("    stack is  %v\n", si.items)
	fmt.Printf("    pop       %v\n", si.pop())
	fmt.Printf("    stack is  %v\n", si.items)

	// The SAME type, now holding strings
	fmt.Println("STACK OF STRINGS")
	var ss stack[string]
	ss.push("a")
	ss.push("b")
	fmt.Printf("    stack is  %v\n", ss.items)
	fmt.Printf("    pop       %v\n", ss.pop())
	fmt.Printf("    stack is  %v\n", ss.items)
}
