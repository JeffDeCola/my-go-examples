package atm

type atm interface {
	InsertCard()
	ShowBalance() int
	DepositCash(int) int
	WithdrawnCash(int) int
	EjectCard()
}

// The interface accepts ANYTHING as long as that
// ANYTHING has a method attached to this interface
