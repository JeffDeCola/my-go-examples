# area-shapes-functions

_Using functions to calculate the area of a rectangle and circle._

Other examples using,

* [Functions](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/functions/area-shapes-functions)
  **<- YOU ARE HERE**
* [Methods](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/methods/area-shapes-methods)
* [Interfaces](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/area-shapes-interfaces)

tl;dr,

```go
// FUNCTION
recArea := areaRectangle(recWidth, recHeight)
// METHOD
rec := Rectangle{2.4, 34.4}
recArea := rec.areaRectangle()
// INTERFACE
var g geometry
g = rec
recArea := g.area()
```

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/functions/area-shapes-functions#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/functions/area-shapes-functions#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/functions/area-shapes-functions#test)
* [AN ILLUSTRATION THAT MAY HELP](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/functions/area-shapes-functions#an-illustration-that-may-help)

Documentation and reference,

* This repos [github webpage](https://jeffdecola.github.io/my-go-examples/)

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
go run area-shapes-functions.go
```

## TEST

To create _test files,

```bash
gotests -w -all area-shapes-functions.go
```

To unit test the code,

```bash
go test -cover ./... 
```

## AN ILLUSTRATION THAT MAY HELP

![IMAGE - functions-methods-interfaces.jpg - IMAGE](../../../docs/pics/basic-syntax/functions-methods-interfaces.jpg)
