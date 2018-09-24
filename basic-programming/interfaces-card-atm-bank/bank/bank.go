package BANK

// Bank interface
type Bank interface {
	Balance() int
	Deposit(int) int
	Withdrew(int) int
}

// TheBalance from bank
func TheBalance(b Bank) int {
	return b.balance()
}

// TheDeposit from bank
func TheDeposit(b Bank, n int) int {
	return b.deposited(n)
}

// TheWithdraw from atm
func TheWithdraw(b Bank, n int) int {
	return b.withdrew(n)
}
