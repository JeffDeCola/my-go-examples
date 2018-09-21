# slices-underlying-array example

`slices-underlying-array` _is an example of a slice being appended to, and shows
the underlying array being added to the slice_.

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## LENGTH AND CAPACITY OF A SLICE

The len of a slice is the current length.

Since slices are build on arrays, the capacity shows what the underlying fixed length array is being used.

This example demonstrates that appending to a slice and showing how the underlying array changes.

## MAKE (LEN & CAP) AND APPEND

When you make and array,

```go
s := make([]int, 0, 5)
```

The length of the slice is 0 and the capacity is 5, so when using
append his will start filling the array at index 0.

## MAKE A SLICE WITH AN ARRAY

Since slices have an underlying array, we can make a slice from an array,

```go
s := make([]int, 10, 18)
```

same as,

```go
s := new([18]int)[0:10]
```

or,

```go
l := new([18]int)
g := l[0:10]
```
