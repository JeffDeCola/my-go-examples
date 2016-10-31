# linked-list example

`linked-list` _is a example of a singly linked list (i.e. With just a head pointer)._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## LINKED-LIST

Linked Lists function as an array that can grow and shrink as needed,
from any point in the array.

Each Node containes a data value and a pointer to the next node in the list.
If the pointer is null, you are at the end of the list.

`head` -> `DATA|NEXT` -> `DATA|NEXT` -> `DATA|(null)`

A local head pointer variable points to the first item of a list.

A example of a item is created by a struct:

```go
type Node struct {
    Value
    next  *Node
}
```

Where Value can be something like,

```go
type Value struct {
    name   string
    age    int
    gender string
}
```

### Advantages over arrays

* Nodes can be added or removed from middle of list.
* There is no need to define size.

### Disadvantages over arrays

* You must iterate over the list until you get to the node you want.
* Dynamic pointers required.
* larger overhead.

## RUN

```bash
go run linked-list.go
```

## RUN TEST

```bash
go test -v -cover ./...
```

`linked-list_test.go` file was created by running:

```bash
gotests -w -all linked-list.go
```