# shapes-using-interfaces example

`shapes-using-interfaces` is an example of
calculating the area and perimeter of
circles, rectangles and triangles using interfaces.

This example is done using,

* [shapes-using-functions](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/functions/shapes-using-functions)
* [shapes-using-methods](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/methods/shapes-using-methods)
* shapes-using-interfaces  (This example)

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

You can see the function and methods are basically the same thing.  But when
we use interfaces we can give a function the type (struct) and the interface
will figure it out.

Ao we go from,

```go
// Get the shape properties
c1Area := circleArea(c1)
c1Perimeter := circlePerimeter(c1)
r1Area := rectangleArea(r1)
r1Perimeter := rectanglePerimeter(r1)
t1Area := triangleArea(t1)
t1Perimeter := trianglePerimeter(t1)
```

To a simple form,

```go
// Get the shape properties
c1Area := area(c1)
c1Perimeter := perimeter(c1)
r1Area := area(r1)
r1Perimeter := perimeter(r1)
t1Area := area(t1)
t1Perimeter := perimeter(t1)
```

## TWO INTERFACES VS ONE INTERFACE

The first example `shapes-using-interfaces1.go` uses 2 interfaces,

```go
type areaer interface {
    area() float64
}

type perimeterer interface {
    perimeter() float64
}
```

The second `shapes-using-interfaces2.go` uses more traditional 1 interface,

```go
type shapes interface {
    area() float64
    perimeter() float64
```

The end result is the same.

## RUN

```bash
go run shapes-using-interfaces1.go
go run shapes-using-interfaces2.go
```
