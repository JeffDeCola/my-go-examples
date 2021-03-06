# interface example

_A simple example of an interface as parameter and an interface as a return.
Also see the next example that uses this in a package._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## THE GOAL

The goal is to have an interface that has can have
different data (e.g. a, b), where you can then use that interface
as a argument in a function and do something,

So if I make an interface with some data,

a := makemyStructA("jeff")
b := makemyStructB(222, 333)

I can use it in the same function to do something,

```go
magic(a)
magic(b)
```

## AN ILLUSTRATION ON HOW WE ACCOMPLISH THIS

This diagram illustrates what we are doing,

![IMAGE - pointers1 - IMAGE](../../../docs/pics/basic-syntax/interface-example.jpg)

## INTERFACE AS A PARAMETER

The code is actually quite easy.

**Step 1:** Create your data types,

```go
type myStructA struct {
    name string
}

type myStructB struct {
    x int
    y int
}
```

**Step 2:** Create methods with same name using the 3 data types as receivers,

```go
func (i myStructA) doThis() {
    fmt.Printf("I'm in doThis() method with receiver myStructA - %v\n", i.name)
}

func (i myStructB) doThis() {
    fmt.Printf("I'm in doThis() method with receiver myStructB - %v %v\n", i.x, i.y)
}
```

**Step 3:** Now create your interface type that will accept any receiver that has
the method name `area()`.  Simple!

```go
type myInterfacer interface {
    doThis()
}
```

**Step 4:** Create a function that uses this interface as a parameter,

```go
// INTERFACE AS A FUNCTION PARAMETER
func magic(i myInterfacer) {
    i.doThis()
}
```

The interface figures out what method to use based on data type.
Its really cool.

## INTERFACE AS A RETURN

It does that same thing as parameter but this time you make an
interface before you pass to magic().

So instead of declaring your data type,

```go
// Declare and assign the struct pointers
var a = myStructA{"jeff"}
var b = myStructB{222, 333}
```

You will create your interface type,

```go
// INTERFACE AS A RETURN
// Get the interface via a function
a := makemyStructA("jeff")
b := makemyStructB(222, 333)
```

Where,

```go
// INTERFACE AS A RETURN
func makemyStructA(name string) myInterfacer {
    return myStructA{name}
}

// INTERFACE AS A RETURN
func makemyStructB(x, y int) myInterfacer {
    return myStructB{x, y}
}
```

## RUN

```bash
go run interface-as-a-parameter.go
go run interface-as-a-return.go
```
