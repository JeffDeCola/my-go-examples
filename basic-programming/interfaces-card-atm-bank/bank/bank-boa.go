package bank

import (
	"fmt"
)

type boABank struct {
	Name     string
	Location string
}

// Balance returns the balance
func (a *boABank) Balance() {
	fmt.Println("hi")
}

// Deposit deposits int the Account
func (a *boABank) Deposit() int {
	fmt.Println("hi")
	return 100
}

// Withdraw withdraws from the account
func (a *boABank) Withdraw(w int) int {
	fmt.Println("hi")
	return 100
}

// NewBoABank creates an ATM instance
func NewBoABank() Bank {
	return &boABank{}
}
