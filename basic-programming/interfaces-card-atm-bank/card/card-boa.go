package card

// BoACard is your BoA Card
type BoACard struct {
	Name          string
	Bank          string
	AccountNumber int
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
