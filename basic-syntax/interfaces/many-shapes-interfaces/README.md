# shapes

_Using interfaces to define and calculate the area and perimeter of geometric shapes._

Other examples using,

* Functions
* Methods
* Interfaces **<- YOU ARE HERE**

Table of Contents,

* tbd

Documentation and reference,

* This repos [github webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

Lets go one more level and have 1 function that can get the area and
perimeter of many different shapes.

```go
type Rectangle struct {
    width  float32
    height float32
}

func (r Rectangle) areaRectangle() float32 {
    area := r.width * r.height
    return area
}

func (r Rectangle)  perimeterRectangle() float32 {
    perimeter := 2 * (r.width + r.height)
    return perimeter
}
```

Define the rectangle,

```go
rec := Rectangle{2.4, 34.4}
```

Calculate using interface,

```go
// CALCULATE AREA & PERIMETER
recArea := area(rec)
recPerimeter := perimeter(rec)
```

## RUN

To run,

```bash
go run shapes-interfaces.go
```

## TEST

To create _test files,

```bash
gotests -w -all shapes-interfaces.go
```

To unit test the code,

```bash
go test -cover ./... 
```
