package bank

import (
	"fmt"
)

type wellsFargoBank struct {
	Name     string
	Location string
}

// Balance returns the balance
func (a *wellsFargoBank) Balance() {
	fmt.Println("hi")
}

// Deposit deposits int the Account
func (a *wellsFargoBank) Deposit() int {
	fmt.Println("hi")
	return 100
}

// Withdraw withdraws from the account
func (a *wellsFargoBank) Withdraw(w int) int {
	fmt.Println("hi")
	return 100
}

// NewwellsFargoBank creates an ATM instance
func NewwellsFargoBank() Bank {
	return &wellsFargoBank{}
}
