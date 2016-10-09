
# gomock example

`gomock` _is a example of using gomock on an interface
for unit testing._

## RUN TEST

```bash
go test -v -cover ./...
```

## MOCKGEN

`mockgen` file was created by running:

```bash
mockgen -source=laboratory.go -package=laboratory -destination=mockcreature.go
```

`laboratory_test.go` file was created by running:

```bash
gotests -w -all laboratory.go
```