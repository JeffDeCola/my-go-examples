package atm

// ATM interface
type ATM interface {
	InsertCard()
	ShowBalance() int
	DepositCash(int) int
	WithdrawCash(int) int
	EjectCard()
}
