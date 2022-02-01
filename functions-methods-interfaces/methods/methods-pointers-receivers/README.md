# methods-pointers-receivers

_Using methods to calculate the area of a rectangle and circle
by passing pointers and using pointer receivers._

Other examples using,

* **FUNCTIONS**
  * [functions](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/functions/functions)
  * [functions-pointers-arguments](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/functions/functions-pointers-arguments)
* **METHODS**
  * [methods](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/methods/methods)
  * [methods-pointers-arguments](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/methods/methods-pointers-arguments)
  * [methods-pointers-receivers](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/methods/methods-pointers-receivers)
    **<- YOU ARE HERE**
* **INTERFACES**
  * [interfaces](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/interfaces/interfaces)
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
        gRec = &rec // Note this
        gRec.size(2)
```

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/methods/methods-pointers-receivers#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/methods/methods-pointers-receivers#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/methods/methods-pointers-receivers#test)
* [AN ILLUSTRATION THAT MAY HELP](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/methods/methods-pointers-receivers#an-illustration-that-may-help)

Documentation and reference,

* This repos [github webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

A method is a function with a special receiver type.

It just makes it a lot easier to use a
struct as a receiver, rather than passing lots of variables around.

In this example, let's use a size method with a pointer receiver
that will increase or decrease the size of the rectangle or circle.

```go
type Rectangle struct {
    width  float64
    height float64
}

rec := Rectangle{2.4, 34.4}
```

Calculate the size of the rectangle using a method with pointer receiver,

```go
rec.size(2)
```

Which will chose the size method that has a Rectangle receiver,

```go
func (r *Rectangle) size(f float64) {
    r.width = r.width * f
    r.height = r.height * f
}
```

## RUN

To run,

```bash
go run methods-pointers-receivers.go
```

## TEST

To create _test files,

```bash
gotests -w -all methods-pointers-receivers.go
```

To unit test the code,

```bash
go test -cover ./... 
```

## AN ILLUSTRATION THAT MAY HELP

![IMAGE - methods-interfaces-pointers-receivers.jpg - IMAGE](../../../docs/pics/basic-syntax/methods-interfaces-pointers-receivers.jpg)
