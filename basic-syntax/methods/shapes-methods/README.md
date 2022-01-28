# shapes

_Using methods to define and calculate the area and perimeter of geometric shapes._

Other examples using,

* Functions
* Methods **<- YOU ARE HERE**
* Interfaces

Table of Contents,

* tbd

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

func (r Rectangle)  perimeterRectangle() float32 {
    perimeter := 2 * (r.width + r.height)
    return perimeter
}
```

Define the rectangle,

```go
rec := Rectangle{2.4, 34.4}
```

Calculate using method,

```go
// CALCULATE AREA & PERIMETER
recArea := rec.areaRectangle()
recPerimeter := rec.perimeterRectangle()
```

## RUN

To run,

```bash
go run shapes-methods.go
```

## TEST

To create _test files,

```bash
gotests -w -all shapes-methods.go
```

To unit test the code,

```bash
go test -cover ./... 
```
