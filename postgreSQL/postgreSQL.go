package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "jeffdecola"
	password = ""
	dbname   = "rm"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d sslmode=disable dbname=%s user=%s password=%s",
		host, port, dbname, user, password)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("could not open sql")
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("no ping")
		panic(err)
	}

	fmt.Println("Successfully connected to db!")

	// WRITE (firstname and lastname) TO TABLE PEOPLE
	/*
		_, err = db.Exec(`
			insert into people (id, first_name, last_name)
			values (3, 'Jeff', 'DeCola')
		`)
		if err != nil {
			fmt.Println("could not write")
			panic(err)
		}

		_, err = db.Exec(`
			insert into people (id, first_name, last_name)
			values (4, 'John', 'Henry')
		`)
		if err != nil {
			fmt.Println("could not write")
			panic(err)
		}
	*/

	// WRITE TO TABLE COLUMN (UPDATE A ROW)
	_, err = db.Exec(`
			update people set first_name = 'fred' where id = 66
		`)
	if err != nil {
		fmt.Println("could not write")
		panic(err)
	}

	// READ (last_name) FROM TABLE PEOPLE with id=3
	var lastname string
	id := 3
	err = db.QueryRow(`
		select last_name from people
		where id = $1
	`, id).Scan(&lastname)
	if err != nil {
		fmt.Println("could not write")
		panic(err)
	}
	fmt.Printf("last_name is %s\n", lastname)

	// READ (last_name) FROM ALL ROWS OF TABLE PEOPLE
	rows, err := db.Query(`select last_name from people`)
	if err != nil {
		fmt.Println("could not write")
		panic(err)
	}
	defer rows.Close()
	lastnames := []string{}
	for rows.Next() {
		var lastname string
		err = rows.Scan(&lastname)
		if err != nil {
			fmt.Println("could scan write")
			panic(err)
		}
		lastnames = append(lastnames, lastname)
	}
	fmt.Printf("lastnames are %s\n", lastnames)

}
