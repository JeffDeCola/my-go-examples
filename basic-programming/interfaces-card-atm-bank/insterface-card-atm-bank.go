package main

import (
	"fmt"

	"github.com/JeffDeCola/my-go-examples/basic-programming/interfaces-card-atm-bank/atm"
	"github.com/JeffDeCola/my-go-examples/basic-programming/interfaces-card-atm-bank/card"
)

func main() {
	w := 0

	fmt.Println("Get your atm card")
	c := card.ATMCard{"Jeff", "BoA", 1001}

	fmt.Println("Walk up to the ATM Machine")
	m := atm.AtmMachine{}
	var a atm.ATM = &m

	fmt.Println("Insert Card", c)
	a.InsertCard(c)

	fmt.Println("Show Balance", a.ShowBalance())
	fmt.Print("How much would you like to withdraw? ")
	fmt.Scan(&w)
	a.Withdraw(w)

}
