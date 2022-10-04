package webserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var PORT = ":8080"

func Start() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/employees", getEmployeesHandler)
	http.HandleFunc("/createEmployee", createEmployeesHandler)

	log.Println("Server is running at PORT", PORT)
	http.ListenAndServe(PORT, nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	msg := "Hello World"
	fmt.Fprint(w, msg)
}

func getEmployeesHandler(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		tpl, err := template.ParseFiles("./webserver/template.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tpl.Execute(w, map[string]interface{}{
			"data":  getEmployees(),
			"title": "Halaman Employee",
			"error": nil,
		})
	}

	// http.Error(w, "Invalid method", http.StatusBadRequest)
}

func createEmployeesHandler(w http.ResponseWriter, r *http.Request) {

	var emp Employee

	if r.Method == "POST" {
		json.NewDecoder(r.Body).Decode(&emp)

		createEmployees(emp)

		json.NewEncoder(w).Encode(map[string]interface{}{
			"data":  getEmployees(),
			"error": nil,
		})
		return
	}
}
