package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

// rootHandler
func rootHandler(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	io.WriteString(res, "You are in the rootHandler\n\n")

	// MAKE JSON PRETTY
	js, _ := json.MarshalIndent(todosDatabase, "", "    ")
	res.Write(js)

}

// GET
func getdataHandler(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	// SEARCH DATA FOR todoID
	for _, item := range todosDatabase {
		if item.ID == params["todoID"] {
			// RETURN ITEM
			// json.NewEncoder(res).Encode(item)
			js, _ := json.MarshalIndent(item, "", "    ")
			res.Write(js)
			return
		}
	}

	// RETURN NOT FOUND
	json.NewEncoder(res).Encode("No Data Found for that id")

}

// POST (add)
func postdataHandler(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	// CREATE AN STRUCT TO ADD
	var newTodo TodoStruct

	// GET THE POST DATA
	_ = json.NewDecoder(req.Body).Decode(&newTodo)

	// PUT IN THE ID
	newTodo.ID = params["todoID"]

	// APPEND TO DATABASE
	todosDatabase = append(todosDatabase, newTodo)

	// RETURN DATABASE
	//json.NewEncoder(res).Encode(todosDatabase)
	js, _ := json.MarshalIndent(todosDatabase, "", "    ")
	res.Write(js)
}

// PUT (replace/update)
func putdataHandler(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	// CREATE AN STRUCT TO REPLACE
	var replaceTodo TodoStruct

	// GET THE POST DATA
	_ = json.NewDecoder(req.Body).Decode(&replaceTodo)

	// PUT IN THE ID
	replaceTodo.ID = params["todoID"]

	// SEARCH DATA FOR todoID
	for i, item := range todosDatabase {
		if item.ID == params["todoID"] {
			// REMOVE FROM DATABASE
			todosDatabase = append(todosDatabase[:i], todosDatabase[i+1:]...)
			// REPLACE DATABASE
			todosDatabase = append(todosDatabase, replaceTodo)
			// RETURN DATABASE
			//json.NewEncoder(res).Encode(todosDatabase)
			js, _ := json.MarshalIndent(todosDatabase, "", "    ")
			res.Write(js)
			return
		}
	}

	// RETURN NOT FOUND
	json.NewEncoder(res).Encode("No Data Found for that id")

}

// DELETE
func deletedataHandler(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	// SEARCH DATA FOR todoID
	for i, item := range todosDatabase {
		if item.ID == params["todoID"] {
			// REMOVE TO DATABASE
			todosDatabase = append(todosDatabase[:i], todosDatabase[i+1:]...)
			break
		}
	}

	// RETURN DATABASE
	//json.NewEncoder(res).Encode(todosDatabase)
	js, _ := json.MarshalIndent(todosDatabase, "", "    ")
	res.Write(js)
}
