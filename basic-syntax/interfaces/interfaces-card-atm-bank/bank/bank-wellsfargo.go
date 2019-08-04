package bank

import (
	"fmt"

	"github.com/JeffDeCola/my-go-examples/basic-programming/interfaces-card-atm-bank/card"
)

type wellsFargoBank struct {
	Name        string
	Location    string
	wellsstruct string
}

// Balance returns the balance
func (a *wellsFargoBank) Balance() int {
	fmt.Println("hi")
	return 100
}

// Deposit deposits int the Account
func (a *wellsFargoBank) Deposit(d int) int {
	fmt.Println("hi")
	return 100
}

// Withdraw withdraws from the account
func (a *wellsFargoBank) Withdraw(w int) int {
	fmt.Println("hi")
	return 100
}

// IssueCard makes new card
func (a *wellsFargoBank) IssueCard() card.Card {
	return card.NewWellsFargoCard()
}

// NewwellsFargoBank creates an ATM instance
func NewWellsFargoBank() BankNetworker {
	return &wellsFargoBank{}
}
