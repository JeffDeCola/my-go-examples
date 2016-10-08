package laboratory

import "fmt"

// Creatures describe creatures
// An interface type is defined as a set of Method signatures.
type Creatures interface {
	Kind() string
	Gender() string
	Fly() bool
	Sound() string
}

// CreatureSound - I only know about Creatures
func CreatureSound(scarysound Creatures) string {
	return scarysound.Sound()
}

// CreatureFly - I only know about Creatures
func CreatureFly(canfly Creatures) string {
	if canfly.Fly() == true {
		return fmt.Sprintf("%s's can Fly", canfly.Kind())
	}
	return fmt.Sprintf("%s's can't Fly", canfly.Kind())
}
