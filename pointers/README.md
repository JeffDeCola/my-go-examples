# pointers example

`pointers` _in an example of a pointer to a struct._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## POINTER

Passing the address of a struct to a function.

```go
changeName(&UserID)
```

So the function can now work on that struct.

```go
func changeName(p *person) {
    p.name = "Fred"
}
```

In this case the userID name is changed from Larry to Fred.
