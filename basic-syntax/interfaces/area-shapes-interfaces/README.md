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

Documentation and reference,

* This repos [github webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

Lets have one interface that can get the area for both
a rectangle and a circle.

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

Calculate using interface,

```go
recArea := area(rec)
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
