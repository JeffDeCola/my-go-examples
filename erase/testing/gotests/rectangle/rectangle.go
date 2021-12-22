// my-go-examples gotests.go

package rectangle

// Rectangle is a Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// RectangleAreaFunction area using function
func RectangleAreaFunction(r Rectangle) float64 {
	return r.Width * r.Height
}

// RectangleAreaMethod area using method
func (r Rectangle) RectangleAreaMethod() float64 {
	return r.Width * r.Height
}
