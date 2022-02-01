# geometry-package

 _Using an interface to calculate the area and perimeter of a rectangle,
circle and triangle via a geometry package.
It uses pointers for arguments and receivers._

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
  * [geometry-package](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/interfaces/geometry-package)
    **<- YOU ARE HERE**

tl;dr,

```go

```

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/shapes-package-ptrs#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/shapes-package-ptrs#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/shapes-package-ptrs#test)
* [AN ILLUSTRATION THAT MAY HELP](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/shapes-package-ptrs#an-illustration-that-may-help)

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
var recArea, circArea, triArea float64
var recPerimeter, circPerimeter, triPerimeter float64

gRec = rec
gCirc = circ
gTri = tri

gRec.Area(&recArea)
gRec.Perimeter(&recPerimeter)
gCirc.Area(&circArea)
gCirc.Perimeter(&circPerimeter)
gTri.Area(&triArea)
gTri.Perimeter(&triPerimeter)
```

Where the interface is,

```go
type Geometry interface {
    Area(*float64) 
    Perimeter(*float64)
}
```

## RUN

To run,

```bash
go run geometry-package.go
```

## TEST

To create _test files,

```bash
cd shapes
gotests -w -all geometry-package.go
```

To unit test the code,

```bash
go test -cover ./... 
```

## AN ILLUSTRATION THAT MAY HELP

![IMAGE - methods-interfaces-pointers-receivers.jpg - IMAGE](../../../docs/pics/basic-syntax/methods-interfaces-pointers-receivers.jpg)
