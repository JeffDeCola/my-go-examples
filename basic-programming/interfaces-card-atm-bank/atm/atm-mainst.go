package atm

import (
	"fmt"
)

// MainStATM atm
type mainStATM struct {
	Name     string
	Location string
}

// InsertCard card inserted
func (a *mainStATM) InsertCard() {
	fmt.Println("hi")
}

// ShowBalance shows balance
func (a *mainStATM) ShowBalance() int {
	fmt.Println("hi")
	return 100
}

// DepositCash deposits cash
func (a *mainStATM) DepositCash(w int) int {
	fmt.Println("hi")
	return 100
}

// WithdrawCash withdraws cash
func (a *mainStATM) WithdrawCash(w int) int {
	fmt.Println("hi")
	return 100
}

// EjectCard ejects card
func (a *mainStATM) EjectCard() {
	fmt.Println("hi")
}

// NewMainStATM creates an ATM instance
func NewMainStATM() ATM {
	return &mainStATM{}
}
