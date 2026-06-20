// interfaces-pointers-receivers
//
// Using an interface to resize a rectangle and circle using pointer receivers.
//

package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type scaler interface {
	scale(factor float64)
}

type rectangle struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

func getArea(s shape) float64 {
	area := s.area()
	return area
}

func (r rectangle) area() float64 {
	area := r.width * r.height
	return area
}

func (c circle) area() float64 {
	area := math.Pi * c.radius * c.radius
	return area
}

func (r *rectangle) scale(factor float64) {
	r.width *= factor
	r.height *= factor
}

func (c *circle) scale(factor float64) {
	c.radius *= factor
}

func main() {
	rec := rectangle{
		width:  10,
		height: 5,
	}

	circ := circle{
		radius: 5,
	}

	// Polymorphism: one func, many types
	recArea := getArea(rec)
	fmt.Printf("The area of the rectangle (%.2f x %.2f) is %.2f\n", rec.width, rec.height, recArea)

	circArea := getArea(circ)
	fmt.Printf("The area of the circle (radius %.2f) is %.2f\n", circ.radius, circArea)

	// Scale via the interface
	var sr scaler = &rec
	sr.scale(3)
	recArea = getArea(&rec)
	fmt.Printf("The area of the rectangle (%.2f x %.2f) is %.2f\n", rec.width, rec.height, recArea)

	var sc scaler = &circ
	sc.scale(4)
	circArea = getArea(&circ)
	fmt.Printf("The area of the circle (radius %.2f) is %.2f\n", circ.radius, circArea)

}
