package atm

import (
	"github.com/JeffDeCola/my-go-examples/basic-syntax/interfaces/interfaces-card-atm-bank/bank"
	"github.com/JeffDeCola/my-go-examples/basic-syntax/interfaces/interfaces-card-atm-bank/card"
)

var banks map[string]bank.Banker

func init() {

	banks = map[string]bank.Banker{}

}

// ATMer interface
type ATMer interface {
	showBalance(card.Carder) int
	depositCash(card.Carder, int) int
	withdrawCash(card.Carder, int) int
}

//LoadBank loads a bank interface to a package level map
func LoadBank(b bank.Banker) {
	BranchName := bank.GetBankInfo(b)
	banks[BranchName] = b
}
