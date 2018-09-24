package card

import (
	"fmt"
)

type boACard struct {
	Bank string
}

// GetBankInfo to get bank info from card
func (c *boACard) GetBankInfo() string {
	fmt.Printf("   Screen - Hi, you inserted your %v card\n", c.Bank)
	return "hi"
}

// NewBoACard creates an ATM instance
func NewBoACard() Card {
	return &boACard{}
}
