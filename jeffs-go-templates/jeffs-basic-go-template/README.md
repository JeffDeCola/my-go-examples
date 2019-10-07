# jeffs-basic-go-template example

_A simple go template with flags, logging & error handling._

Just copy/paste to use in your projects.

Refer to the following go examples for more information,

* [flag](https://github.com/JeffDeCola/my-go-examples/tree/master/packages/flag)
* [logrus](https://github.com/JeffDeCola/my-go-examples/tree/master/packages/logrus)
* [errors](https://github.com/JeffDeCola/my-go-examples/tree/master/packages/errors)

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## RUN

Remember to get the packages,

```bash
go get -u -v github.com/sirupsen/logrus
go get -u -v github.com/pkg/errors
```

The default flag i=42 which will be an error. Hence set the flag i to
any other integer,

```bash
go run jeffs-basic-go-template.go
go run jeffs-basic-go-template.go -i 42
go run jeffs-basic-go-template.go -i 10
```
