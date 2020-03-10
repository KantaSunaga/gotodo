package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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
	length, _ := strconv.Atoi(r.Header.Get("Content-Length"))
	body := make([]byte, length)
	length, _ = r.Body.Read(body)
	var jsonBody map[string]interface{}
	json.Unmarshal(body[:length], &jsonBody)
	fmt.Printf("%v\n", jsonBody)

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

