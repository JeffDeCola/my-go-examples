# area-shapes-interfaces

_Using an interface to calculate the area of a rectangle and circle._

Other examples using,

* [Functions](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/functions/area-shapes-functions)
* [Methods](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/methods/area-shapes-methods)
* [Interfaces](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/area-shapes-interfaces)
  **<- YOU ARE HERE**

tl;dr,

```go
// Function, method, interface
recArea := areaRectangle(recWidth, recHeight)
recArea := rec.areaRectangle()
recArea := area(rec)
```

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/area-shapes-interfaces#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/area-shapes-interfaces#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/area-shapes-interfaces#test)
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

Calculate the area using an interface,

```go
recArea := area(rec)
```

The rectangle interface,

```go
type areaInterface interface {
    theArea() float64
}

func area(a areaInterface) float64 {
    // CALL THE METHOD - WILL CHOSE RIGHT ONE
    area := a.theArea()
    return area
}
```

where,

```go
func (r Rectangle) theArea() float32 {
    area := r.width * r.height
    return area
}
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
