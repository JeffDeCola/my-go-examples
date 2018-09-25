package atm

import (
	"github.com/JeffDeCola/my-go-examples/basic-programming/interfaces-card-atm-bank/bank"
)

var banks map[string]bank.BankNetworker

func init() {

	banks = map[string]bank.BankNetworker{}

}

// ATM interface
type ATM interface {
	InsertCard()
	ShowBalance() int
	DepositCash(int) int
	WithdrawCash(int) int
	EjectCard()
}

//LoadBank loads a bank interface to a package level map
func LoadBank(name string, b bank.BankNetworker) {

	banks[name] = b

}
