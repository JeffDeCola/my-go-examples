# FUNCTIONS EXAMPLE

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Using functions to calculate the area of a rectangle and circle._

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

tl;dr

```go
// SYNTAX:  func (receiver) name(parameters) returnValues

type rectangle struct{ width, height float64 }
type circle    struct{ radius float64 }

// FUNCTIONS - standalone; the struct goes in, a value comes out.
// Functions can't share a name, so they're areaRectangle/areaCircle.
func areaRectangle(r rectangle) float64 { return r.width * r.height }
func areaCircle(c circle) float64       { return math.Pi * c.radius * c.radius }
recArea  := areaRectangle(rec)
circArea := areaCircle(circ)

// FUNCTIONS-POINTERS-ARGUMENTS - pass a *pointer to mutate the caller's value
func scale(r *rectangle, f float64) { r.width *= f; r.height *= f }
scale(&rec, 2)

// METHODS - bound to a receiver, so BOTH shapes can just be area()
func (r rectangle) area() float64 { return r.width * r.height }
func (c circle)    area() float64 { return math.Pi * c.radius * c.radius }
recArea  := rec.area()
circArea := circ.area()

// METHODS-POINTERS-RECEIVERS - pointer receiver mutates the receiver in place.
// value receiver reads; pointer receiver mutates.
func (r *rectangle) scale(f float64) { r.width *= f; r.height *= f }
rec.scale(2)

// INTERFACES - a func accepts the interface; any type with area() satisfies it.
// This is polymorphism: one func, many types, the right area() chosen at runtime.
type geometry interface { area() float64 }
func getArea(g geometry) float64 { return g.area() }
recArea  := getArea(rec)    // rectangle
circArea := getArea(circ)   // circle - one func, dispatch by type

// INTERFACES-POINTERS-RECEIVERS - scale has a *receiver, so only *rectangle is
//                                 in the method set -> assign &rec, not rec
type geometry interface {
    area() float64
    scale(float64)
}
var g geometry = &rec
g.scale(2)
```

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/functions/functions#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/functions/functions#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/functions/functions#test)
* [AN ILLUSTRATION THAT MAY HELP](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/functions/functions#an-illustration-that-may-help)

## OVERVIEW

Define the rectangle,

```go
var recWidth float64 = 2.4
var recHeight float34 = 34.4
```

Calculate the area using a function,

```go
recArea := areaRectangle(recWidth, recHeight)
```

The rectangle function,

```go
func areaRectangle(w float64, h float64) float64 {
    area := w * h
    return area
}
```

## RUN

To run,

```bash
go run main.go
```

The out should look like

