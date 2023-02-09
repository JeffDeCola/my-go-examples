package main

import (
	"fmt"

	geometry "github.com/JeffDeCola/my-go-packages/geometry/shapes"
)

func main() {

	// DEFINE
	rec := geometry.Rectangle{
		Width:  2.4,
		Height: 34.4,
	}
	circ := geometry.Circle{
		Radius: 2.3,
	}
	tri := geometry.Triangle{
		A: 2,
		B: 3.3,
		C: 4,
	}

	// THE INTERFACE
	var gRec geometry.Geometry
	var gCirc geometry.Geometry
	var gTri geometry.Geometry

	// AREA AND PERIMETER
	var recArea, circArea, triArea float64
	var recPerimeter, circPerimeter, triPerimeter float64

	// ASSIGN STRUCT TO INTERFACE
	gRec = &rec
	gCirc = &circ
	gTri = &tri

	// CALCULATE AREA PERIMETER
	gRec.Area(&recArea)
	gRec.Perimeter(&recPerimeter)
	gCirc.Area(&circArea)
	gCirc.Perimeter(&circPerimeter)
	gTri.Area(&triArea)
	gTri.Perimeter(&triPerimeter)

	// PRINT
	fmt.Printf("Rectangle (%.2f x %.2f): Area=%.2f, Perimeter=%.2f\n", rec.Width, rec.Height, recArea, recPerimeter)
	fmt.Printf("Circle (%.2f): Area=%.2f, Perimeter=%.2f\n", circ.Radius, circArea, circPerimeter)
	fmt.Printf("Triangle (%.2f x %.2f x %.2f ): Area=%.2f, Perimeter=%.2f\n", tri.A, tri.B, tri.C, triArea, triPerimeter)

	// INCREASE SIZE BY A FACTOR OF 2 USING POINTER RECIEVER
	fmt.Println("Increase size by a factor of 2")
	gRec.Size(2)
	gCirc.Size(2)
	gTri.Size(2)

	// CALCULATE AREA USING A STRUCT TYPE (INCREASED BY 2)
	gRec.Area(&recArea)
	gRec.Perimeter(&recPerimeter)
	gCirc.Area(&circArea)
	gCirc.Perimeter(&circPerimeter)
	gTri.Area(&triArea)
	gTri.Perimeter(&triPerimeter)

	// PRINT
	fmt.Printf("Rectangle (%.2f x %.2f): Area=%.2f, Perimeter=%.2f\n", rec.Width, rec.Height, recArea, recPerimeter)
	fmt.Printf("Circle (%.2f): Area=%.2f, Perimeter=%.2f\n", circ.Radius, circArea, circPerimeter)
	fmt.Printf("Triangle (%.2f x %.2f x %.2f ): Area=%.2f, Perimeter=%.2f\n", tri.A, tri.B, tri.C, triArea, triPerimeter)

}
