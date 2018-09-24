package atm

type atm interface {
	insertCard()
	showBalance() int
	depositCash(int) int
	withdrawnCash(int) int
	ejectCard()
}

// The interface accepts ANYTHING as long as that
// ANYTHING has a method attached to this interface
