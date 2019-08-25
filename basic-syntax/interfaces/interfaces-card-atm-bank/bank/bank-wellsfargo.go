package bank

import (
	"fmt"

	"github.com/JeffDeCola/my-go-examples/basic-syntax/interfaces/interfaces-card-atm-bank/card"
)

type wellsFargoBank struct {
	BranchName string
	Location   string
}

// Balance returns the balance
func (a *wellsFargoBank) Balance(name string) int {
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

// NewWellsFargoBank creates an ATM instance
func NewWellsFargoBank() Networker {
	return &wellsFargoBank{}
}
