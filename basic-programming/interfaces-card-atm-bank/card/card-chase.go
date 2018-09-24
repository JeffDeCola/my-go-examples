package card

import (
	"fmt"
)

type chaseCard struct {
	Bank string
}

// GetBankInfo to get bank info from card
func (c *chaseCard) GetBankInfo() string {
	fmt.Printf("   Screen - Hi, you inserted your %v card\n", c.Bank)
	return "hi"
}

// NewChaseCard creates an ATM instance
func NewChaseCard() Card {
	return &chaseCard{}
}
