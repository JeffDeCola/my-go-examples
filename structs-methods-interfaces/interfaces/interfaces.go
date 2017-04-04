// my-go-examples (methods-and-interfaces) interfaces.go

package main

import (
	"fmt"
	"math"
)

// Shape is the shape of all the objects
type Shape interface {
	area() float64
}

// Rectangle defines a rectangle
type Rectangle struct {
	x1 float64
	x2 float64
	y1 float64
	y2 float64
}

// Circle defines a circle
type Circle struct {
	x float64
	y float64
	r float64
}

// distance between two points
func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

// theArea will call the appropriate function based on receiver type.
func theArea(s Shape) float64 {
	return s.area()
}

// Compute the area of a rectagle
// Value receiver type rectangle
// Can not change the original struct in this function
func (r Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

// Compute the area of a circle
func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func main() {

	// rect1
	rect1 := Rectangle{x1: 1, x2: 11, y1: 1, y2: 11}

	// circ1 (no field names)
	circ1 := Circle{x: 0, y: 0, r: 5}

	fmt.Println("Area of rect1 is: ", theArea(&rect1))
	fmt.Println("Area of circ1 is: ", theArea(&circ1))

}
