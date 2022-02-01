# interfaces-pointers-receivers

_Using an interface to calculate the area of a rectangle and circle
by passing pointers and using pointer receivers._

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
  * [interfaces-pointers-arguments](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/interfaces/interfaces-pointers-arguments)
  * [interfaces-pointers-receivers](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/interfaces/interfaces-pointers-receivers)
    **<- YOU ARE HERE**
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

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/area-shapes-interfaces-ptrs-x2#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/area-shapes-interfaces-ptrs-x2#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/area-shapes-interfaces-ptrs-x2#test)
* [AN ILLUSTRATION THAT MAY HELP](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/area-shapes-interfaces-ptrs-x2#an-illustration-that-may-help)

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

Calculate the size of the rectangle using an interface with pointer receiver,

```go
var gRec geometry
var recArea float64
gRec = &rec
gRec.size(2)
```

Where the interface is,

```go
type geometry interface {
    area(*float64)
    size(float64)
}
```

Which will chose the size method that has a **Rectangle struct pointer receiver**,

```go
func (r *Rectangle) size(f float64) {
    r.width = r.width * f
    r.height = r.height * f
}
```

## RUN

To run,

```bash
go run interfaces-pointers-receivers.go
```

## TEST

To create _test files,

```bash
gotests -w -all interfaces-pointers-receivers.go
```

To unit test the code,

```bash
go test -cover ./... 
```

## AN ILLUSTRATION THAT MAY HELP

![IMAGE - methods-interfaces-pointers-receivers.jpg - IMAGE](../../../docs/pics/basic-syntax/methods-interfaces-pointers-receivers.jpg)
