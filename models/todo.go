package models

import (
	"encoding/json"
	"fmt"
	"log"
   "database/sql"
)

type ToDo struct{
	Id 	  int 	 `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	Done  bool	 `json:"done"`
}

func JsonToTodoStruct (todoPram []byte) ToDo {
	var  ToDo ToDo
	json.Unmarshal(todoPram, &ToDo)
	ToDo.Done = false
	fmt.Printf("(%%+v) %+v\n", ToDo)
	return ToDo
}

func (td *ToDo) Update() error {
	query := fmt.Sprintf(`
		UPDATE todos SET title = ?, body = ?, done = ? WHERE id = ?;
	`)
	_, err := Dbconecction.Exec(query, td.Title, td.Body, td.Done, td.Id)
	if err != nil{
		log.Println(err)
	}
	return err
}

func (td *ToDo) Delete() error {
	query := fmt.Sprintf(`
		DELETE FROM todos WHERE id = ?
	`)
	_, err := Dbconecction.Exec(query, td.Id)
	if err != nil{
		log.Println("Delete")
		log.Println(err)
	}
	return err
}

func GetTodo(id int) (ToDo, error){
	query := fmt.Sprintf(`SELECT * FROM todos WHERE id = ?`)
	row := Dbconecction.QueryRow(query, id)
	var todo ToDo
	err := row.Scan(&todo.Id, &todo.Title, &todo.Body, &todo.Done)
	if err != nil {
		log.Println(err)
	}
	return todo, err
}

func (td *ToDo) Save() error {
	query := fmt.Sprintf(`
		INSERT INTO todos (title, body, done) VALUES (?,?,?)
	`)
	_, err := Dbconecction.Exec(query, td.Title, td.Body, td.Done)
	if err != nil{
		log.Println(err)
	}
	return err
}

func GetAllTodo( done bool ) ([]ToDo) {
	query := fmt.Sprintf(`SELECT * FROM todos WHERE done = ?`)
	rows := Dbconecction.QueryRow(query, done)
	var todos []ToDo
	for rows.Next()  {
		var todo ToDo
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Body, &todo.Done)
		if err != nil {
			log.Println(err)
		}
		todos = append(todos, todo)
	}
	return todos
}