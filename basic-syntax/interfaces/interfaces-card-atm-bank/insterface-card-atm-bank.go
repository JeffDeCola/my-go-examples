package main

import (
	"fmt"

	"github.com/JeffDeCola/my-go-examples/basic-syntax/interfaces/interfaces-card-atm-bank/bank"
)

func main() {

	// 1. Make a Bank Interface
	b1 := bank.NewBoaBank("elmStBranch")
	b2 := bank.NewBoaBank("downtownBranch")

	// 2. Add Some accounts to that bank with initial deposit
	//bank.NewBoaAccount(b1, "Jeff", 100)
	//bank.NewBoaAccount(b1, "Jill", 100)

	// Now I have the bank and I can get info like
	fmt.Println(b1)
	fmt.Println(b2)

	// Send Networker interface (the Bank Struct) to atm to put in map
	//atm.LoadBank("Boa", b)
	//atm.LoadBank("Chase", c)
	//atm.LoadBank("WellsFargo", w)

	// Lets look at the map of Bank Structs
	//atm.ShowBanksMap("BoA")
	//atm.ShowBanksMap("Chase")
	//atm.ShowBanksMap("WellsFargo")

	// 2. Issue cards from Bank
	// Via the BankNetworker interface, get Card interface (the card struct)
	//boaCard := b.IssueCard()
	//chaseCard := c.IssueCard()
	//wellsFargoCard := w.IssueCard()

	// 3. Create an ATM

	// 4. Put card in ATM

	// 5. Make a Transaction

	// 6. Eject Card

}
