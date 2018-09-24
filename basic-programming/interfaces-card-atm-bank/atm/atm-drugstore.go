package atm

// DrugStoreATM atm
type DrugStoreATM struct {
	Location string
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
