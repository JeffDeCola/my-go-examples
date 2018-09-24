package atm

// BarATM atm
type BarATM struct {
	ATMName string
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

func NewBarATM() BarATM {
	return &BarATM{}
}
