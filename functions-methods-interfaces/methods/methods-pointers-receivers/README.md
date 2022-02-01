# methods-pointers-receivers

_Using methods to calculate the area of a rectangle and circle
by passing pointers and using pointer receivers._

Other examples using,

* **FUNCTIONS**
  * [functions](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/functions/functions)
  * [functions-pointers-arguments](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/functions/functions-pointers-arguments)
* **METHODS**
  * [methods](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/methods/methods)
  * [methods-pointers-arguments](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/methods/methods-pointers-arguments)
  * [methods-pointers-receivers](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/methods/methods-pointers-receivers)
    **<- YOU ARE HERE**
* **INTERFACES**
  * [interfaces](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/interfaces/interfaces)
  * [interfaces-pointers-arguments](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/interfaces/interfaces-pointers-arguments)
  * [interfaces-pointers-receivers](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/interfaces/interfaces-pointers-receivers)
  * [geometry-package](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/interfaces/geometry-package)
  
tl;dr,

```go

```

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/methods/area-shapes-methods-ptrs#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/methods/area-shapes-methods-ptrs#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/methods/area-shapes-methods-ptrs#test)
* [AN ILLUSTRATION THAT MAY HELP](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/methods/area-shapes-methods-ptrs#an-illustration-that-may-help)

Documentation and reference,

* This repos [github webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

A method is a function with a special receiver type.

It just makes it a lot easier to use a
struct as a receiver, rather than passing lots of variables around.

```go
type Rectangle struct {
    width  float64
    height float64
}

rec := Rectangle{2.4, 34.4}
```

Calculate the area using a method,

```go
rec.area(&recArea)
```

Which will chose the area method that has a Rectangle struct,

```go
func (r Rectangle) area(a *float64) {
    *a = r.width * r.height
}
```

## RUN

To run,

```bash
go run methods-pointers-receivers.go
```

## TEST

To create _test files,

```bash
gotests -w -all methods-pointers-receivers.go
```

To unit test the code,

```bash
go test -cover ./... 
```

## AN ILLUSTRATION THAT MAY HELP

![IMAGE - methods-interfaces-pointers-receivers.jpg - IMAGE](../../../docs/pics/basic-syntax/methods-interfaces-pointers-receivers.jpg)
