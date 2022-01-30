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

	// CALCULATE AREA AND PERIMETER
	var gRec shapes.Geometry
	var gCirc shapes.Geometry
	var gTri shapes.Geometry

	gRec = rec
	gCirc = circ
	gTri = tri

	recArea := gRec.Area()
	recPerimeter := gRec.Perimeter()
	circArea := gCirc.Area()
	circPerimeter := gCirc.Perimeter()
	triArea := gTri.Area()
	triPerimeter := gTri.Perimeter()

	// PRINT
	fmt.Printf("Rectangle (%.2f x %.2f): Area=%.2f, Perimeter=%.2f\n", rec.Width, rec.Height, recArea, recPerimeter)
	fmt.Printf("Circle (%.2f): Area=%.2f, Perimeter=%.2f\n", circ.Radius, circArea, circPerimeter)
	fmt.Printf("Triangle (%.2f x %.2f x %.2f ): Area=%.2f, Perimeter=%.2f\n", tri.A, tri.B, tri.C, triArea, triPerimeter)

	// COULD PUT THE INTERFACE IN A FUNCTION
	recArea = shapes.GetArea(rec)
	circArea = shapes.GetArea(circ)
	triArea = shapes.GetArea(tri)
	recPerimeter = shapes.GetPerimeter(rec)
	circPerimeter = shapes.GetPerimeter(circ)
	triPerimeter = shapes.GetPerimeter(tri)

	// PRINT
	fmt.Printf("Rectangle (%.2f x %.2f): Area=%.2f, Perimeter=%.2f\n", rec.Width, rec.Height, recArea, recPerimeter)
	fmt.Printf("Circle (%.2f): Area=%.2f, Perimeter=%.2f\n", circ.Radius, circArea, circPerimeter)
	fmt.Printf("Triangle (%.2f x %.2f x %.2f ): Area=%.2f, Perimeter=%.2f\n", tri.A, tri.B, tri.C, triArea, triPerimeter)

}
