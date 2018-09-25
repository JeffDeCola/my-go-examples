package bank

import (
	"fmt"

	"github.com/JeffDeCola/my-go-examples/basic-programming/interfaces-card-atm-bank/card"
)

type chaseBank struct {
	Name            string
	Location        string
	ChaseBankStruct string
}

// Balance returns the balance
func (a *chaseBank) Balance() int {
	fmt.Println("hi")
	return 100
}

// Deposit deposits int the Account
func (a *chaseBank) Deposit(d int) int {
	fmt.Println("hi")
	return 100
}

// Withdraw withdraws from the account
func (a *chaseBank) Withdraw(w int) int {
	fmt.Println("hi")
	return 100
}

// IssueCard makes new card
func (a *chaseBank) IssueCard() card.Card {
	return card.NewChaseCard()
}

// NewChaseBank creates an ATM instance
func NewChaseBank() BankNetworker {
	return &chaseBank{}
}
