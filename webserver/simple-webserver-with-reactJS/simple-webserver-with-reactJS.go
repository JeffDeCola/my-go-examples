// my-go-examples simple-webserver-reactjs.go

package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const (
	ip   = "127.0.0.1"
	port = "1234"
)

const htmlMonkeyIndex = `
<html>
    <body>
        <p> You are in the monkeyHandler </p>
        <a href="/jeff">Goto /jeff</a>
    </body>
</html>
`

// rootHandler
func rootHandler(w http.ResponseWriter, r *http.Request) {
	// print out the req struct for fun
	fmt.Printf("Request is %+v\n\n", r.URL)

	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)

	//io.WriteString(res, "You are in the rootHandler\n")
	//io.WriteString(res, "hello, world!\n")
	fmt.Println("done index")

}

// jeffHandler
func jeffHandler(w http.ResponseWriter, r *http.Request) {
	// print out the req struct for fun
	fmt.Printf("Request is %+v\n\n", r.URL)
	//io.WriteString(w, "You are in the jeffHandler\n")
	//io.WriteString(w, "hello, Jeff!\n")

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := `{"hello":"hi"}`
	json.NewEncoder(w).Encode(payload)
	fmt.Println("done jeff")

}

// monkeyHandler
func monkeyHandler(w http.ResponseWriter, r *http.Request) {
	// print out the req struct for fun
	fmt.Printf("Request is %+v\n\n", r.URL)
	fmt.Fprintf(w, htmlMonkeyIndex)
	fmt.Println("done monkey")

}

func startServer() {
	// LISTEN ON IP AND PORT
	fmt.Printf("\nListening on %s:%s\n\n", ip, port)
	log.Fatal(http.ListenAndServe(ip+":"+port, nil))
}

func main() {

	// HANDLER
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/jeff", jeffHandler)
	http.HandleFunc("/monkey", monkeyHandler)

	// START WEBSERVER
	go startServer()

	// PRESS RETURN TO EXIT
	fmt.Scanln()
	fmt.Println("done")

}
