package atm

import (
	"fmt"

	"github.com/JeffDeCola/my-go-examples/basic-syntax/interfaces/interfaces-card-atm-bank/bank"
	"github.com/JeffDeCola/my-go-examples/basic-syntax/interfaces/interfaces-card-atm-bank/card"
)

// ATM data
type ATM struct {
	Name string
}

//***************************************************************************************************

// showBalance shows balance
func (a *ATM) showBalance(c card.Carder) int {
	Name, BranchName := card.GetBankInfo(c)
	balance := bank.GetBalance(banks[BranchName], Name)
	fmt.Printf("    Hi %v, you are at ATM %v and you have %v in your account at bank %v\n", Name, a.Name, balance, BranchName)
	return balance
}

// depositCash deposits cash
func (a *ATM) depositCash(c card.Carder, deposit int) int {
	Name, BranchName := card.GetBankInfo(c)
	total := bank.MakeDeposit(banks[BranchName], Name, deposit)
	fmt.Printf("    Hi %v, you are at ATM %v and you have %v after your deposit of %v at bank %v\n", Name, a.Name, total, deposit, BranchName)
	return total
}

// withdrawCash withdraws cash
func (a *ATM) withdrawCash(c card.Carder, withdraw int) int {
	Name, BranchName := card.GetBankInfo(c)
	total := bank.MakeWithdraw(banks[BranchName], Name, withdraw)
	fmt.Printf("    Hi %v, you are at ATM %v and you have %v after your withdrawal of %v at bank %v\n", Name, a.Name, total, withdraw, BranchName)
	return total
}

//***************************************************************************************************

// NewATM creates an ATM
func NewATM(name string) ATMer {
	return &ATM{Name: name}
}

// ShowBalance shows balance
func ShowBalance(a ATMer, c card.Carder) int {
	return a.showBalance(c)
}

// Deposit creates an ATM
func Deposit(a ATMer, c card.Carder, deposit int) int {
	return a.depositCash(c, deposit)
}

// Withdraw creates an ATM
func Withdraw(a ATMer, c card.Carder, withdraw int) int {
	return a.withdrawCash(c, withdraw)
}
