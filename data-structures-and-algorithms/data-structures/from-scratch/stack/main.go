// STACK EXAMPLE
//
// A LIFO (last-in first-out) built from a slice.
//

package main

import "fmt"

type stack struct {
	items []int
}

func (s *stack) push(x int) {
	s.items = append(s.items, x)
}

func (s *stack) pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false // empty - nothing to pop
	}
	top := s.items[len(s.items)-1]     // read tail
	s.items = s.items[:len(s.items)-1] // chop tail
	return top, true
}

func main() {

	s := &stack{}

	s.push(1)
	s.push(2)
	s.push(3)
	fmt.Println("Stack:", s.items)

	v, _ := s.pop()
	fmt.Println("Pop:  ", v) // 3
	v, _ = s.pop()
	fmt.Println("Pop:  ", v) // 2
	v, _ = s.pop()
	fmt.Println("Pop:  ", v) // 1

	_, ok := s.pop()
	fmt.Println("Pop empty:", ok) // false

}
