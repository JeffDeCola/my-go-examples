package creatures

// Warewolf impliments creature
type Warewolf struct {
	TimeofDay string
}

// Kind is a method - Receiver Type is Warewolf
func (w Warewolf) Kind() string {
	if w.TimeofDay == "day" {
		return "human"
	}
	return "warewolf"
}

// Fly is a method  - Receiver Type is Warewolf
// Exists to satify creature interface
func (w Warewolf) Fly() bool {
	return false
}

// Sound is a method - Receiver Type is Warewolf
func (w Warewolf) Sound() string {
	if w.TimeofDay == "day" {
		return "hello"
	}
	return "howl"
}

// Vampire impliments creature
type Vampire struct {
	Age int
}

// Kind is a method - Receiver Type is Vampire
func (z Vampire) Kind() string {
	if z.Age > 100 {
		return "Vampire"
		//return fmt.Sprintf("Vampire that is %d years old\n", s.Age)
	}
	return "human"
}

// Fly is a method - Receiver Type is Vampire
func (z Vampire) Fly() bool {
	return true
}

// Sound is a method - Receiver Type is Vampire
func (z Vampire) Sound() string {
	return "I want to drink your blood"
}
