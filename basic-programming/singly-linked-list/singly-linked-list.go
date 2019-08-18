// my-go-examples singly-linked-list.go

package main

import (
	"fmt"
)

// SinglyList is the List itself
type SinglyList struct {
	head *Node
	tail *Node
}

// Node is the node of a Linked List
type Node struct {
	Value
	next *Node
}

// Value is the data of the linked-list
type Value struct {
	name   string
	age    int
	gender string
}

// Push a node onto a linked-list
func (l *SinglyList) Push(v Value) {

	// Declair & Assign data to type struct Node
	n := Node{Value: v, next: nil}
	// Declair a Pointer to that node
	myNode := &n

	fmt.Printf("    For the new Node Address %p ", myNode)
	fmt.Printf("adding Value=%v\n", v)
	// myList HEAD
	// FOR FIRST NODE ONLY - head is address of the first Node
	if l.head == nil {
		l.head = myNode
	}

	// NEXT (NOT FIRST NODE)
	// This is actually the previous node
	if l.tail != nil {
		l.tail.next = myNode
	}

	// myList TAIL
	// Always update tail when adding a Node
	l.tail = myNode
}

// InsertAfter - Provide the Value.name to insert after
func (l *SinglyList) InsertAfter(nodePtr *Node, v Value) {

	// If the last on the list
	if l.tail == nodePtr {

		l.Push(v)
		return
	}

	// To insert a node
	// Declare & Assign data to type struct Node
	n := Node{Value: v, next: nodePtr.next}
	fmt.Printf("    1. Create new node %v\n", n)
	fmt.Printf("       Notice the next pointer is %p\n", n.next)
	fmt.Printf("    2. Update the next pointer in %v from %p to %p\n", nodePtr.name, nodePtr.next, &n)
	// Declair a Pointer to that node
	nodePtr.next = &n
}

func main() {

	// Declair a pointer to a type struct SinglyList
	myList := new(SinglyList)
	fmt.Printf("\n    *myList{head,tail}=%v\n", *myList)

	// Declair & Assign data to type struct Value
	// This will be copied into struct (This stuff will just hang around)
	value1 := Value{name: "Jeff", age: 48, gender: "male"}
	value2 := Value{name: "Clif", age: 32, gender: "male"}
	value3 := Value{name: "Julie", age: 22, gender: "female"}
	value4 := Value{name: "Larry", age: 25, gender: "male"}
	value5 := Value{name: "Fred", age: 25, gender: "male"}

	// PUSH Node onto a linked list
	fmt.Println("Pushing 4 Nodes onto myList")
	myList.Push(value1)
	fmt.Printf("    *myList{head,tail} updated to %v\n", *myList)
	myList.Push(value2)
	fmt.Printf("    *myList{head,tail} updated to %v\n", *myList)
	myList.Push(value3)
	fmt.Printf("    *myList{head,tail} updated to %v\n", *myList)
	myList.Push(value4)
	fmt.Printf("    *myList{head,tail} updated to %v\n", *myList)

	// Print all the node Values on the list
	fmt.Println("Print all the node Values on the list")
	for nodePtr := myList.head; nodePtr != nil; nodePtr = nodePtr.next {
		fmt.Printf("    For Node Addres %p   ", nodePtr)
		fmt.Printf("Value=%6v next=%p\n", nodePtr.Value, nodePtr.next)
	}

	// Insert NODE Fred after Larry
	fmt.Println("Insert NODE Fred after Clif")
	for nodePtr := myList.head; nodePtr != nil; nodePtr = nodePtr.next {
		if nodePtr.name == "Clif" {
			fmt.Printf("    We found node Clif at Node Address %p, so lets insert %v\n", nodePtr, value5)
			myList.InsertAfter(nodePtr, value5)
		}
	}
	fmt.Printf("    *myList{head,tail} still the same %v\n", *myList)

	// Print all the node Values on the list
	fmt.Println("Print all the node Values on the list")
	for nodePtr := myList.head; nodePtr != nil; nodePtr = nodePtr.next {
		fmt.Printf("    For Node Addres %p   ", nodePtr)
		fmt.Printf("Value=%6v next=%p\n", nodePtr.Value, nodePtr.next)
	}

}
