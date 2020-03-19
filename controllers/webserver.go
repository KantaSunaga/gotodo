package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gotodo/models"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strconv"
)

type Ping struct {
	Status 	int  	`json:"status"`
	Result  string  `json:"result"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	todos := models.GetAllTodo(false)
	undones := models.GetAllTodo(true)
	AllTodos := map[string][]models.ToDo
	AllTodos["todos"] = todos
	AllTodos["undones"] = undones
	w.Header().Set("Content-type", "application/json")
	respBody, _ := json.Marshal(AllTodos)
	w.Write(respBody)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	todo := models.JsonToTodoStruct(body)
	err := todo.Update()
	if err != nil {
		http.Error(w, "Invalid access!!", http.StatusInternalServerError)
	} else {
		ReturnStatusOk(w)
	}
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	todo := models.JsonToTodoStruct(body)
	err := todo.Save()
	if err != nil {
		http.Error(w, "Invalid access!!", http.StatusInternalServerError)
	} else {
		ReturnStatusOk(w)
	}
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	log.Println(id)
	todo, err := models.GetTodo( id )
	if err != nil {
		log.Println("todo not found")
		http.Error(w, "todo not found", http.StatusInternalServerError)
	} else {
		todo.Delete()
		ReturnStatusOk(w)
	}
}

func ReturnStatusOk(w http.ResponseWriter) {
	ping := Ping{http.StatusOK, "ok" }
	res, _ := json.Marshal(ping)
	w.Header().Set("Content-type", "application/json")
	w.Write(res)
}


func StartWebServer() {
	router := mux.NewRouter()
	router.HandleFunc("/index/", indexHandler).Methods("GET")
	router.HandleFunc("/create/", createHandler).Methods("POST")
	router.HandleFunc("/update/", updateHandler).Methods("POST")
	router.HandleFunc("/delete/{id:[0-9]+}/", deleteHandler).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}

