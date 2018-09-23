package atm

import (
	"fmt"

	"github.com/JeffDeCola/my-go-examples/basic-programming/interfaces-card-atm-bank/card"
)

// ATM Machine interface
type ATM interface {
	InsertCard(card.ATMCard)
	ShowBalance() int
	Withdraw(int)
}

type AtmMachine struct {
	cardName          string
	cardBank          string
	cardAccountNumber int
	cardBalance       int
}

func (a *AtmMachine) ShowBalance() int {
	if a.cardAccountNumber == 1001 {
		a.cardBalance = 100
		return a.cardBalance
	}
	return 0
}

func (a *AtmMachine) InsertCard(c card.ATMCard) {
	fmt.Printf("Welcome %v\n", c.Name)
	a.cardName = c.Name
	a.cardBank = c.Bank
	a.cardAccountNumber = c.AccountNumber
}

func (a *AtmMachine) Withdraw(w int) {
	fmt.Printf("Going to withdraw %v for you %v\n", w, a.cardName)
	a.cardBalance = a.cardBalance - w
	fmt.Printf("Your new balance is %v\n", a.cardBalance)
}
