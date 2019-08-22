# postgreSQL example

`postgreSQL` _is an example of
a relational database.
Read/write from/to a table._

Documentation and reference,

* My
  [postgreSQL cheat sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/database/postgresql-cheat-sheet)
  on how to get postgreSQL running on your machine.

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## CREATE A POSTGRESQL DATABASE AND A TABLE (SCHEMA)

Obviously you need postgreSQL server running. Check,

```go
ps aux | grep -i postgres
```

Create a database `jeff_db_example` and a table `people` using `psql`,

```bash
sudo -u postgres psql
CREATE USER jeffd WITH ENCRYPTED PASSWORD 'mypass';
CREATE DATABASE jeff_db_example OWNER jeffd;
\c jeff_db_example
CREATE TABLE people (id int primary key not null, first_name text, last_name text);
GRANT ALL PRIVILEGES ON TABLE people TO jeffd;
```

## CHECK TABLE

Check your table using psql,

```bash
sudo -u postgres psql
\c jeff_db_example
\d
\d people
\d+ people
select * from people;
select last_name from people;
```

## RUN

First, make sure you have the go library

```bash
go get -u github.com/lib/pq
```

Run,

```bash
go run postgreSQL.go
```

## RUN WILL DO THE FOLLOWING

Lets dive into the code.

### CREATE A NEW ROW (id=3)

```go
_, err = db.Exec(`
    insert into people (id, first_name, last_name)
    values (3, 'Jeff', 'DeCola')
`)
```

### CREATE A NEW ROW (id=4)

```go
fmt.Printf("CREATE A NEW ROW (id=4)\n")
_, err = db.Exec(`
        insert into people (id, first_name, last_name)
        values (4, 'John', 'Henry')
    `)
checkErr(err)
```

### UPDATE A COLUMN IN A ROW (id=4)

```go
_, err = db.Exec(`
    update people set first_name = 'fred' where id = 4
`)
```

### READ A COLUMN (last_name) FROM ROW (id=3)

```go
var lastname string
id := 3
err = db.QueryRow(`
    select last_name from people
    where id = $1
`, id).Scan(&lastname)
fmt.Printf("last_name is %s\n", lastname)
```

### READ AN ENTIRE COLUMN (last_name) FROM ALL ROWS

```go
rows, err := db.Query(`select last_name from people`)
defer rows.Close()
lastnames := []string{}
for rows.Next() {
    var lastname string
    err = rows.Scan(&lastname)
    lastnames = append(lastnames, lastname)
}
fmt.Printf("lastnames are %s\n", lastnames)
```

### READ AN ENTIRE ROW (id=3)

```go
var theid int32
var firstName, lastName string
err = db.QueryRow(`
    select * from people where id = 3
    `).Scan(&theid, &firstName, &lastName)
fmt.Printf("row %d is: %s, %s\n", theid, firstName, lastName)
```

## YOUR OUTPUT SHOULD LOOK LIKE

```bash
Successfully connected to db!
CREATE A NEW ROW (id=3)
CREATE A NEW ROW (id=4)
UPDATE A COLUMN IN A ROW (id=4)
READ A COLUMN (last_name) FROM ROW (id=3)
    last_name is DeCola
READ AN ENTIRE COLUMN (last_name) FROM ALL ROWS
    lastnames are [DeCola Henry]
READ AN ENTIRE ROW (id=3)
    row 3 is: Jeff, DeCola
```
