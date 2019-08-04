# interface example

`interface` _is an example of
accepting ANYTHING (as long as anything has a method attached
to an interface)._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## INTERFACE

The code is actually quite easy.

1; Just define the interface. Attach methods to it.

```go
type theInterface interface {
    doThis(n int)
}
```

1; Create a function that has the interface as a parameter.

Therefor anything can be sent to it as long as it has
a method doThis().

```go
func magic(i theInterface, n int) {
    i.doThis(n)
}
```

1; Call the function

```go
var a1 = who cares as long as can use method doThis()
magic(a1, 44)
```
