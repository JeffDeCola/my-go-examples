// my-go-examples main.go

package main

import (
	"fmt"

	rec "./rectangle"
)

func main() {

	r1 := rec.Rectangle{Width: 5, Height: 3}
	r2 := rec.Rectangle{Width: 5, Height: 3}

	r1Area := rec.RectangleAreaFunction(r1)
	r2Area := r2.RectangleAreaMethod()

	fmt.Printf("The area of a rectangle with width %v and heigh %v is %f\n", r1.Width, r1.Height, r1Area)
	fmt.Printf("The area of a rectangle with width %v and heigh %v is %f\n", r2.Width, r2.Height, r2Area)

}
