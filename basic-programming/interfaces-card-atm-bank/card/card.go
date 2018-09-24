package card

// BoACard is your BoA Card
type BoACard struct {
	Name          string
	Bank          string
	AccountNumber int
}

//ChaseCard is your Chase Card
type ChaseCard struct {
	NameOnCard string
}

//WellsFargo is your WellsFargo Card
type WellsFargo struct {
	Bank string
}
