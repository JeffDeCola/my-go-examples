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
}

// InsertAfter - provide the Value name to insert after
func (l *SinglyList) InsertAfter(n *Node, v Value) {
	if l.tail == n {
		l.Push(v)
		return
	}
	n.next = &Node{Value: v, next: n.next}
}

// Push a node onto a linked-list
func (l *SinglyList) Push(v Value) {
	// Create a node
	n := &Node{Value: v, next: nil}
	// Update Next
	if l.tail != nil {
		l.tail.next = n
	}
	// Always update tail
	l.tail = n
	// Set first Node
	if l.head == nil {
		l.head = n
	}
}

func stupid() {

	l := new(SinglyList)

	// PUSH Create another node
	l.Push(Value{name: "Jeff", age: 48, gender: "male"})
	l.Push(Value{name: "Clif", age: 32, gender: "male"})
	l.Push(Value{name: "Julie", age: 22, gender: "female"})
	l.Push(Value{name: "Larry", age: 25, gender: "male"})

	for e := l.head; e != nil; e = e.next {
		fmt.Println("ELEMENT", e.name, e.age, e.gender)
		if e.name == "Larry" {
			l.InsertAfter(e, Value{name: "Fred", age: 25, gender: "male"})
		}
	}

}
