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
