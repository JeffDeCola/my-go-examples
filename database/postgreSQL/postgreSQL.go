package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "jeffd"
	password = "mypass"
	dbname   = "jeff_db_example"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d sslmode=disable dbname=%s user=%s password=%s",
		host, port, dbname, user, password)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("ERROR - Could not open sql")
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("ERROR - No ping")
		panic(err)
	}

	fmt.Println("Successfully connected to db!")

	// CREATE A NEW ROW (id=3)
	fmt.Printf("CREATE A NEW ROW (id=3)\n")
	_, err = db.Exec(`
			insert into people (id, first_name, last_name)
			values (3, 'Jeff', 'DeCola')
		`)
	if err != nil {
		fmt.Println("ERROR - Could not write")
		panic(err)
	}

	// CREATE A NEW ROW (id=4)
	fmt.Printf("CREATE A NEW ROW (id=4)\n")
	_, err = db.Exec(`
			insert into people (id, first_name, last_name)
			values (4, 'John', 'Henry')
		`)
	if err != nil {
		fmt.Println("ERROR - Could not write")
		panic(err)
	}

	// UPDATE A COLUMN IN A ROW (id=4)
	fmt.Printf("UPDATE A COLUMN IN A ROW (id=4)\n")
	updateResult, err := db.Exec(`
			update people set first_name = 'larry' where id = 4
		`)
	if err != nil {
		fmt.Println("ERROR - Could not write")
		panic(err)
	}
	rowsresultResult, err := updateResult.RowsAffected()
	if err != nil {
		fmt.Printf("ERROR - No Rows Updated %d\n", rowsresultResult)
		panic(err)
	}
	if rowsresultResult == 0 {
		fmt.Printf("No Rows Updated %d\n", rowsresultResult)
		err = errors.New("ERROR - Guess What - No Rows Updated")
		panic(err)
	}

	// READ A COLUMN (last_name) FROM ROW (id=3)
	fmt.Printf("READ A COLUMN (last_name) FROM ROW (id=3)\n")
	var lastname string
	id := 3
	err = db.QueryRow(`
		select last_name from people
		where id = $1
	`, id).Scan(&lastname)
	if err != nil {
		fmt.Println("ERROR - Could not write")
		panic(err)
	}
	fmt.Printf("    last_name is %s\n", lastname)

	// READ AN ENTIRE COLUMN (last_name) FROM ALL ROWS
	fmt.Printf("READ AN ENTIRE COLUMN (last_name) FROM ALL ROWS\n")
	rows, err := db.Query(`select last_name from people`)
	if err != nil {
		fmt.Println("ERROR - Could not write")
		panic(err)
	}
	defer rows.Close()
	lastnames := []string{}
	for rows.Next() {
		var lastname string
		err = rows.Scan(&lastname)
		if err != nil {
			fmt.Println("ERROR - Could not scan")
			panic(err)
		}
		lastnames = append(lastnames, lastname)
	}
	fmt.Printf("    lastnames are %s\n", lastnames)

	// READ AN ENTIRE ROW (id=3)
	fmt.Printf("READ AN ENTIRE ROW (id=3)\n")
	var theid int32
	var firstName, lastName string
	err = db.QueryRow(`
		select * from people where id = 3
		`).Scan(&theid, &firstName, &lastName)
	if err != nil {
		fmt.Println("ERROR - Could not read")
		panic(err)
	}
	fmt.Printf("    row %d is: %s, %s\n", theid, firstName, lastName)

}
