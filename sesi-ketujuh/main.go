package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)


const (
	HOST = "localhost"
	PORT = 5432
	USER = "postgres"
	PASSWORD = "root"
	DBNAME = "dbgosql"
)


var (
	db *sql.DB
	err error
	row *sql.Row
)


type Employee struct {
	ID int
	Full_name string
	Email string
	Age int
	Division string
}


func CreateEmployee() {
	var employee Employee = Employee{}

	var sqlStatement string = `
		INSERT INTO employees (full_name, email, age, division) VALUES ($1, $2, $3, $4) 
		returning *
	`

	err = db.QueryRow(sqlStatement, "Zulkarnen", "zulkarnen@gmail.com", 23, "IT").
		Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

	if err != nil {
		panic(err)
	}

	fmt.Printf("New Employee Data: %+v", employee)
}


func GetEmployees()  {
	var results = []Employee{}

	var sqlStatement string = `SELECT * FROM employees`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var employees = Employee{}

		err := rows.Scan(&employees.ID, &employees.Full_name, &employees.Email, &employees.Age, &employees.Division)
		
		if err != nil {
			panic(err)
		}
		
		results = append(results, employees)
	}

	fmt.Println("Employee datas: ", results)
}



func UpdateEmployee() {
	var sqlStatement = `
		UPDATE employees
			SET full_name = $2, email = $3, division = $4, age = $5
		WHERE id = $1
		` 
	
		res, err := db.Exec(sqlStatement, 1, "Abdul Rosad", "abdul@gmail.com", "Finance", 50)

		if err != nil {
			panic(err)
		}

		count, err := res.RowsAffected()

		if err != nil {
			panic(err)
		}

		fmt.Println("updated data amount: ", count)
}


func DeleteEmployee() {
	var sqlStatement = `
		DELETE FROM employees WHERE id = $1
		` 
	
		res, err := db.Exec(sqlStatement, 1)

		if err != nil {
			panic(err)
		}

		count, err := res.RowsAffected()

		if err != nil {
			panic(err)
		}

		fmt.Println("Deleted data amount: ", count)
}


func main()  {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s " +
		"dbname=%s sslmode=disable", HOST, PORT, USER, PASSWORD, DBNAME)

	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err)
	}



	fmt.Println("successfully connected to database")

	// CreateEmployee()
	// GetEmployees()
	UpdateEmployee()
	GetEmployees()
	fmt.Println("Deleting...")
	DeleteEmployee()
	GetEmployees()
}