package card

import (
	"fmt"
)

type wellsFargoCard struct {
	Bank        string
	ThisisWells string
}

// GetBankInfo to get bank info from card
func (c *wellsFargoCard) GetBankInfo() string {
	fmt.Printf("   Screen - Hi, you inserted your %v card\n", c.Bank)
	return "hi"
}

// NewWellsFargoCard creates an ATM instance
func NewWellsFargoCard() Card {
	return &wellsFargoCard{}
}
