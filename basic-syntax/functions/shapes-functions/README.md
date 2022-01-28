# shapes

_Using functions to define and calculate the area and perimeter of geometric shapes._

Other examples using,

* Functions **<- YOU ARE HERE**
* Methods
* Interfaces

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/functions/shapes#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/functions/shapes#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/functions/shapes#test)

Documentation and reference,

* This repos [github webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

The rectangle functions,

```go
func areaRectangle(w float32, h float32) float32 {
    area := w * h
    return area
}

func perimeterRectangle(w float32, h float32) float32 {
    perimeter := 2 * (w + h)
    return perimeter
}
```

Define the rectangle,

```go
var recWidth float32 = 2.4
var recHeight float32 = 34.4
```

Calculate using function,

```go
// CALCULATE AREA & PERIMETER
recArea := areaRectangle(recWidth, recHeight)
recPerimeter := perimeterRectangle(recWidth, recHeight)
```

## RUN

To run,

```bash
go run shapes-functions.go
```

## TEST

To create _test files,

```bash
gotests -w -all shapes-functions.go
```

To unit test the code,

```bash
go test -cover ./... 
```
