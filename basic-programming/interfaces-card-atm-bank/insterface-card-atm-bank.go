package main

import (
	"fmt"

	"github.com/JeffDeCola/my-go-examples/basic-programming/interfaces-card-atm-bank/card"

	"github.com/JeffDeCola/my-go-examples/basic-programming/interfaces-card-atm-bank/atm"
)

func chooseCard(c *card.ATMCard) bool {
	var q int

	c1 := card.ATMCard{Name: "Jeff", Bank: "BoA", AccountNumber: 2001}
	c2 := card.ATMCard{Name: "Jeffry", Bank: "Chase", AccountNumber: 232123212}

	fmt.Println("What card would you like to use")
	fmt.Println("1. BoA")
	fmt.Println("2. Chase")
	fmt.Scan(&q)
	switch {
	case q == 1:
		*c = c1
	case q == 2:
		*c = c2
	default:
		fmt.Println("I don;t know that card")
		return false
	}
	return true
}

func main() {

	// Your ATM card
	c := card.ATMCard{}

	// The ATM machine interface
	m := atm.Machine{}
	var a atm.ATM = &m

	if !(chooseCard(&c)) {
		// Bad Card
	}

    fmt.Printf("Insert your %v Card\n", c.Bank)
    // Card is
	if !(a.InsertCard(c)) {
		// Pin Wrong
		// Eject card
	}

	if !(a.AtMachinde(c)) {
		// Eject card
	}

    fmt.Printf("\nWhat would you like to do?")
			fmt.Println("1. Show Balance")
			fmt.Println("2. Deposit Cash")
			fmt.Println("3. Withdraw Cash")
			fmt.Println("4. Eject Card")
			fmt.Scan(&q)
			switch {
			case q == 1:
				fmt.Println("You Balance", a.ShowBalance())
			case q == 2:
				fmt.Println("You Balance", a.ShowBalance())
			case q == 3:
				fmt.Print("How much would you like to withdraw? ")
				fmt.Scan(&q)
				a.Withdraw(q)
			case q == 4:
				fmt.Println("Ejecting Card")
				break
			default:
				fmt.Println("Ejecting Card")
				break
			}
		
		fmt.Println("Thank you - Have a nice day")

	}

}
