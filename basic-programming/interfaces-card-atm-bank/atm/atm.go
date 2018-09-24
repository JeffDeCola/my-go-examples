package atm

import (
	"github.com/JeffDeCola/my-go-examples/basic-programming/interfaces-card-atm-bank/bank"
)

type atm interface {
	inserted()
	getBalance(bank.Bank)
	deposited(bank.Bank, int) int
	withdrawn(bank.Bank, int) int
	ejected()
}

// The interface accepts ANYTHING as long as that
// ANYTHING has a method attached to this interface

// InsertCard into atm
func InsertCard(a atm) {
	a.inserted()
}

// ShowBalance from bank
func ShowBalance(a atm, b bank.Bank) {
	a.getBalance(b)
}

// Deposit into bank
func Deposit(a atm, b bank.Bank, d int) int {
	return a.deposited(b, d)
}

// Withdraw from bank
func Withdraw(a atm, b bank.Bank, w int) int {
	return a.withdrawn(b, w)
}

// EjectCard from atm
func EjectCard(a atm) {
	a.ejected()
}
