package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Hello Sublime Text World!")

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/users", HomeHandler)
	r.HandleFunc("/articles", HomeHandler)
	http.Handle("/", r)
	http.ListenAndServe(":9000", r)
}

func HomeHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Cache-Control", "no-cache")
	rw.WriteHeader(http.StatusOK)

	reply := make(map[string]interface{})
	reply["ver"] = 1.0

	myobj := make(map[string]interface{})
	myobj["id"] = 2
	myobj["name"] = "PenguenUmut"

	reply["who"] = myobj

	myarr := []string{"a", "bb", "ccc"}
	reply["arr"] = myarr
	json.NewEncoder(rw).Encode(reply)
}
