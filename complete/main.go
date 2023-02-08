package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

const createTable string = `CREATE TABLE courses(
	department TEXT,
  	code INT,
  	description TEXT
);`

const insertTuple string = `insert into courses (department, code, description) VALUES (?,?,?);`

const getCseCourses string = `select department, code, description 
							  from courses where department = ? 
							  order by code;`

func main() {
	// create .db file
	database, err := sql.Open("sqlite3", "./ucsd.db")
	if err != nil {
		fmt.Println(err.Error())
	}

	// create table in .db file
	statement, err := database.Prepare(createTable)
	if err != nil {
		fmt.Println(err.Error())
	}
	statement.Exec()

	// insert tuples into table 'courses'
	// notice ? in the insert statement.
	statement, _ = database.Prepare(insertTuple)
	statement.Exec("cse", 100, "data structure")
	statement.Exec("cse", 224, "network systems")
	statement.Exec("cse", 221, "operating systems")
	statement.Exec("cse", 232, "database")
	statement.Exec("cse", 202, "algorithms")
	statement.Exec("cse", 240, "computer architecture")
	statement.Exec("ece", 225, "Probability")
	statement.Exec("ece", 260, "VLSI")

	// retrieve all cse courses tuples in ascending order by code.
	rows, err := database.Query(getCseCourses, "cse")
	if err != nil {
		fmt.Println(err.Error())
	}
	var department string
	var code int
	var description string
	for rows.Next() {
		rows.Scan(&department, &code, &description)
		fmt.Println(department + strconv.Itoa(code) + ": " + description)
	}

}
