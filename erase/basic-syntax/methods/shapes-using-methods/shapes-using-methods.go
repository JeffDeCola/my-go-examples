// my-go-examples shapes-using-methods.go

package main

import (
	"fmt"
	"math"
)

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
func (c Circle) circleArea() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// Circle circumference (will call it perimeter for this example)
func (c Circle) circlePerimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Rectangle area
func (r Rectangle) rectangleArea() float64 {
	return r.width * r.height
}

// Rectangle perimeter
func (r Rectangle) rectanglePerimeter() float64 {
	return (r.width + r.height) * 2
}

// Triangle area
func (t Triangle) triangleArea() float64 {
	// Heron's Formula to get area from 3 sides
	s := ((t.a + t.b + t.c) / 2)
	return math.Sqrt(s * (s - t.a) * (s - t.b) * (s - t.c))
}

// Triangle perimeter
func (t Triangle) trianglePerimeter() float64 {
	return t.a + t.b + t.c
}

func main() {

	// Define some shapes
	c1 := Circle{5}
	r1 := Rectangle{5, 3}
	t1 := Triangle{3, 4, 5}

	// Get the shape properties
	c1Area := c1.circleArea()
	c1Perimeter := c1.circlePerimeter()
	r1Area := r1.rectangleArea()
	r1Perimeter := r1.rectanglePerimeter()
	t1Area := t1.triangleArea()
    t1Perimeter := t1.trianglePerimeter()

	fmt.Printf("\nThe area of a circle with radius %v is                          %.3f\n", c1.radius, c1Area)
	fmt.Printf("The perimeter of a circle with radius %v is                     %.3f\n", c1.radius, c1Perimeter)
	fmt.Printf("The area of a rectangle with width %v and heigh %v is            %.3f\n", r1.width, r1.height, r1Area)
	fmt.Printf("The perimeter of a rectangle with with width %v and heigh %v is  %.3f\n", r1.width, r1.height, r1Perimeter)
	fmt.Printf("The area of a triangle with sides %v %v %v is                     %.3f\n", t1.a, t1.b, t1.c, t1Area)
	fmt.Printf("The perimeter of a triangle with sides %v %v %v is                %.3f\n", t1.a, t1.b, t1.c, t1Perimeter)
}
