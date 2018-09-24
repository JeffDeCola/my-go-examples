package bank

import (
	"fmt"
)

type chaseBank struct {
	Name     string
	Location string
}

// Balance returns the balance
func (a *chaseBank) Balance() {
	fmt.Println("hi")
}

// Deposit deposits int the Account
func (a *chaseBank) Deposit() int {
	fmt.Println("hi")
	return 100
}

// Withdraw withdraws from the account
func (a *chaseBank) Withdraw(w int) int {
	fmt.Println("hi")
	return 100
}

// NewchaseBank creates an ATM instance
func NewchaseBank() Bank {
	return &chaseBank{}
}
