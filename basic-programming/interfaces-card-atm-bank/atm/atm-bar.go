package atm

import (
	"fmt"
)

// barATM atm
type barATM struct {
	Name     string
	Location string
}

// InsertCard card inserted
func (a *barATM) InsertCard() {
	fmt.Println("hi")
}

// ShowBalance shows balance
func (a *barATM) ShowBalance() int {
	fmt.Println("hi")
	return 100
}

// DepositCash deposits cash
func (a *barATM) DepositCash(w int) int {
	fmt.Println("hi")
	return 100
}

// WithdrawCash withdraws cash
func (a *barATM) WithdrawCash(w int) int {
	fmt.Println("hi")
	return 100
}

// EjectCard ejects card
func (a *barATM) EjectCard() {
	fmt.Println("hi")
}

// NewBarATM creates an ATM instance
func NewBarATM() ATM {
	return &barATM{}
}
