# pointers example

`pointers` _in an example of a pointer to a struct._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## POINTER

Passing the address of a struct to a function,

```go
b := person{"Larry", "male", 25}
changeNamePtr(&b)
```

So the function can now work on that original struct,

```go
func changeNamePtr(p *person) {
    p.name = "Fred"
}
```

In this case the name is changed from Larry to Fred.
