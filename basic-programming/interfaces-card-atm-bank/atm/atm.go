package atm

// ATM interface
type ATM interface {
	InsertCard()
	ShowBalance() int
	DepositCash(int) int
	WithdrawCash(int) int
	EjectCard()
}

// The interface accepts ANYTHING as long as that
// ANYTHING has a method attached to this interface
