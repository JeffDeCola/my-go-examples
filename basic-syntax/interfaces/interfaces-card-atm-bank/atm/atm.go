package atm

import (
	"fmt"

	"github.com/JeffDeCola/my-go-examples/basic-syntax/interfaces/interfaces-card-atm-bank/bank"
)

var banks map[string]bank.Networker

func init() {

	banks = map[string]bank.Networker{}

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
func LoadBank(name string, b bank.Networker) {

	banks[name] = b

}

// ShowBanksMap will show the banks Map in ATM
func ShowBanksMap(name string) {
	fmt.Printf("%+v\n", banks[name])
}
