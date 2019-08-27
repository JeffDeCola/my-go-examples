package bank

import (
	"fmt"

	"github.com/JeffDeCola/my-go-examples/basic-syntax/interfaces/interfaces-card-atm-bank/card"
)

// BoaBank is the data
type BoaBank struct {
	BranchName string
	Accounts   map[string]int
}

//***************************************************************************************************

// AddAccount makes new account
func (b *BoaBank) addAccount(name string, deposit int) {
	fmt.Printf("    For bank %v, You are creating an account for %v with %d\n", b, name, deposit)
	b.Accounts[name] = deposit
}

// issueCard issues an atm card
func (b *BoaBank) issueCard(name string) card.Carder {
	fmt.Printf("    For bank %v, You are creating an card for %v\n", b, name)
	return card.NewBoaCard(name, b.BranchName)
}

// GetBankInfo to get bank info from card
func (b *BoaBank) getBankInfo() string {
	return b.BranchName
}

// Balance returns the balance
func (b *BoaBank) balance(name string) int {
	//fmt.Printf("    Hi %v. You have %d in your account.\n", name, b.Accounts[name])
	return b.Accounts[name]
}

// Deposit deposits int the Account
func (b *BoaBank) deposit(name string, d int) int {
	b.Accounts[name] = b.Accounts[name] + d
	//fmt.Printf("    Hi %v. You deposited %d in your account and now have a total of %v.\n", name, d, b.Accounts[name])
	return b.Accounts[name]
}

// Withdraw withdraws from the account
func (b *BoaBank) withdraw(name string, w int) int {
	b.Accounts[name] = b.Accounts[name] - w
	//fmt.Printf("    Hi %v. You withdrew %d from your account and now have a total of %v.\n", name, w, b.Accounts[name])
	return b.Accounts[name]
}

//***************************************************************************************************

// NewBoaBank creates a Bank
func NewBoaBank(name string) Banker {
	// Notice how we have to init the map because we can use it
	return &BoaBank{BranchName: name, Accounts: make(map[string]int)}
}

// NewBoaAccount to crate an account at a bank
func NewBoaAccount(b Banker, name string, deposit int) {
	b.addAccount(name, deposit)
}

// NewBoaCard to crate a card that can be used at an atm
func NewBoaCard(b Banker, name string) card.Carder {
	return b.issueCard(name)
}

// GetBankInfo gets bank info
func GetBankInfo(b Banker) string {
	return b.getBankInfo()
}

// GetBalance gets balance
func GetBalance(b Banker, name string) int {
	return b.balance(name)
}

// MakeDeposit makes a deposit
func MakeDeposit(b Banker, name string, deposit int) int {
	return b.deposit(name, deposit)
}

// MakeWithdraw makes a withdrawal
func MakeWithdraw(b Banker, name string, withdraw int) int {
	return b.withdraw(name, withdraw)
}
