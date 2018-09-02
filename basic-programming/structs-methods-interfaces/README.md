# structs-methods-interfaces example

`structs-methods-interfaces` _is an example of structs, methods and interfaces._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## EXAMPLE

Consider a coordinate system with a rectangle  `rect1` and a circle `circ1`.

`rect1` lives at coordinates:

* x1 = 1
* x2 = 11
* y1 = 1
* y2 = 11

`circ1` lives at coordinates:

* x = 10
* y = 10
* r = 5

## WITHOUT STRUCTS

Just using go datatypes will eventually become tedius.

To describe rect1 you need to keep track of its coordinates
and pass them to functions.

```go
// rect1
var rx1, ry1 float64 = 1, 1
var rx2, ry2 float64 = 11, 11
```

And then pass this to the rectangleArea Function,

```go
rectangleArea(rx1, ry1, rx2, ry2)
```

Just Ugly.

Then your function looks like,

```go
func rectangleArea(x1, y1, x2, y2 float64) float64 {
  l := distance(x1, y1, x1, y2)
  w := distance(x1, y1, x2, y1)
  return l * w
}
```

## STRUCTS

Structs allow you to keep all the types about a rectangle or a circle together,

```go
type Rectangle struct {
    x1 float64
    x2 float64
    y1 float64
    y2  float64
}
```

Now rect1 is describes as,

```go
// rect1
rect1 := Rectangle{x1: 1, x2: 11, y1: 1, y2: 11}
```

And then pass this to the rectangleArea Function,

```go
rectangleArea(rect1)
```

Not Ugly.

Then in your function you access the field by using `.`,

```go
func rectangleArea(r Rectangle) float64 {
  l := distance(r.x1, r.y1, r.x1, r.y2)
  w := distance(r.x1, r.y1, r.x2, r.y1)
  return l * w
}
```

## FUNCTIONS

As seen above Functions just use input and output.

But we can do better using a Method.

## METHODS

Methods are are a special type of Function.
Where Functions stand on their own, Methods
are declaired with a receiver type, hence Methods
are attached to data (type).

From,

```go
func rectangleArea(r Rectangle) float64 {
```

to,

```go
func (r Rectangle) area() float64 {
```

We can just call it area() for both the Circle and Rectangle,
because its only associated with its receiver type.

Hence,

```go
func (c Circle) area() float64 {
```

You can call this special function (i.e. Method) using the `.` operator,

From,

```go
rectangleArea(rect1)
```

to,

```go
rect1.area()
```

Furthermore Methods can use a Pointer Receiver Type or a Value Receiver Type.

Value Receiver,

```go
func (r Rectangle) area() float64 {
```

Pointer Receiver,

```go
func (r *Rectangle) areaPtr() float64 {
```

Go knows how to sort this out when its called. I don't agree.
I feel it should be explicit as follows,

```go
(&rect1).area()
```

## INTERFACES

Note we were able to name the Rectangle's area() method the same thing
as the Circle's area() method.

Go has a way of making these accidental similarities explicit
through a type known as an Interface.

Interfaces are named collections of method signatures or method sets.

So lets make a Shape interface that uses the area() method,

```go
type Shape interface {
  area() float64
}
```

By itself this is not useful, but we can use interface types
as arguments to functions,

```go
func theArea(s Shape) float64 {
  return s.area()
}
```

And now we just call one fuction for getting the area.  As with methods,
the function will figure out the receiver type and call the appropriate
function.

```go
theArea(&circ1)
theArea(&rect1)
```
