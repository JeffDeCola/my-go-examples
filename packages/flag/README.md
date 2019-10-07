# flag (standard) example

_The 'flag` package makes it easy to implement command-line flag parsing._

Refer to the
[flag](https://golang.org/pkg/flag/)
package for more info.

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## RUN

Define flags using,

```go
// STRING
stringPtr := flag.String("s", "default", "This is the flag for a string")
// INTEGER
integerPtr := flag.Int("i", 1, "This is the flag for an integer")
// BOOLEAN
boolPtr := flag.Bool("b", false, "This is the flag for a boolean")
```

The following formats are permitted,

* -flagname
* -flagname=x
* -flagname x

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
