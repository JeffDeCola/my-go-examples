# shapes-package

_Using an interface to calculate the area and perimeter of a rectangle,
circle and triangle via a shapes package._

Other examples using,

* Functions using
  * [returns](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/functions/area-shapes-functions)
  * [pointers](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/functions/area-shapes-functions-ptrs)
* Methods using
  * [returns](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/methods/area-shapes-methods)
  * [pointers](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/methods/area-shapes-methods-ptrs)
* Interface using
  * [returns](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/area-shapes-interfaces)
  * [pointers](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/area-shapes-interfaces-ptrs)
  * [pointers](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/area-shapes-interfaces-ptrs-x2)
    (multiply structs by 2)
  * [returns using a package](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/shapes-package)
    **<- YOU ARE HERE**
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

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/shapes-package#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/shapes-package#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/shapes-package#test)
* [AN ILLUSTRATION THAT MAY HELP](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/shapes-package#an-illustration-that-may-help)

Documentation and reference,

* This repos [github webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

The power of interfaces really shine when you have a package.
Because you can have one interface that does a lot of things
like a geometry shapes package.

Defines the shapes using a struct,

```go
rec := shapes.Rectangle{
    Width:  2.4,
    Height: 34.4,
}
circ := shapes.Circle{
    Radius: 2.3,
}
tri := shapes.Triangle{
    A: 2,
    B: 3.3,
    C: 4,
}
```

Calculate area and perimeter using a interface,

```go
var gRec shapes.Geometry
var gCirc shapes.Geometry
var gTri shapes.Geometry

gRec = rec
gCirc = circ
gTri = tri

recArea := gRec.Area()
recPerimeter := gRec.Perimeter()
circArea := gCirc.Area()
circPerimeter := gCirc.Perimeter()
triArea := gTri.Area()
triPerimeter := gTri.Perimeter()
```

Where the interface is,

```go
type Geometry interface {
    Area() float64
    Perimeter() float64
}
```

## RUN

To run,

```bash
go run shapes-package.go
```

## TEST

To create _test files,

```bash
cd shapes
gotests -w -all shapes.go
```

To unit test the code,

```bash
go test -cover ./... 
```

## AN ILLUSTRATION THAT MAY HELP

![IMAGE - functions-methods-interfaces.jpg - IMAGE](../../../docs/pics/basic-syntax/functions-methods-interfaces.jpg)
