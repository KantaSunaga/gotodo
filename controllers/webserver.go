package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gotodo/models"
	"io/ioutil"
	"log"
	"net/http"
)

type Ping struct {
	Status 	int  	`json:"status"`
	Result  string  `json:"result"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	ping := Ping{http.StatusOK, "ok" }
	res, _ := json.Marshal(ping)
	w.Header().Set("Content-type", "application/json")
	w.Write(res)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	todo := models.NewTodo(body)
	err := todo.Save()
	if err != nil {
		log.Println("失敗")
		http.Error(w, "Invalid access!!", http.StatusInternalServerError)
	} else {
		log.Println("成功")
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
	http.ListenAndServe(":8080", router)
}

