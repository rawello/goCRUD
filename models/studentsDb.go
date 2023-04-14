package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"fmt"
)

const dbuser = "root"
const dbpass = "password"
const dbname = "students"

func GetStudents() []Student {

	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM student")

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	students := []Student{}

	for results.Next() {
		var stud Student

		err = results.Scan(&stud.Id, &stud.Name, &stud.Age, &stud.Country, &stud.City)

		if err != nil {
			panic(err.Error())
		}

		students = append(students, stud)
	}
	return students
}

func GetStudent(id string) *Student {

	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)
	stud := &Student{}

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM student where id=?", id)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	if results.Next() {
		err = results.Scan(&stud.Id, &stud.Name, &stud.Age, &stud.Country, &stud.City)

		if err != nil {
			return nil
		}
	} else {
		return nil
	}

	return stud

}

func AddStudent(name string, age int, country string, city string) {
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query(
		"INSERT INTO student (name, age, country, city) VALUES (?,?,?,?)",
		name, age, country, city)

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
}

func DeleteStudent(id int) {
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)

	if err != nil {
		panic(err.Error())
	}

	delete, err := db.Query(
		"DELETE FROM student where  id=?", id)

	if err != nil {
		panic(err.Error())
	}

	defer delete.Close()
}

func AutorizeStudent(name string) {
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)

	if err != nil {
		panic(err.Error())
	}
	autorize, err := db.Query("SELECT name FROM student WHERE Name=?", name)

	if err != nil {
		panic(err.Error())
	}

	defer autorize.Close()
}
