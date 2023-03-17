# SHAPES PACKAGE EXAMPLE

_Using an interface to calculate and manipulate simple 2D and 3D geometric shapes._

Other Examples

* **FUNCTIONS**
  * [functions](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/functions/functions)
  * [functions-pointers-arguments](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/functions/functions-pointers-arguments)
* **METHODS**
  * [methods](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/methods/methods)
  * [methods-pointers-arguments](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/methods/methods-pointers-arguments)
  * [methods-pointers-receivers](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/methods/methods-pointers-receivers)
* **INTERFACES**
  * [interfaces](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/interfaces/interfaces)
  * [interfaces-pointers-arguments](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/interfaces/interfaces-pointers-arguments)
  * [interfaces-pointers-receivers](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/interfaces/interfaces-pointers-receivers)
  * [shapes-package](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/interfaces/shapes-package)
    **<- YOU ARE HERE**

tl;dr,

```go
// SYNTAX
    // (receiver type) func (parameters) (return types)

// FUNCTIONS

    // USING RETURNS
        func areaRectangle(w float64, h float64) float64 {
        recArea := areaRectangle(recWidth, recHeight)
    // USING POINTERS IN ARGUMENTS
        func areaRectangle(w float64, h float64, a *float64) {
        var recArea float64
        areaRectangle(recWidth, recHeight, &recArea)

// METHODS

    // USING RETURNS
        func (r Rectangle) area() float64 {
        rec := Rectangle{2.4, 34.4}
        recArea := rec.area()
    // USING POINTERS IN PARAMETERS
        func (r Rectangle) area(a *float64) {
        rec := Rectangle{2.4, 34.4}
        var recArea float64
        rec.area(&recArea)
     // USING POINTERS IN RECEIVERS
        func (r *Rectangle) size(f factor) {
        rec := Rectangle{2.4, 34.4} // We want to change this
        rec.size(2)

// INTERFACES

    // USING RETURNS
        func (geometry).area() float64 // Abstract representation
        rec := Rectangle{2.4, 34.4}
        var gRec geometry
        gRec = rec
        recArea := gRec.area()
    // USING POINTERS IN ARGUMENTS
        func (geometry).area()(*float64) // Abstract representation
        rec := Rectangle{2.4, 34.4}
        var gRec geometry
        var recArea float64
        gRec = rec
        gRec.area(&recArea)
    // USING POINTERS IN RECEIVERS
        func (geometry).size()(*float64) // Abstract representation
        rec := Rectangle{2.4, 34.4} // We want to change this
        var gRec geometry
        gRec = &rec // Note this
        gRec.size(2)
```

Table of Contents

* [PREREQUISITES](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/interfaces/shapes-package#prerequisites)
* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/interfaces/shapes-package#overview)
* [EXAMPLE](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/interfaces/shapes-package#example)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/interfaces/shapes-package#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/interfaces/shapes-package#test)
* [AN ILLUSTRATION THAT MAY HELP](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/interfaces/shapes-package#an-illustration-that-may-help)

Documentation and Reference

* My shapes
  [package](https://github.com/JeffDeCola/my-go-packages/tree/master/geometry/shapes)

## OVERVIEW

The use of interfaces is to create a common type that can be used by
multiple types.

```go
func GetArea(t TwoDCalculations) float64
func GetPerimeter(t TwoDCalculations) float64
func GetVolume(t ThreeDCalculations) float64
func GetSurfaceArea(t ThreeDCalculations) float64
func ChangeSize(t ShapeManipulations, f float64)
```

## PREREQUISITES

```go
go get -u -v github.com/JeffDeCola/my-go-packages
go install -v github.com/JeffDeCola/my-go-packages@latest
```

## EXAMPLE

```go
package main

import (
    "fmt"

    "github.com/JeffDeCola/my-go-packages/geometry/shapes"
)

func main() {

    // Create a Rectangle and Circle type
    r := shapes.Rectangle{Width: 10, Height: 10}
    c := shapes.Circle{Radius: 5}

    // Get the area (using the interface)
    //a := shapes.GetArea(r)
    a := r.Area()
    fmt.Println("Area of rectangle =", a)
    a = shapes.GetArea(c)
    fmt.Println("Area of circle =", a)

    // Change the size (x2)
    shapes.ChangeSize(&r, 2)
    shapes.ChangeSize(&c, 2)

    // Get the area (using the interface)
    a = shapes.GetArea(r)
    fmt.Println("Area of rectangle =", a)
    a = shapes.GetArea(c)
    fmt.Println("Area of circle =", a)

}
```

## RUN

To run,

```bash
go run main.go
```

## TEST

To create _test files,

```bash
gotests -w -all main.go
```

To unit test the code,

```bash
go test -cover ./...
```

## AN ILLUSTRATION THAT MAY HELP

![IMAGE - methods-interfaces-pointers-receivers.jpg - IMAGE](../../../docs/pics/functions-methods-interfaces/methods-interfaces-pointers-receivers.jpg)
