# area-shapes-functions

_Using functions to calculate the area of a rectangle and circle._

Other examples using,

* [Functions](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/functions/area-shapes-functions)
  **<- YOU ARE HERE**
* [Methods](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/functions/area-shapes-methods)
* [Interfaces](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/functions/area-shapes-interfaces)

tl;dr,

```go
// Function, method, interface
recArea := areaRectangle(recWidth, recHeight)
recArea := rec.areaRectangle()
recArea := area(rec)
```

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/functions/area-shapes-functions#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/functions/area-shapes-functions#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/functions/area-shapes-functions#test)

Documentation and reference,

* This repos [github webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

The rectangle functions,

```go
func areaRectangle(w float64, h float64) float64 {
    area := w * h
    return area
}
```

Define the rectangle,

```go
var recWidth float32 = 2.4
var recHeight float32 = 34.4
```

Calculate using function,

```go
recArea := areaRectangle(recWidth, recHeight)
```

## RUN

To run,

```bash
go run area-shapes-functions.go
```

## TEST

To create _test files,

```bash
gotests -w -all area-shapes-functions.go
```

To unit test the code,

```bash
go test -cover ./... 
```
