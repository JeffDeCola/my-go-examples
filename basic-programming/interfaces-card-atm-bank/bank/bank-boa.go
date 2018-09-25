package bank

import (
	"fmt"

	"github.com/JeffDeCola/my-go-examples/basic-programming/interfaces-card-atm-bank/card"
)

type boABank struct {
	Name          string
	boABankStruct string
}

// Balance returns the balance
func (a *boABank) Balance() int {
	fmt.Println("hi")
	return 100
}

// Deposit deposits int the Account
func (a *boABank) Deposit(d int) int {
	fmt.Println("hi")
	return 100
}

// Withdraw withdraws from the account
func (a *boABank) Withdraw(w int) int {
	fmt.Println("hi")
	return 100
}

// IssueCard makes new card
func (a *boABank) IssueCard() card.Card {
	return card.NewBoACard()
}

// NewBoABank creates an ATM instance
func NewBoABank() BankNetworker {
	return &boABank{}
}
