# METHODS EXAMPLE

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Using methods to calculate the area of a rectangle and circle._

Examples

* **FUNCTIONS**
  * [functions](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/functions/functions)
  * [functions-pointers-arguments](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/functions/functions-pointers-arguments)
* **METHODS**
  * [methods](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/methods/methods)
    **YOU ARE HERE**
  * [methods-pointers-receivers](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/methods/methods-pointers-receivers)
* **INTERFACES**
  * [interfaces](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/interfaces/interfaces)
  * [interfaces-pointers-receivers](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/interfaces/interfaces-pointers-receivers)

tl;dr

```go
// replace
```

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/methods/methods#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/methods/methods#run)

Documentation and Reference

* [go spec - method declarations](https://go.dev/ref/spec#Method_declarations)
* [effective go - methods](https://go.dev/doc/effective_go#methods)

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

Methods are also very simple with the addition of a receiver
as well as normal input and output.
We pass the rectangle by value in the receiver so a complete copy is made.
**The big win is area() function can be used by methods.**

```go
func (r rectangle) area() float64 {
    area := r.width * r.height
    return area
}
```

And they are called,

```go
recArea  := rec.area()
circArea := circ.area()
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
