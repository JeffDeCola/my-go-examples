# shapes-using-methods example

`shapes-using-methods` is an example of
calculating the area and perimeter of
circles, rectangles and triangles using methods.

This example is done using,

* [shapes-using-functions](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/functions/shapes-using-functions)
* shapes-using-methods (This example)
* [shapes-using-interfaces](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/shapes-using-interfaces)

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

This is pretty straight forward using methods.
It is very similar to using functions except we're using a
function receiver type.  Refer to code.

As you can see its similar to a function,

```go
// Get the shape properties
c1Area := c1.circleArea()
c1Perimeter := c1.circlePerimeter()
r1Area := r1.rectangleArea()
r1Perimeter := r1.rectanglePerimeter()
t1Area := t1.triangleArea()
t1Perimeter := t1.trianglePerimeter()
```

You could argue there is no advantage and you'll be right.
But with interfaces, you can reduce the function name to
just area() and parameter().  See interface example.

## RUN

```go
go run shapes-using-methods.go
```
