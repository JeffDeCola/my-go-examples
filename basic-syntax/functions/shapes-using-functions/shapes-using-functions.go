// my-go-examples shapes-using-functions.go

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
func circleArea(c Circle) float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// Circle circumference (will call it perimeter for this example)
func circlePerimeter(c Circle) float64 {
	return 2 * math.Pi * c.radius
}

// Rectangle area
func rectangleArea(r Rectangle) float64 {
	return r.width * r.height
}

// Rectangle perimeter
func rectanglePerimeter(r Rectangle) float64 {
	return (r.width + r.height) * 2
}

// Triangle area
func triangleArea(t Triangle) float64 {
	// Heron's Formula to get area from 3 sides
	s := ((t.a + t.b + t.c) / 2)
	return math.Sqrt(s * (s - t.a) * (s - t.b) * (s - t.c))
}

// Triangle perimeter
func trianglePerimeter(t Triangle) float64 {
	return t.a + t.b + t.c
}

func main() {

	// Define some shapes
	c1 := Circle{5}
	r1 := Rectangle{5, 3}
	t1 := Triangle{3, 4, 5}

	// Get the shape properties
	c1Area := circleArea(c1)
	c1Perimeter := circlePerimeter(c1)
	r1Area := rectangleArea(r1)
	r1Perimeter := rectanglePerimeter(r1)
	t1Area := triangleArea(t1)
	t1Perimeter := trianglePerimeter(t1)

	fmt.Printf("The area of a circle with radius %v is                          %.3f\n", c1.radius, c1Area)
	fmt.Printf("The perimeter of a circle with radius %v is                     %.3f\n", c1.radius, c1Perimeter)
	fmt.Printf("The area of a rectangle with width %v and heigh %v is            %.3f\n", r1.width, r1.height, r1Area)
	fmt.Printf("The perimeter of a rectangle with with width %v and heigh %v is  %.3f\n", r1.width, r1.height, r1Perimeter)
	fmt.Printf("The area of a triangle with sides %v %v %v is                     %.3f\n", t1.a, t1.b, t1.c, t1Area)
	fmt.Printf("The perimeter of a triangle with sides %v %v %v is                %.3f\n", t1.a, t1.b, t1.c, t1Perimeter)
}
