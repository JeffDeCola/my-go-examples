package laboratory

import "fmt"

// Creatures describe Behaivior of creatures
type Creatures interface {
	Kind() string
	Fly() bool
	Sound() string
}

// Greet  - I only know about Creatures
func Greet(c Creatures) string {
	return c.Sound()
}

// FlyAway - I only know about Creatures
func FlyAway(canfly Creatures) string {
	if canfly.Fly() == true {
		return fmt.Sprintf("%s's can Fly - flap flap flap", canfly.Kind())
	}
	return fmt.Sprintf("%s's can't Fly", canfly.Kind())
}
