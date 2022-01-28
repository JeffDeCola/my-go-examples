package main

import (
	"fmt"
	"shapes-package/shapes"
)

func main() {

	// DEFINE
	rec := shapes.Rectangle{
		Width:  2.4,
		Height: 34.4,
	}
	circ := shapes.Circle{
		Radius: 2.3,
	}
	tri := shapes.Triangle{
		A: 2,
		B: 3.3,
		C: 4,
	}

	// CALCULATE AREA - THE POWER OF THE INTERFACE
	recArea := shapes.Area(rec)
	circArea := shapes.Area(circ)
	triArea := shapes.Area(tri)

	// CALCULATE PERIMETER - THE POWER OF THE INTERFACE
	recPerimeter := shapes.Perimeter(rec)
	circPerimeter := shapes.Perimeter(circ)
	triPerimeter := shapes.Perimeter(tri)

	// PRINT
	fmt.Printf("Rectangle (%.2f x %.2f): Area=%.2f, Perimeter=%.2f\n", rec.Width, rec.Height, recArea, recPerimeter)
	fmt.Printf("Circle (%.2f): Area=%.2f, Perimeter=%.2f\n", circ.Radius, circArea, circPerimeter)
	fmt.Printf("Triangle (%.2f x %.2f x %.2f ): Area=%.2f, Perimeter=%.2f\n", tri.A, tri.B, tri.C, triArea, triPerimeter)
}
