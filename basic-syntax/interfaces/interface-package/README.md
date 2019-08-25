# interface-package example

`interface-package` _is an example of
an interface placed in a package `example`._

This example is based on the previous example
[interface](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/interface).

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## THE BOTTOM LINE

This is exactly the same as the previous example except I am using a package.
This makes the beauty of the interface come out, since we just make
the interface and use the interface. Everything is hidden in the
package except for the following,

* type **MyInterfacer** interface {
* func **MakemyStructA**(name string) MyInterfacer {
* func **MakemyStructB**(x, y int) MyInterfacer {
* func **Magic**(i MyInterfacer) {

As you can see it makes things easy.

I did add one thing, the ability to update the data for StructA.

* func **UpdateName**(i MyInterfacer, name string) {

## RUN

```bash
go run interface-package.go
```
