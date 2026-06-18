# METHODS POINTERS RECEIVERS EXAMPLE

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Using methods to resize a rectangle and circle using pointer receivers._

tl;dr

```go
// replace
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
pointer receiver (`*rectangle`, not `rectangle`), so it works on the original
and can change it in place. **The big win: a value receiver reads, a pointer
receiver mutates.**

```go
func (r *rectangle) scale(factor float64) {
    r.width *= factor
    r.height *= factor
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
