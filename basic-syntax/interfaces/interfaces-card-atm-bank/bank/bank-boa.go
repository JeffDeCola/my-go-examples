package bank

import (
	"fmt"

	"github.com/JeffDeCola/my-go-examples/basic-syntax/interfaces/interfaces-card-atm-bank/card"
)

type boaBank struct {
	BranchName string
	Accounts   map[string]int
}

// Balance returns the balance
func (a *boaBank) Balance(name string) int {
	fmt.Println("hi", name, a.Accounts[name])
	return 100
}

// Deposit deposits int the Account
func (a *boaBank) Deposit(d int) int {
	fmt.Println("hi")
	return 100
}

// Withdraw withdraws from the account
func (a *boaBank) Withdraw(w int) int {
	fmt.Println("hi")
	return 100
}

// IssueCard makes new card
func (a *boaBank) IssueCard() card.Card {
	return card.NewBoACard()
}

// IssueCard makes new card
func (a *boaBank) AddAccount(name string) {
	// Put this in a database - but lets just use a map

}

// NewBoaBank creates an ATM instance
func NewBoaBank(name string) Networker {
	return &boaBank{BranchName: name}
}

// NewBoaAccount to get balance
func NewBoaAccount(n Networker, name string, deposit int) Networker {
	var a map[string]int
	a[name] = deposit
	return &boaBank{Accounts: a}
}
