# slices-underlying-array example

_A slice being appended to showing the underlying array being added to the slice._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## LENGTH AND CAPACITY OF A SLICE

Slices are build on arrays,

* The `length` of the slice is just the length of the slice.
* The `capacity` shows what the underlying fixed length array is being used.

This example demonstrates this by appending to a slice and showing how the
underlying array changes (the capacity).

## RUN

```bash
go run slices-underlying-array.go
```
