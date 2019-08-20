# shapes-using-functions example

`shapes-using-functions` is an example of
calculating the area and perimeter of
circles, rectangles and triangles using functions.

* shapes-using-functions (This example)
* [shapes-using-methods](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/methods/shapes-using-methods)
* [shapes-using-interfaces](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/shapes-using-interfaces)

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

This is pretty straight forward using functions.
You just call all the different functions,

```go
// Get the shape properties
c1Area := circleArea(c1)
c1Perimeter := circlePerimeter(c1)
r1Area := rectangleArea(r1)
r1Perimeter := rectanglePerimeter(r1)
t1Area := triangleArea(t1)
t1Perimeter := trianglePerimeter(t1)
```

We will see, we can simplify this using an interface.

## RUN

```bash
go run shapes-using-functions.go
```
