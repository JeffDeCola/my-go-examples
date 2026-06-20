# METHODS POINTERS RECEIVERS EXAMPLE

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Using methods to resize a rectangle and circle using pointer receivers._

tl;dr

```go

// A HIGH HIGH LEVEL VIEW

    func jeff(in1, in2 int) (out1, out2 int) { ... }         // Function  - Values in, values out (many of each)
        a, b := jeff(1, 2)
    func jeff(in inStruct) (out outStruct) { ... }           //           - Using Struct
        out := jeff(in)
    func jeff[T int | float64](in T) (out T) { ... }         //           - Using Generics - Parametric polymorphism
        out := jeff(5)
    func (i inStructOne) jeff() (out outStruct) { ... }      // Method    - Receiver; same jeff() works across many types
        out := one.jeff()
    func doSomething(s shape) float64 { return s.jeff() }    // Interface - Subtype polymorphism (any type with jeff())
        result := doSomething(one)

----------

// SYNTAX:  func (receiver) name(parameters) returnValues
    type rectangle struct{ width, height float64 }
    type circle    struct{ radius float64 }

----------

// FUNCTIONS - Pass by value - A complete copy is made
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

// METHODS - Value receiver - reads a copy - both shapes can just be area()
    func (r rectangle) area() float64 { ... }
    func (c circle) area() float64 { ... }
    recArea  := rec.area()
    circArea := circ.area()

// METHODS-POINTERS-RECEIVERS - Pointer receiver mutates the original
    func (r *rectangle) scale(factor float64) { ... }
    func (c *circle) scale(factor float64) { ... }
    rec.scale(3)
    circ.scale(4)

----------

// INTERFACES - Polymorphism: One func, many types
    type shape interface {
        area() float64
    }
    func getArea(s shape) float64 { return s.area() }
    func (r rectangle) area() float64 { ... }
    func (c circle) area() float64 { ... }
    recArea  := getArea(rec)
    circArea := getArea(circ)

// INTERFACES-POINTERS-RECEIVERS - Scale via interface
    type scaler interface {
        scale(float64)
    }
    func (r *rectangle) scale(factor float64) { ... }
    func (c *circle) scale(factor float64) { ... }
    var sr scaler = &rec
    var sc scaler = &circ
    sr.scale(3)
    sc.scale(4)
```

Examples

* **FUNCTIONS**
  * [functions](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/functions/functions)
  * [functions-pointers-arguments](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/functions/functions-pointers-arguments)
* **METHODS**
  * [methods](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/methods/methods)
  * [methods-pointers-receivers](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/methods/methods-pointers-receivers)
    **YOU ARE HERE**
* **INTERFACES**
  * [interfaces](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/interfaces/interfaces)
  * [interfaces-pointers-receivers](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/interfaces/interfaces-pointers-receivers)

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/methods/methods-pointers-receivers#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/methods/methods-pointers-receivers#run)

Documentation and Reference

* [go spec - method declarations](https://go.dev/ref/spec#Method_declarations)
* [effective go - pointers vs. values](https://go.dev/doc/effective_go#pointers_vs_values)

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

`area()` uses a value receiver, so it reads a copy. But `scale()` uses a
**pointer** receiver (`*rectangle`, not `rectangle`), so it works on the original
and can change it in place. **The big win is a value receiver reads, a pointer
receiver mutates.**

```go
func (r *rectangle) scale(factor float64) {
    r.width *= factor
    r.height *= factor
}
func (c *circle) scale(factor float64) {
    c.radius *= factor
}
```

And they are called,

```go
rec.scale(3)
circ.scale(4)
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
The area of the rectangle (30.00 x 15.00) is 450.00
The area of the circle (radius 20.00) is 1256.64
```
