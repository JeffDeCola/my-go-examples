package card

// ChaseCard is your Chase Card
type ChaseCard struct {
	NameOnCard string
	Bank       string
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
