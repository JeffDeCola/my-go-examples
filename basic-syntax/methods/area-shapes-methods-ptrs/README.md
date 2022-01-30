# area-shapes-methods-ptrs

_Using methods to calculate the area of a rectangle and circle
by passing pointers._

Other examples using,

* Functions using
  * [returns](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/functions/area-shapes-functions)
  * [pointers](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/functions/area-shapes-functions-ptrs)
* Methods using
  * [returns](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/methods/area-shapes-methods)
  * [pointers](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/methods/area-shapes-methods-ptrs)
    **<- YOU ARE HERE**
* Interface using
  * [returns](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/area-shapes-interfaces)
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

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/methods/area-shapes-methods-ptrs#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/methods/area-shapes-methods-ptrs#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/methods/area-shapes-methods-ptrs#test)
* [AN ILLUSTRATION THAT MAY HELP](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/methods/area-shapes-methods-ptrs#an-illustration-that-may-help)

Documentation and reference,

* This repos [github webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

I love structs. So do methods.

The benefit of methods over functions is that you use a struct that defines
the characteristics of the shape. It just makes it a lot easier to use a
struct, rather than passing lots of variables around.

Define the rectangle using a struct,

```go
type Rectangle struct {
    width  float64
    height float64
}

rec := Rectangle{2.4, 34.4}
```

Calculate the area using a method,

```go
rec.area(&recArea)
```

Which will chose the area method that has a Rectangle struct,

```go
func (r Rectangle) area(a *float64) {
    *a = r.width * r.height
}
```

## RUN

To run,

```bash
go run area-shapes-methods-ptrs.go
```

## TEST

To create _test files,

```bash
gotests -w -all area-shapes-methods-ptrs.go
```

To unit test the code,

```bash
go test -cover ./... 
```

## AN ILLUSTRATION THAT MAY HELP

![IMAGE - functions-methods-interfaces.jpg - IMAGE](../../../docs/pics/basic-syntax/functions-methods-interfaces.jpg)
