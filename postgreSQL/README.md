# postgreSQL example

`postgreSQL` _setup and read/write from/to a table._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

A cheat cheet is
[here](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/postgreSQL-cheat-sheet)

## CREATE A TABLE

Make sure you have the go library

```bash
go get -u github.com/lib/pq
```

Create a database `rm` and a table `people`

```bash
psql -d postgres
create database rm
\c rm
create table people (id int primary key not null, first_name text, last_name text);
```




