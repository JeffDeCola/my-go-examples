package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "rm"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("could not open sql")
		panic(err)
	}
	defer db.Close()

	//err = db.Ping()
	//if err != nil {
	//		fmt.Println("no ping")
	//		panic(err)
	//	}

	fmt.Println("Successfully connected to db!")

	// WRITE (firstname and lastname) TO TABLE PEOPLE
	_, err = db.Exec(`
		insert into people (id, first_name, last_name)
		values (1, 'Jeff', 'DeCola')
	`)

	_, err = db.Exec(`
		insert into people (id, first_name, last_name)
		values (2, 'John', 'Henry')
	`)

	// READ (lastname) FROM TABLE PEOPLE with id=2
	var lastname string
	err = db.QueryRow(`
		select lastname from people
		where id = $2
	`, 1).Scan(&lastname)
	// sql.ErrNoRows
	fmt.Printf("lastname is %s\n", lastname)

	// READ ROWS TABLE PEOPLE
	rows, err := db.Query(`select * from people`)
	defer rows.Close()
	lastnames := []string{}
	for rows.Next() {
		var lastname string
		rows.Scan(&lastname)
		lastnames = append(lastnames, lastname)
	}
	fmt.Printf("lastnames are %s\n", lastnames)

}
