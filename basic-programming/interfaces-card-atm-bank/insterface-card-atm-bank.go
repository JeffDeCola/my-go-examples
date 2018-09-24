package main

import (
	"fmt"

	"github.com/JeffDeCola/my-go-examples/basic-programming/interfaces-card-atm-bank/atm"
	"github.com/JeffDeCola/my-go-examples/basic-programming/interfaces-card-atm-bank/bank"
)

func main() {

	// WHAT ATM CARD ARE YOU USING
	// var c = &atm.BoACard{Name: "jeff", Bank: "BoA", AccountNumber: 001}
	// var c = &atm.ChaseCard{NameOnCard: "jeff", Bank: "Chase"}
	var c = &atm.WellsFargo{Bank: "BoA"}

	// WHAT ATM ARE YOU USING
	// var a = &bank.MainStATM{Name: "MainSt", Location: "Los Angeles"}
	// var a = &bank.DrugStoreATM{Location: "Elm Street"}
	var a = &bank.BarATM{ATMName: "BarATM"}

	fmt.Println("Insert your Card")

	// Every one of these function uses an interface for atm card, bank or both.
	atm.InsertCard(c)
	atm.ShowBalance(c, a)
	//fmt.Printf("You are deposited a total of %v in cash\n", atm.Deposit(c, a, 100))
	fmt.Printf("You withdrew a total of %v in cash\n", atm.Withdraw(c, a, 50))
	atm.EjectCard(c)

}
