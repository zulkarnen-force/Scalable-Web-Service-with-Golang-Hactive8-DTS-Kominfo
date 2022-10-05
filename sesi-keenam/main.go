package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

const PORT string = ":8080"

type Employee struct {
	ID int
	Name string
	Age int
	Division string
}

var employees = []Employee {
	{ID: 1, Name: "Airlee", Age: 20, Division: "IT"},
	{ID: 2, Name: "Nanda", Age: 21, Division: "Finance"},
	{ID: 3, Name: "Mailo", Age: 23, Division: "IT"},
}

func getEmployees(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(employees)
		return
	}

	http.Error(w, "invalid request", http.StatusBadRequest)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		name := r.FormValue("name")
		age := r.FormValue("age")
		division := r.FormValue("division")

		convertAge, err := strconv.Atoi(age)

		if err != nil {
			http.Error(w, "invalid age", http.StatusBadRequest)
		}

		newEmployee := Employee{
			ID: len(employees) + 1,
			Name: name,
			Age: convertAge,
			Division: division,
		}

		employees = append(employees, newEmployee)

		json.NewEncoder(w).Encode(newEmployee)
		return

	}

	http.Error(w, "invalid method", http.StatusBadRequest)

}


func showEmployee(w http.ResponseWriter, r *http.Request) {
	
	if r.Method == "GET" {
		view, err := template.ParseFiles("./index.html")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		view.Execute(w, employees)
		return
	}

	http.Error(w, "invalid method", http.StatusBadRequest)
}

func main() {

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {
		msg := "Hello World"
		fmt.Fprintln(w, msg)
	})

	http.HandleFunc("/employees", getEmployees)
	http.HandleFunc("/employees/create", createEmployee)
	http.HandleFunc("/employees/view", showEmployee)

	http.ListenAndServe(PORT, nil)
 
}