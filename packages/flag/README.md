# flag (standard) example

`flag` _standard package makes it easy to implement command-line flag parsing._

Refer to the
[flag](https://golang.org/pkg/flag/)
package for more info.

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## RUN

Define flags using flag.String(), Bool(), Int(), etc.

The following formats are permitted,

* -flag
* -flag=x
* -flag x  // non-boolean flags only

String,

```bash
go run flag.go -s "happy days"
```

Integer,

```bash
go run flag.go -i 56
```

Boolean,

```bash
go run flag.go -b
```
