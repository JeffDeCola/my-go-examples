package creatures

// Werewolf implements creature
type Werewolf struct {
	TimeofDay string
}

// Kind - Changes into a werewolf at night
func (w Werewolf) Kind() string {
	if w.TimeofDay == "day" {
		return "human"
	}
	return "werewolf"
}

// Fly - Nope
func (w Werewolf) Fly() bool {
	return false
}

// Sound depends on TimeofDay
func (w Werewolf) Sound() string {
	if w.TimeofDay == "day" {
		return "hello"
	}
	return "howl"
}

// Vampire implements creature
type Vampire struct {
	Age int
}

// Kind - after 100 becomes a vampire
func (z Vampire) Kind() string {
	if z.Age > 100 {
		return "Vampire"
	}
	return "human"
}

// Fly - Yes
func (z Vampire) Fly() bool {
	return true
}

// Sound - he talks funny
func (z Vampire) Sound() string {
	return "I want to drink your blood"
}
