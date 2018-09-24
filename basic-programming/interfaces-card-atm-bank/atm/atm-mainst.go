package atm

// MainStATM atm
type MainStATM struct {
	Name     string
	Location string
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
