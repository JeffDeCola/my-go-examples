# markdown-create-table-of-contents example

`markdown-create-table-of-contents` _is a example of
parsing threw a markdown file to find the ## and create a table
of contents for links of github.  I used it at the beginning of
my readme files._

tl;dr,

```bash
markdown-create-table-of-contents -i input.md
markdown-create-table-of-contents -i input.md -h3
```

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## RUN

```bash
go run markdown-create-table-of-contents.go -i input.md
```

You can also include headings 3 `###`,

```bash
go run markdown-create-table-of-contents.go -i input.md -h3
```

To install (place an executable in your $GOPATH/bin),

```bash
go install markdown-create-table-of-contents.go
```
