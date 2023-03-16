package main

import (
	"fmt"
	"math"
)

// INTERFACE TYPES - TIES INTERFACE WITH METHODS
type geometry interface {
	area(*float64)
	size(float64)
}

type rectangle struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

// MATH USING METHODS
func (r rectangle) area(a *float64) {
	*a = r.width * r.height
}

func (c circle) area(a *float64) {
	*a = math.Pi * math.Pow(c.radius, 2)
}

func (r *rectangle) size(f float64) {
	r.width = r.width * f
	r.height = r.height * f
}

func (c *circle) size(f float64) {
	c.radius = c.radius * f
}

func main() {

	// DEFINE
	rec := rectangle{2.4, 34.4}
	circ := circle{2.3}

	// CALCULATE AREA USING AN INTERFACE TYPE
	var gRec geometry
	var gCirc geometry
	var recArea float64
	var circArea float64

	// Note using the address
	gRec = &rec
	gCirc = &circ

	// CALCULATE AREA USING A STRUCT TYPE
	gRec.area(&recArea)
	gCirc.area(&circArea)

	// PRINT
	fmt.Printf("BEFORE: Rectangle (%.2f x %.2f): Area=%.2f\n", rec.width, rec.height, recArea)
	fmt.Printf("BEFORE: Circle (%.2f): Area=%.2f\n", circ.radius, circArea)

	// INCREASE SIZE BY A FACTOR OF 2 USING POINTER RECIEVER
	fmt.Println("Increase size by a factor of 2")
	gRec.size(2)
	gCirc.size(2)

	// CALCULATE AREA USING A STRUCT TYPE (INCREASED BY 2)
	gRec.area(&recArea)
	gCirc.area(&circArea)

	// PRINT
	fmt.Printf("AFTER:  Rectangle (%.2f x %.2f): Area=%.2f\n", rec.width, rec.height, recArea)
	fmt.Printf("AFTER:: Circle (%.2f): Area=%.2f\n", circ.radius, circArea)

}
