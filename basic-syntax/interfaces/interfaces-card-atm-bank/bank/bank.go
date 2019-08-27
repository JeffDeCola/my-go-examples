package bank

import "github.com/JeffDeCola/my-go-examples/basic-syntax/interfaces/interfaces-card-atm-bank/card"

// Banker interface
type Banker interface {
	addAccount(string, int)
	issueCard(string) card.Carder
	getBankInfo() string
	balance(string) int
	deposit(string, int) int
	withdraw(string, int) int
}
