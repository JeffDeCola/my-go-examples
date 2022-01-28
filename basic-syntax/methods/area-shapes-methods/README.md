# area-shapes-methods

_Using methods to calculate the area of a rectangle and circle._

Other examples using,

* Functions
* Methods **<- YOU ARE HERE**
* Interfaces

tl;dr,

```go
// Function, method, interface
recArea := areaRectangle(recWidth, recHeight)
recArea := rec.areaRectangle()
recArea := area(rec)
```

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/methods/area-shapes-methods#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/methods/area-shapes-methods#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/methods/area-shapes-methods#test)

Documentation and reference,

* This repos [github webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

I love structs. So do methods.

The benefit of methods over functions is that you use a struct that defines
the characteristics of the shape. It just makes it a lot easier to use a
struct, rather than passing lots of variables around.

The rectangle methods,

```go
type Rectangle struct {
    width  float32
    height float32
}

func (r Rectangle) areaRectangle() float32 {
    area := r.width * r.height
    return area
}
```

Define the rectangle,

```go
rec := Rectangle{2.4, 34.4}
```

Calculate using method,

```go
recArea := rec.areaRectangle()
```

## RUN

To run,

```bash
go run area-shapes-methods.go
```

## TEST

To create _test files,

```bash
gotests -w -all area-shapes-methods.go
```

To unit test the code,

```bash
go test -cover ./... 
```
