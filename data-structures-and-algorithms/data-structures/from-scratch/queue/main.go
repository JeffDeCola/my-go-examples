// QUEUE EXAMPLE
//
// A FIFO (first-in first-out) built from a slice.
//

package main

import "fmt"

type queue struct {
	items []int
}

func (q *queue) enqueue(i int) {
	q.items = append(q.items, i)
}

func (q *queue) dequeue() (int, bool) {
	if len(q.items) == 0 {
		return 0, false // empty - nothing to pop
	}
	top := q.items[0]     // read head
	q.items = q.items[1:] // chop head
	return top, true
}

func main() {

	q := &queue{}

	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	fmt.Println("Queue:    ", q.items)

	v, _ := q.dequeue()
	fmt.Println("Dequeue:  ", v) // 1
	v, _ = q.dequeue()
	fmt.Println("Dequeue:  ", v) // 2
	v, _ = q.dequeue()
	fmt.Println("Dequeue:  ", v) // 3
	_, ok := q.dequeue()
	fmt.Println("Dequeue empty:", ok) // false

}
