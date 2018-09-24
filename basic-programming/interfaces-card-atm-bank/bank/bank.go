package bank

// Bank interface
type Bank interface {
	balance() int
	deposited(int) int
	withdrew(int) int
}

// MainStATM atn
type MainStATM struct {
	Name     string
	Location string
}

// DrugStoreATM atm
type DrugStoreATM struct {
	Location string
}

// BarATM atm
type BarATM struct {
	ATMName string
}

func (a MainStATM) balance() int {
	return 100
}

func (a MainStATM) deposited(d int) int {
	return 100 + d
}

func (a MainStATM) withdrew(w int) int {
	return 100 - w
}

func (a DrugStoreATM) balance() int {
	return 100
}

func (a DrugStoreATM) deposited(d int) int {
	return 100 + d
}

func (a DrugStoreATM) withdrew(w int) int {
	return 100 - w
}

func (a BarATM) balance() int {
	return 100
}

func (a BarATM) deposited(d int) int {
	return 100 + d
}

func (a BarATM) withdrew(w int) int {
	return 100 - w
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
