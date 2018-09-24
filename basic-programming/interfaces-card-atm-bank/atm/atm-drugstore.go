package atm

import (
	"fmt"
)

type drugStoreATM struct {
	Name     string
	Location string
}

// InsertCard card inserted
func (a *drugStoreATM) InsertCard() {
	fmt.Println("hi")
}

// ShowBalance shows balance
func (a *drugStoreATM) ShowBalance() int {
	fmt.Println("hi")
	return 100
}

// DepositCash deposits cash
func (a *drugStoreATM) DepositCash(w int) int {
	fmt.Println("hi")
	return 100
}

// WithdrawCash withdraws cash
func (a *drugStoreATM) WithdrawCash(w int) int {
	fmt.Println("hi")
	return 100
}

// EjectCard ejects card
func (a *drugStoreATM) EjectCard() {
	fmt.Println("hi")
}

// NewdrugStoreATM creates an ATM instance
func NewdrugStoreATM() ATM {
	return &drugStoreATM{}
}
