package bank

// Bank interface
type Bank interface {
	Networker
	CustomerServicer
}

// Networker interface
type Networker interface {
	Balance(string) int
	Deposit(int) int
	Withdraw(int) int
}

// CustomerServicer interface
type CustomerServicer interface {
	IssueCard()
	AddAccount()
}
