package card

// BoaCard Data
type boaCard struct {
	Name       string
	BranchName string
}

//***************************************************************************************************

// GetBankInfo to get bank info from card
func (c *boaCard) getBankInfo() (string, string) {
	return c.Name, c.BranchName
}

//***************************************************************************************************

// NewBoaCard creates a Card
func NewBoaCard(name string, BranchName string) Carder {
	// Notice how we have to init the map because we can use it
	return &boaCard{Name: name, BranchName: BranchName}
}

// GetBankInfo gets bank info
func GetBankInfo(c Carder) (string, string) {
	return c.getBankInfo()
}
