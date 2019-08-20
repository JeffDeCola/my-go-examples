# closure example

`closure` _is an example of
closure (a function that references variables from outside
its body)._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## CLOSURE

A closure is a function value that references variables from outside its body.

Helps us limit the scope of variables used by multiple functions.
Without closure, for two or more functions to have access to the same variable,
that variable would need to be at the package scope.

There are two methods as follows.

### ASSIGN ANONYMOUS FUNCTION (func LITERAL) TO A VARIABLE

Here is an example of closure, assigning an anonymous function
(also called func literal) to a variable `increment1`.

This is called closure because the function captures (enclose)
the surrounding environment and can use it.

### RETURN A FUNCTION TO A FUNCTION

Same program as above, but with function outside main.

But this acts like a variable where once `increment`
is declared and assigned, x is set in that scope.

Just think/treat the function as a variable and closure makes sense.
People try to make this too complicated, like I did above but its just
treating the function `inc(x)` as a variable `increment`.  Simple.

## RUN

```bash
go run anonymous-function.go
go run return-function-to-function.go
```
