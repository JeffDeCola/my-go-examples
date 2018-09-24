package atm

import (
	"fmt"
)

// BarATM atm
type BarATM struct {
	Name     string
	Location string
}

// InsertCard card inserted
func (a *BarATM) InsertCard() {
	fmt.Println("hi")
}

// ShowBalance shows balance
func (a *BarATM) ShowBalance() int {
	fmt.Println("hi")
	return 100
}

// DepositCash deposits cash
func (a *BarATM) DepositCash(w int) int {
	fmt.Println("hi")
	return 100
}

// WithdrawCash withdraws cash
func (a *BarATM) WithdrawCash(w int) int {
	fmt.Println("hi")
	return 100
}

// EjectCard ejects card
func (a *BarATM) EjectCard() {
	fmt.Println("hi")
}

// NewBarATM creates an ATM instance
func NewBarATM() ATM {
	return &BarATM{}
}
