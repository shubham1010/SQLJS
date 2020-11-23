package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Student struct {
	Name   string `json:"name"`
	RollNo int64 `json:"rollno"`
	Course string `json:"course"`
}

func main() {
	SetRouter()
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	data, err := ioutil.ReadFile("./students.json")
	if err != nil {
		fmt.Print(err)
	}

	var students []Student
	err = json.Unmarshal(data, &students)
	if err != nil {
		fmt.Println("error:", err)
	}

	json.NewEncoder(w).Encode(students)
}

func SetRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/api/students", Home).Methods("GET")
	http.ListenAndServe(":8080", router)
	log.Printf("Server is started at 8080 port")
}
