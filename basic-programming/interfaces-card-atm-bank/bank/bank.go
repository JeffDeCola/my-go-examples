package BANK

// Bank interface
type Bank interface {
	Balance() int
	Deposit(int) int
	Withdrew(int) int
}
