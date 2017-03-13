package main

import (
	"fmt"
)

type request struct {
	name string
}

type logic func(request)

var registry = map[string]logic{}

func addLogicForString(s string, l logic) {
	registry[s] = l
}

func callLogicForString(s string) {
	registry[s](request{name: s})
}

/*----------------------------------------*/

func sayHello(r request) {
	fmt.Printf("Hello, %s\n", r.name)
}

func monkeypoop(asdf request) {
	fmt.Printf("you suck")

}

/*----------------------------------------*/

func main() {
	addLogicForString("bob", sayHello)
	addLogicForString("jeff", monkeypoop)
	fmt.Printf("%+v\n", registry)
	//callLogicForString("bob")
//	callLogicForString("jeff")
}
