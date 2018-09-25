package bank

import (
	"fmt"
)

type wellsFargoBank struct {
	Name     string
	Location string
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
func (a *wellsFargoBank) IssueCard() Card {
	return MakeNewCard()
}

// NewwellsFargoBank creates an ATM instance
func NewWellsFargoBank() BankNetworker {
	return &wellsFargoBank{}
}
