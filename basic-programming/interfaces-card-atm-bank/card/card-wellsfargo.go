package card

// WellsFargo is your WellsFargo Card
type WellsFargo struct {
	Bank string
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

