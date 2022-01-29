# area-shapes-interfaces

_Using an interface to calculate the area of a rectangle and circle._

Other examples using,

* [Functions](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/functions/area-shapes-functions)
* [Methods](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/methods/area-shapes-methods)
* [Interfaces](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/area-shapes-interfaces)
  **<- YOU ARE HERE**

tl;dr,

```go
// FUNCTION
recArea := areaRectangle(recWidth, recHeight)
// METHOD
rec := Rectangle{2.4, 34.4}
recArea := rec.area()
// INTERFACE
var g geometry
g = rec
recArea := g.area()
```

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/area-shapes-interfaces#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/area-shapes-interfaces#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/area-shapes-interfaces#test)
* [TWO WAYS TO CODE INTERFACES](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/area-shapes-interfaces#two-ways-to-code-interfaces)
* [AN ILLUSTRATION THAT MAY HELP](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/area-shapes-interfaces#an-illustration-that-may-help)

Documentation and reference,

* This repos [github webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

Lets have one function that can get the area for both
a rectangle and a circle. That's an interface.

Define the rectangle using a struct,

```go
type Rectangle struct {
    width  float64
    height float64
}

rec := Rectangle{2.4, 34.4}
```

Calculate the area using area method,

```go
var g geometry
g = rec
recArea := g.area()
```

Where the interface is,

```go
type geometry interface {
    area() float64
}
```

Or you could create a function,

```go
func theArea(g geometry) float64 {
    area := g.area()
    return area
}
```

To Calculate the area,

```go
recArea := theArea(rec)
```

## RUN

To run,

```bash
go run area-shapes-interfaces.go
```

## TEST

To create _test files,

```bash
gotests -w -all area-shapes-interfaces.go
```

To unit test the code,

```bash
go test -cover ./... 
```

## AN ILLUSTRATION THAT MAY HELP

![IMAGE - functions-methods-interfaces.jpg - IMAGE](../../../docs/pics/basic-syntax/functions-methods-interfaces.jpg)
