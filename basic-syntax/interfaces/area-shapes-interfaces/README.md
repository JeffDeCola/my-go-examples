# area-shapes-interfaces

_Using an interface to calculate the area of a rectangle and circle._

Other examples using,

* Functions using
  * [returns](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/functions/area-shapes-functions)
  * [pointers](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/functions/area-shapes-functions-ptrs)
* Methods using
  * [returns](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/methods/area-shapes-methods)
  * [pointers](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/methods/area-shapes-methods-ptrs)
* Interface using
  * [returns](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/area-shapes-interfaces)
    **<- YOU ARE HERE**
  * [pointers](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/area-shapes-interfaces-ptrs)
  * [returns using a package](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/shapes-package)
  * [pointers using a package](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/shapes-package-ptrs)
  
tl;dr,

```go
// FUNCTION
    // USING RETURN
    recArea := areaRectangle(recWidth, recHeight)
    // USING POINTER
    var recArea float64
    areaRectangle(recWidth, recHeight, &recArea)
// METHOD
    // USING RETURN
    rec := Rectangle{2.4, 34.4}
    recArea := rec.area()
    // USING POINTER
    rec := Rectangle{2.4, 34.4}
    var recArea float64
    rec.area(&recArea)
// INTERFACE
    // USING RETURN
    rec := Rectangle{2.4, 34.4}
    var gRec geometry
    gRec = rec
    recArea := gRec.area()
    // USING POINTER
    rec := Rectangle{2.4, 34.4}
    var gRec geometry
    var recArea float64
    gRec = rec
    gRec.area(&recArea)
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
