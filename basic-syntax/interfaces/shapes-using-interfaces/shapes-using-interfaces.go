// my-go-examples shapes-using-interfaces.go

package main

import (
	"fmt"
	"math"
)

type areaer interface {
	area() float64
}

type perimeterer interface {
	perimeter() float64
}

// Circle description
type Circle struct {
	radius float64
}

// Rectangle description
type Rectangle struct {
	width  float64
	height float64
}

// Triangle description
type Triangle struct {
	a float64
	b float64
	c float64
}

// Circle area
func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// Circle circumference (will call it perimeter for this example)
func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Rectangle area
func (r Rectangle) area() float64 {
	return r.width * r.height
}

// Rectangle perimeter
func (r Rectangle) perimeter() float64 {
	return (r.width + r.height) * 2
}

// Triangle area
func (t Triangle) area() float64 {
	// Heron's Formula to get area from 3 sides
	s := ((t.a + t.b + t.c) / 2)
	return math.Sqrt(s * (s - t.a) * (s - t.b) * (s - t.c))
}

// Triangle perimeter
func (t Triangle) perimeter() float64 {
	return t.a + t.b + t.c
}

func area(a areaer) float64 {
	return a.area()
}

func perimeter(p perimeterer) float64 {
	return p.perimeter()
}

func main() {

	// Define some shapes
	c1 := Circle{5}
	r1 := Rectangle{5, 3}
	t1 := Triangle{3, 4, 5}

	// We could fod it this way
	// Get the shape properties
	c1Area := c1.area()
	c1Perimeter := c1.perimeter()
	r1Area := r1.area()
	r1Perimeter := r1.perimeter()
	t1Area := t1.area()
	t1Perimeter := t1.perimeter()

	fmt.Printf("\nThe area of a circle with radius %v is                          %.3f\n", c1.radius, c1Area)
	fmt.Printf("The perimeter of a circle with radius %v is                     %.3f\n", c1.radius, c1Perimeter)
	fmt.Printf("The area of a rectangle with width %v and heigh %v is            %.3f\n", r1.width, r1.height, r1Area)
	fmt.Printf("The perimeter of a rectangle with with width %v and heigh %v is  %.3f\n", r1.width, r1.height, r1Perimeter)
	fmt.Printf("The area of a triangle with sides %v %v %v is                     %.3f\n", t1.a, t1.b, t1.c, t1Area)
	fmt.Printf("The perimeter of a triangle with sides %v %v %v is                %.3f\n", t1.a, t1.b, t1.c, t1Perimeter)

	// But this way is easier and makes more sense
	// Get the shape properties
	c1Area = area(c1)
	c1Perimeter = perimeter(c1)
	r1Area = area(r1)
	r1Perimeter = perimeter(r1)
	t1Area = area(t1)
	t1Perimeter = perimeter(t1)

	fmt.Printf("\nThe area of a circle with radius %v is                          %.3f\n", c1.radius, c1Area)
	fmt.Printf("The perimeter of a circle with radius %v is                     %.3f\n", c1.radius, c1Perimeter)
	fmt.Printf("The area of a rectangle with width %v and heigh %v is            %.3f\n", r1.width, r1.height, r1Area)
	fmt.Printf("The perimeter of a rectangle with with width %v and heigh %v is  %.3f\n", r1.width, r1.height, r1Perimeter)
	fmt.Printf("The area of a triangle with sides %v %v %v is                     %.3f\n", t1.a, t1.b, t1.c, t1Area)
	fmt.Printf("The perimeter of a triangle with sides %v %v %v is                %.3f\n", t1.a, t1.b, t1.c, t1Perimeter)
}
