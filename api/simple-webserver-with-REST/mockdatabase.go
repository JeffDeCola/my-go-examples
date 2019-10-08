package main

// TodoStruct struct
type TodoStruct struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

// Todos Slice
type Todos []TodoStruct

// MIMIC DATABASE
var todosDatabase = Todos{
	TodoStruct{ID: "10", Name: "Write presentation", Completed: false},
	TodoStruct{ID: "20", Name: "Eat Lunch", Completed: false},
	TodoStruct{ID: "30", Name: "Pick up Milk", Completed: true},
}
