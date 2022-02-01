# interfaces

_Using an interface to calculate the area of a rectangle and circle._

Other examples using,

* **FUNCTIONS**
  * [functions](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/functions/functions)
  * [functions-pointers-arguments](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/functions/functions-pointers-arguments)
* **METHODS**
  * [methods](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/methods/methods)
  * [methods-pointers-arguments](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/methods/methods-pointers-arguments)
  * [methods-pointers-receivers](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/methods/methods-pointers-receivers)
* **INTERFACES**
  * [interfaces](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/interfaces/interfaces)
    **<- YOU ARE HERE**
  * [interfaces-pointers-arguments](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/interfaces/interfaces-pointers-arguments)
  * [interfaces-pointers-receivers](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/interfaces/interfaces-pointers-receivers)
  * [geometry-package](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/interfaces/geometry-package)
  
tl;dr,

```go
// SYNTAX
    // (receiver) func (arguments) (return arguments)

// FUNCTIONS

    // USING RETURNS
        func areaRectangle(w float64, h float64) float64 {
        // To Use
        recArea := areaRectangle(recWidth, recHeight)
    // USING POINTERS IN ARGUMENTS
        func areaRectangle(w float64, h float64, a *float64) {
        // To Use
        var recArea float64
        areaRectangle(recWidth, recHeight, &recArea)

// METHODS

    // USING RETURNS
        func (r Rectangle) area() float64 {
        // To Use
        rec := Rectangle{2.4, 34.4}
        recArea := rec.area()
        // To Use
    // USING POINTERS IN ARGUMENTS
        func (r Rectangle) area(a *float64) {
        // To Use
        rec := Rectangle{2.4, 34.4}
        var recArea float64
        rec.area(&recArea)
     // USING POINTERS IN RECEIVERS
        func (r *Rectangle) size(f factor) {
        // To Use
        rec := Rectangle{2.4, 34.4}
        rec.size(2)


// INTERFACES
    
    // USING RETURNS
        func (geometry).area() float64 // Abstract representation
        // To Use
        rec := Rectangle{2.4, 34.4}
        var gRec geometry
        gRec = rec
        recArea := gRec.area()
    // USING POINTERS IN ARGUMENTS
        func (geometry).area()(*float64) // Abstract representation
        // To Use
        rec := Rectangle{2.4, 34.4}
        var gRec geometry
        var recArea float64
        gRec = rec
        gRec.area(&recArea)
    // USING POINTERS IN RECEIVERS
        func (geometry).size()(*float64) // Abstract representation
        // To Use
        rec := Rectangle{2.4, 34.4}
        var gRec geometry
        var recArea float64
        gRec = &rec // Note this
        gRec.size(2)
```

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/area-shapes-interfaces#overview)
  * [WRAP INTERFACE IN A FUNCTION](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/area-shapes-interfaces#wrap-interface-in-a-function)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/area-shapes-interfaces#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/area-shapes-interfaces#test)
* [AN ILLUSTRATION THAT MAY HELP](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/area-shapes-interfaces#an-illustration-that-may-help)

Documentation and reference,

* This repos [github webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

Define the rectangle using a struct,

```go
type Rectangle struct {
    width  float64
    height float64
}

rec := Rectangle{2.4, 34.4}
```

Calculate the area using a interface,

```go
var gRec geometry
gRec = rec
recArea := gRec.area()
```

Where the interface is,

```go
type geometry interface {
    area() float64
    perimeter()float64
}
```

Which will chose the area method that has a Rectangle struct,

```go
func (r Rectangle) area() float64 {
    area := r.width * r.height
    return area
}
```

### WRAP INTERFACE IN A FUNCTION

You could create a function,

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
go run interfaces.go
```

## TEST

To create _test files,

```bash
gotests -w -all interfaces.go
```

To unit test the code,

```bash
go test -cover ./... 
```

## AN ILLUSTRATION THAT MAY HELP

![IMAGE - functions-methods-interfaces.jpg - IMAGE](../../../docs/pics/basic-syntax/functions-methods-interfaces.jpg)
