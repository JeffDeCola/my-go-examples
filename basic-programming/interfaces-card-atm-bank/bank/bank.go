package bank

// Bank interface
type Bank interface {
	balance() int
	deposited(int) int
	withdrew(int) int
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
