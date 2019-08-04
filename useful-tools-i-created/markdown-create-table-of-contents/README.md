# markdown-create-table-of-contents example

`markdown-create-table-of-contents` _is a example of
parsing a markdown file to find ##, ### to create a table
of contents (TOC) for links at github.  I like to use a
TOC it at the beginning of my readme files._

tl;dr,

```bash
markdown-create-table-of-contents
markdown-create-table-of-contents -i input.md
markdown-create-table-of-contents -i input.md -h3
```

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## RUN & INSTALL

Will always assume README.md

```bash
markdown-create-table-of-contents
```

If if not README, it will add that filename `input` to the link,

```bash
go run markdown-create-table-of-contents.go -i input.md
```

You can also include headings 3 switch `-h3`,

```bash
go run markdown-create-table-of-contents.go -h3
go run markdown-create-table-of-contents.go -i input.md -h3
```

To install (place an executable in your $GOPATH/bin),

```bash
go install markdown-create-table-of-contents.go
```
