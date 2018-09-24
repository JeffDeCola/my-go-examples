package atm

import (
	"fmt"
)

// MainStATM atm
type MainStATM struct {
	Name     string
	Location string
}

// InsertCard card inserted
func (a *MainStATM) InsertCard() {
	fmt.Println("hi")
}

// ShowBalance shows balance
func (a *MainStATM) ShowBalance() int {
	fmt.Println("hi")
	return 100
}

// DepositCash deposits cash
func (a *MainStATM) DepositCash(w int) int {
	fmt.Println("hi")
	return 100
}

// WithdrawCash withdraws cash
func (a *MainStATM) WithdrawCash(w int) int {
	fmt.Println("hi")
	return 100
}

// EjectCard ejects card
func (a *MainStATM) EjectCard() {
	fmt.Println("hi")
}

// NewMainStATM creates an ATM instance
func NewMainStATM() ATM {
	return &MainStATM{}
}
