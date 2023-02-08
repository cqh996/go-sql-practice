package main

import (
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

	// create table in .db file

	// insert tuples into table 'courses'
	// notice ? in the insert statement.

	// retrieve all cse courses tuples in ascending order by code.
}
