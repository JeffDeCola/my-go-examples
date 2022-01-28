# shapes-package

_Using an interface to calculate the area and perimeter of a rectangle,
circle and triangle via a shapes package._

tl;dr,

```go
// Function, method, interface
recArea := areaRectangle(recWidth, recHeight)
recArea := rec.areaRectangle()
recArea := area(rec)
```

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/shapes-package#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/shapes-package#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/shapes-package#test)

Documentation and reference,

* This repos [github webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

The goal is to just have an interface that calculates area,

```go
recArea := area(rec)
circArea := area(circ)
triArea := area(tri)
```

Or perimeter,

```go
recPerimeter := area(rec)
circPerimeter := area(circ)
triPerimeter := area(tri)
```

## RUN

To run,

```bash
go run shapes-package.go
```

## TEST

To create _test files,

```bash
gotests -w -all shapes-package.go
```

To unit test the code,

```bash
go test -cover ./... 
```
