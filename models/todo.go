package models

import (
	"encoding/json"
	"fmt"
	"log"
)

type ToDo struct{
	id 	  int 	 `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	Done  bool	 `json:"done"`
}

func NewTodo (todoPram []byte) ToDo {
	var  ToDo ToDo
	json.Unmarshal(todoPram, &ToDo)
	ToDo.Done = false
	fmt.Printf("(%%+v) %+v\n", ToDo)
	return ToDo
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