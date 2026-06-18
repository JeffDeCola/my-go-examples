# FUNCTIONS EXAMPLE

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Using functions to calculate the area of a rectangle and circle._

tl;dr

```go
// SYNTAX:  func (receiver) name(parameters) returnValues
    type rectangle struct{ width, height float64 }
    type circle    struct{ radius float64 }

----------

// FUNCTIONS - Pass by value -  A complete copy is made
    func areaRectangle(r rectangle) float64 { ... }
    func areaCircle(c circle) float64 { ... }
    recArea  := areaRectangle(rec)
    circArea := areaCircle(circ)

// FUNCTIONS-POINTERS-ARGUMENTS - Pass by pointer
    func scaleRectangle(r *rectangle, factor float64) { ... }
    func scaleCircle(c *circle, factor float64) { ... }
    scaleRectangle(&rec, 3)
    scaleCircle(&circ, 4)

----------

// METHODS - Value receiver - Reads a copy - Both shapes can just be area()
    func (r rectangle) area() float64 { ... }
    func (c circle) area() float64 { ... }
    recArea  := rec.area()
    circArea := circ.area()

// METHODS-POINTERS-RECEIVERS - Pointer receiver mutates the original
    func (r *rectangle) scale(factor float64) { ... }
    rec.scale(2)

----------

// INTERFACES - a func accepts the interface; any type with area() satisfies it.
// This is polymorphism: one func, many types, the right area() chosen at runtime.
    type geometry interface { area() float64 }
    func getArea(g geometry) float64 { ... }
    recArea  := getArea(rec)    // rectangle
    circArea := getArea(circ)   // circle - one func, dispatch by type

// INTERFACES-POINTERS-RECEIVERS - scale has a *receiver, so only *rectangle is
// in the method set -> assign &rec, not rec
    type geometry interface {
        area() float64
        scale(float64)
    }
    var g geometry = &rec
    g.scale(2)
```

Examples

* **FUNCTIONS**
  * [functions](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/functions/functions)
    **YOU ARE HERE**
  * [functions-pointers-arguments](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/functions/functions-pointers-arguments)
* **METHODS**
  * [methods](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/methods/methods)
  * [methods-pointers-receivers](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/methods/methods-pointers-receivers)
* **INTERFACES**
  * [interfaces](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/interfaces/interfaces)
  * [interfaces-pointers-receivers](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/interfaces/interfaces-pointers-receivers)

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/functions/functions#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/functions/functions#run)

Documentation and Reference

* [go spec - function declarations](https://go.dev/ref/spec#Function_declarations)
* [effective go - functions](https://go.dev/doc/effective_go#functions)

## OVERVIEW

Given these structs,

```go
type rectangle struct {
    width  float64
    height float64
}

type circle struct {
    radius float64
}
```

Functions are very simple; just input and output.
We pass the rectangle by value so a complete copy is made.

```go
func areaRectangle(r rectangle) float64 {
    area := r.width * r.height
    return area
}
```

And they are called,

```go
recArea  := areaRectangle(rec)
circArea := areaCircle(circ)
```

## RUN

To run,

```bash
go run main.go
```

The output should look like,

```bash
The area of the rectangle (10.00 x 5.00) is 50.00
The area of the circle (radius 5.00) is 78.54
```
