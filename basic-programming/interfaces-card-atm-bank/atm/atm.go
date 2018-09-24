package atm

import (
	"fmt"

	"github.com/JeffDeCola/my-go-examples/basic-programming/interfaces-card-atm-bank/bank"
)

type atm interface {
	inserted()
	getBalance(bank.Bank)
	deposited(bank.Bank, int) int
	withdrawn(bank.Bank, int) int
	ejected()
}

// MainStATM atn
type MainStATM struct {
	Name     string
	Location string
}

// DrugStoreATM atm
type DrugStoreATM struct {
	Location string
}

// BarATM atm
type BarATM struct {
	ATMName string
}

func (c BoACard) inserted() {
	fmt.Printf("   Screen - Hi %v, you inserted your %v card\n", c.Name, c.Bank)
}

func (c BoACard) getBalance(b bank.Bank) {
	balance := bank.TheBalance(b)
	fmt.Printf("   Screen - Your balance is %v\n", balance)
}

func (c BoACard) deposited(b bank.Bank, d int) int {
	balance := bank.TheDeposit(b, d)
	fmt.Printf("   Screen - Deposited %v. Your balance is %v\n", d, balance)
	return d
}

func (c BoACard) withdrawn(b bank.Bank, w int) int {
	balance := bank.TheWithdraw(b, w)
	fmt.Printf("   Screen - Dispensing %v. Your balance is %v\n", w, balance)
	return w
}

func (c BoACard) ejected() {
	fmt.Printf("   Screen - Goodbye %v, you ejected your %v card\n", c.Name, c.Bank)
}

func (c ChaseCard) inserted() {
	fmt.Printf("   Screen - Hi %v, you inserted your %v card\n", c.NameOnCard, c.Bank)
}

func (c ChaseCard) getBalance(b bank.Bank) {
	balance := bank.TheBalance(b)
	fmt.Printf("   Screen - Your balance is %v\n", balance)
}

func (c ChaseCard) deposited(b bank.Bank, d int) int {
	balance := bank.TheDeposit(b, d)
	fmt.Printf("   Screen - Deposited %v. Your balance is %v\n", d, balance)
	return d
}

func (c ChaseCard) withdrawn(b bank.Bank, w int) int {
	balance := bank.TheWithdraw(b, w)
	fmt.Printf("   Screen - Dispensing %v. Your balance is %v\n", w, balance)
	return w
}

func (c ChaseCard) ejected() {
	fmt.Printf("   Screen - Goodbye %v, you ejected your %v card\n", c.NameOnCard, c.Bank)
}

func (c WellsFargo) inserted() {
	fmt.Printf("   Screen - Hi, you inserted your %v card\n", c.Bank)
}

func (c WellsFargo) getBalance(b bank.Bank) {
	balance := bank.TheBalance(b)
	fmt.Printf("   Screen - Your balance is %v\n", balance)
}

func (c WellsFargo) deposited(b bank.Bank, d int) int {
	balance := bank.TheDeposit(b, d)
	fmt.Printf("   Screen - Deposited %v. Your balance is %v\n", d, balance)
	return d
}

func (c WellsFargo) withdrawn(b bank.Bank, w int) int {
	balance := bank.TheWithdraw(b, w)
	fmt.Printf("   Screen - Dispensing %v. Your balance is %v\n", w, balance)
	return w
}

func (c WellsFargo) ejected() {
	fmt.Printf("   Screen - Goodbye, you ejected your %v card\n", c.Bank)
}

// The interface accepts ANYTHING as long as that
// ANYTHING has a method attached to this interface

// InsertCard into atm
func InsertCard(a atm) {
	a.inserted()
}

// ShowBalance from bank
func ShowBalance(a atm, b bank.Bank) {
	a.getBalance(b)
}

// Deposit into bank
func Deposit(a atm, b bank.Bank, d int) int {
	return a.deposited(b, d)
}

// Withdraw from bank
func Withdraw(a atm, b bank.Bank, w int) int {
	return a.withdrawn(b, w)
}

// EjectCard from atm
func EjectCard(a atm) {
	a.ejected()
}
