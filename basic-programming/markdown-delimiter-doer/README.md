# markdown-delimiter-doer example

`markdown-delimiter-doer` _is a example of
taking a markdown file and "do whatever you want" between the delimiters
and output new markdown file._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## RUN

Run using delimiters `$$` and the `-htmltable` switch,

```bash
go run markdown-delimter-doer.go -htmltable -delimeter \$\$ -i input.md -o output.md
```

To install (place an executable in your $GOPATH/bin),

```bash
go install markdown-delimter-doer.go
```

## SWITCHES

You can make many different switches to do different things.

### HTML TABLE SWITCH

Here is an illustration using the `-htmltable` switch,

![IMAGE - markdown-delimiter-doer - IMAGE](../../docs/pics/markdown-delimiter-doer.jpg)

It will even check the dates and subscript them automatically.
