package creatures

// Jeff is a warewolf
type Jeff struct {
	TimeofDay string
}

// Kind is a method Receiver Type Jeff
func (j Jeff) Kind() string {
	if j.TimeofDay == "day" {
		return "human"
	}
	return "warewolf"
}

// Fly is a method Receiver Type Jeff
func (j Jeff) Fly() bool {
	return false
}

// Sound is a method Receiver Type Jeff
func (j Jeff) Sound() string {
	if j.TimeofDay == "day" {
		return "hello"
	}
	return "howl"
}

// Gender is a method Receiver Type Jeff
func (j Jeff) Gender() string {
	return "male"
}

// Steven is a zombie
type Steven struct {
	Age    int
	Height int
}

// Kind is a Method with Receiver Type Steven
func (s Steven) Kind() string {
	if s.Age > 100 {
		return "zombie"
		//return fmt.Sprintf("Zombie that is %d years old\n", s.Age)
	}
	return "human"
}

// Fly is a Method with Receiver Type Steven
func (s Steven) Fly() bool {
	return false
}

// Sound is a Method with Receiver Type Steven
func (s Steven) Sound() string {
	return "growl"
}

// Gender is a method Receiver Type Doris
func (s Steven) Gender() string {
	return "male"
}

// Doris is a witch
type Doris string

// Kind is a Method with Receiver Type Doris
func (d Doris) Kind() string {
	return "witch"
}

// Fly is a Method with Receiver Type Doris
func (d Doris) Fly() bool {
	return true
}

// Sound is a Method with Receiver Type Steven
func (d Doris) Sound() string {
	return "Weeeeeeee"
}

// Gender is a Method with Receiver Type Doris
func (d Doris) Gender() string {
	return "female"
}
