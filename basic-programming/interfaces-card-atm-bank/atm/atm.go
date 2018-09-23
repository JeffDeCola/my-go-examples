package atm

import (
	"fmt"

	"github.com/JeffDeCola/my-go-examples/basic-programming/interfaces-card-atm-bank/card"
)

// ATM Machine interface
type ATM interface {
	InsertCard(card.ATMCard) bool
	ShowBalance() int
	Withdraw(int)
}

// Machine is the data in the machine
type Machine struct {
	cardName          string
	cardBank          string
	cardAccountNumber int
	cardBalance       int
}

// ShowBalance will show your balance from ATMMachine
func (a *Machine) ShowBalance() int {
	if a.cardAccountNumber == 1001 {
		a.cardBalance = 100
		return a.cardBalance
	}
	return 0
}

// InsertCard will put card info into ATMMachine
func (a *Machine) InsertCard(c card.ATMCard) bool {
	p := 0
	fmt.Printf("Enter your pin number: ")
	fmt.Scan(&p)
	// check Pin
	if p != 11 {
		// Bad pin
		fmt.Println("Sorry, your pin has been rejected.")
		return false
	}
	fmt.Printf("Welcome %v\n", c.Name)
	a.cardName = c.Name
	a.cardBank = c.Bank
	a.cardAccountNumber = c.AccountNumber
	return true
}

// Withdraw will withdraw from ATMMachine
func (a *Machine) Withdraw(w int) {
	fmt.Printf("Going to withdraw %v for you %v\n", w, a.cardName)
	a.cardBalance = a.cardBalance - w
	fmt.Printf("Your new balance is %v\n", a.cardBalance)
}
