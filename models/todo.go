package models

import (
	"encoding/json"
	"fmt"
	"log"
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