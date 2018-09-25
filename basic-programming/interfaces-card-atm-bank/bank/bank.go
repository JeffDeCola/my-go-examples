package bank

// Bank interface
type Bank interface {
	BankNetworker
	BankCustomerServicer
}

// BankNetworker interface
type BankNetworker interface {
	Balance() int
	Deposit(int) int
	Withdraw(int) int
}

// BankCustomerServicer interface
type BankCustomerServicer interface {
	IssueCard()
}
