# pointers example

`pointers` _in an example of
a pointer to a struct (pass by value, pass by reference)._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## PASS BY REFERENCE (MAKES POINTER IN FUNCTION)

Passing the address of a struct to a function (by reference),

```go
b := person{"Jill", "female", 27}
changeNamePtr(&b)
```

So the function can now work on that original struct,

```go
func changeNamePtr(p *person) {
    p.name = "Lisa"
}
```

In this case the name is changed from Jill to Lisa.

This diagram may help,

![IMAGE - pointers-pass-by-value-and-reference - IMAGE](../../docs/pics/pointers-pass-by-value-and-reference.jpg)

## RUN

Run,

```go
go run pointers.go
```
