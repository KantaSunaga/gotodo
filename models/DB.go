package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Dbconecction *sql.DB

func DbInit() {
	  createTodoTable()
}

func createTodoTable() {
	Dbconecction, _ = sql.Open("sqlite3", "./data.sql")
	query := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS "todos"(
			id 		INTEGER PRIMARY KEY AUTOINCREMENT,
			titile	STRING,
			body	INTEGER,
			done	INTEGER
		);`)
	_, err := Dbconecction.Exec(query)
	if err != nil{
		log.Fatalln(err)
	}
}

