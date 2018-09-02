// my-go-examples (methods-and-interfaces) structs.go

package main

import ("fmt"; "math")

type Rectangle struct {
	x1 float64
	x2 float64
	y1 float64
	y2 float64
}

type Circle struct {
	x float64
	y float64
	r float64
}

// Find the distance between two points
func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

// Compute the area of a rectagle
func rectangleArea(r Rectangle) float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

// Compute the area of a circle
func circleArea(c Circle) float64 {
	return math.Pi * c.r*c.r
}

func main() {
 	
	// rect1
	rect1 := Rectangle{x1: 1, x2: 11, y1: 1, y2: 11}
	
	// circ1 (no field names)
	circ1 := Circle{x: 0, y: 0, r: 5}

  	fmt.Println("Area of this rect1 is: ", rectangleArea(rect1))
  	fmt.Println("Area of this circ1 is: ", circleArea(circ1))

}