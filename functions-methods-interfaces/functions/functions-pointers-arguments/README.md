# FUNCTIONS POINTERS ARGUMENTS EXAMPLE

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Using functions to resize a rectangle and circle by passing pointers._

tl;dr

```go
// replace
```

Examples

* **FUNCTIONS**
  * [functions](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/functions/functions)
  * [functions-pointers-arguments](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/functions/functions-pointers-arguments)
    **YOU ARE HERE**
* **METHODS**
  * [methods](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/methods/methods)
  * [methods-pointers-receivers](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/methods/methods-pointers-receivers)
* **INTERFACES**
  * [interfaces](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/interfaces/interfaces)
  * [interfaces-pointers-receivers](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/interfaces/interfaces-pointers-receivers)

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/functions/functions-pointers-arguments#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/functions/functions-pointers-arguments#run)

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

These functions take a pointer and change the struct in place.
No value comes back.

```go
func scaleRectangle(r *rectangle, factor float64) {
    r.width *= factor
    r.height *= factor
}
```

And they are called,

```go
scaleRectangle(&rec, 3)
scaleCircle(&circ, 4)
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
