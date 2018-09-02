// my-go-examples singly-linked-list.go

package main

import (
	"fmt"
)

// Value is the data of the linked-list
type Value struct {
	name   string
	age    int
	gender string
}

// Node is just a node of a Linked List
type Node struct {
	Value // embedded struct
	next  *Node
}

// SinglyList is the List itself
type SinglyList struct {
	head *Node
	tail *Node
}

// InsertAfter - Provide the Value.name to insert after
func (l *SinglyList) InsertAfter(n *Node, v Value) {
	if l.tail == n {
		l.Push(v)
		return
	}
	n.next = &Node{Value: v, next: n.next}
}

// Push a node onto a linked-list
func (l *SinglyList) Push(v Value) {
	// Create a node with next nil
	n := &Node{Value: v, next: nil}
	// Update Next except but first one
	if l.tail != nil {
		l.tail.next = n
	}
	// Always update tail
	l.tail = n
	// update head on First only
	if l.head == nil {
		l.head = n
	}
}

func main() {

	myList := new(SinglyList)

	// PUSH Create another node
	fmt.Println("Pushing 4 Nodes onto a myList")
	myList.Push(Value{name: "Jeff", age: 48, gender: "male"})
	myList.Push(Value{name: "Clif", age: 32, gender: "male"})
	myList.Push(Value{name: "Julie", age: 22, gender: "female"})
	myList.Push(Value{name: "Larry", age: 25, gender: "male"})

	// Print all the node Values on the list
	for n := myList.head; n != nil; n = n.next {
		fmt.Println("    NODE", n.name, n.age, n.gender)
	}

	// Insert NODE Fred after Larry
	fmt.Println("Insert NODE Fred after Larry")
	for n := myList.head; n != nil; n = n.next {
		if n.name == "Julie" {
			myList.InsertAfter(n, Value{name: "Fred", age: 25, gender: "male"})
		}
	}

	// Print all the node Values on the list
	for n := myList.head; n != nil; n = n.next {
		fmt.Println("    NODE", n.name, n.age, n.gender)
	}

}
