package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Todo - The struct for one of the routes/endpoints
type Todo struct {
	ID        int       `json:"id"` // for the mock database
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// Todos - Slice of Todo - Will build this in /jeffHandler
type Todos []Todo

// Person - The struct for one of the routes/endpoints
type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

// Address is the address
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

const htmlIndex = `
<html>
<body>
<p>Welcome to the simple webserver with rest!</p>
<p>This is Index "/" </p>
<p>todos is using a mock database: </p>
<a href="/todos">Goto /todos</a> which will use GET TodosIndex() <br />
<a href="/todos/hi">Goto /todos/hi </a> which will use GET TodosShow() <br />
<br />
<p>POST something to /todos - which will use POST TodosCreate()</p>
<p>curl -X POST -H "Content-Type: application/json" -d '{"name":"Get a Cat"}' http://127.0.0.1:1234/todos </p>
<br />
<p>people is just using a slice to save data: </p> 
<a href="/people">Goto /people </a> which will use GET GetPeopleEndpoint() <br />
<a href="/people/1">Goto /people/1 </a> which will use GET GetPersonEndpoint() <br />
<br />
<p>POST a person to /people/1 or /people/2 - which will use POST CreatePersonEndpoint()</p>
<p>curl -X POST -H "Content-Type: application/json" -d '{"Firstname":"Jeff", "Lastname":"DeCola"}' http://127.0.0.1:1234/people/1 </p>
<p>curl -X POST -H "Content-Type: application/json" -d '{"Firstname":"Brian", "Lastname":"Smith"}' http://127.0.0.1:1234/people/2 </p>
<br />
<p>DELETE a person to /people/1 - which will use DELETE DeletePersonEndpoint()</p>
<p>curl -X DELETE -H "Content-Type: application/json" -d http://127.0.0.1:1234/people/1 </p>
</body>
</html>
`

// Index is "/"
func Index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, htmlIndex)
}

// TodosIndex is /todos
func TodosIndex(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(http.StatusOK)
	// encode json
	if err := json.NewEncoder(res).Encode(todos); err != nil {
		panic(err)
	}
}

// TodosShow is /todos/{ids}
func TodosShow(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	todoID := vars["todoID"]
	fmt.Fprintln(res, "Todo show:", todoID)
}

// TodosCreate is /todos/{ids}
// e.g. curl -X POST -H "Content-Type: application/json" -d '{"name":"Get a Cat"}' http://127.0.0.1:1234/todos
func TodosCreate(res http.ResponseWriter, req *http.Request) {
	var todo Todo

	// Read POST
	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))
	fmt.Printf("    TodosCreate - Body is %s\n", body)

	if err != nil {
		panic(err)
	}

	// Close
	if err := req.Body.Close(); err != nil {
		panic(err)
	}

	// umarshals json into todo struct
	if err := json.Unmarshal(body, &todo); err != nil {
		res.Header().Set("Content-Type", "application/json; charset=UTF-8")
		res.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(res).Encode(err); err != nil {
			panic(err)
		}
	}

	fmt.Printf("    TodosCreate - todo is %v\n", todo)

	// Save this data in mockdatabase
	t := RepoCreateTodo(todo)
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(res).Encode(t); err != nil {
		panic(err)
	}
}

// GetPersonEndpoint is /people
func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

// GetPeopleEndpoint is /people/{id}
func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

// CreatePersonEndpoint is /people/{id}
func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

// DeletePersonEndpoint is /people/{id}
func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}
