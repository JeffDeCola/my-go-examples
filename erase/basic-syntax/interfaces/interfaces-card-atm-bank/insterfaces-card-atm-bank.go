package main

import (
	"fmt"

	"github.com/JeffDeCola/my-go-examples/basic-syntax/interfaces/interfaces-card-atm-bank/atm"
	"github.com/JeffDeCola/my-go-examples/basic-syntax/interfaces/interfaces-card-atm-bank/bank"
)

func main() {

	fmt.Println(" ")

	fmt.Println("Make some banks")
	b1 := bank.NewBoaBank("elmStBranch")
	b2 := bank.NewBoaBank("downtownBranch")

	fmt.Println("Add Some accounts with initial deposits to banks")
	bank.NewBoaAccount(b1, "Jeff", 100)
	bank.NewBoaAccount(b1, "Jill", 1000)
	fmt.Printf("Jeff, your balance is %v\n", bank.GetBalance(b1, "Jeff"))
	fmt.Printf("Jill, your balance is %v\n", bank.GetBalance(b1, "Jill"))

	fmt.Println("3. Issue an atm card from that bank")
	c1 := bank.NewBoaCard(b1, "Jeff")
	c2 := bank.NewBoaCard(b1, "Jill")

	fmt.Println("4. Load Banks into ATM")
	atm.LoadBank(b1)
	atm.LoadBank(b2)

	fmt.Println("5. Create an ATM machine")
	a1 := atm.NewATM("123Main")
	a2 := atm.NewATM("sunset")
	//fmt.Println(a1)
	//fmt.Println(a2)

	fmt.Println(" ")
	fmt.Println("----------------")
	fmt.Println(" ")

	fmt.Println("Jeff Make a Transaction")
	atm.ShowBalance(a1, c1)
	atm.Deposit(a1, c1, 740)
	atm.Withdraw(a1, c1, 50)

	fmt.Println("Jill Make a Transaction")
	atm.ShowBalance(a2, c2)
	atm.Deposit(a2, c2, 1740)

	fmt.Println(" ")

}
