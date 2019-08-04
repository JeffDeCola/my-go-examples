package main

import (
	"github.com/JeffDeCola/my-go-examples/basic-programming/interfaces-card-atm-bank/atm"
	"github.com/JeffDeCola/my-go-examples/basic-programming/interfaces-card-atm-bank/bank"
)

func main() {

	// 1. Load Banks Structs into the ATM (So the ATM knows about Banks)
	// Create a Bank (return Bank Struct via interface BankNetworker)
	b := bank.NewBoABank()
	c := bank.NewChaseBank()
	w := bank.NewWellsFargoBank()
	// Send BankNetworker interface (the Bank Struct) to atm to put in map
	atm.LoadBank("BoA", b)
	atm.LoadBank("Chase", c)
	atm.LoadBank("WellsFargo", w)
	// Lets look at the map of Bank Structs
	atm.ShowBanksMap("BoA")
	atm.ShowBanksMap("Chase")
	atm.ShowBanksMap("WellsFargo")

	// 2. Issue cards from Bank
	// Via the BankNetworker interface, get Card interface (the card struct)
	boaCard := b.IssueCard()
	chaseCard := c.IssueCard()
	wellsFargoCard := w.IssueCard()

	// 3. Create an ATM

	// 4. Put card in ATM

	// 5. Make a Transaction

	// 6. Eject Card

}
