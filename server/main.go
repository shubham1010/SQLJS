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
	RollNo int64  `json:"rollno"`
	Course string `json:"course"`
}

type Message struct {
	Content  string `json:"content"`
	Response bool   `json:"response"`
}

func main() {
	SetRouter()
}

func GetStudentsInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get request is made..")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	data, err := ioutil.ReadFile("./students.json")
	if err != nil {
		log.Fatal(err)
	}

	var students []Student
	err = json.Unmarshal(data, &students)
	if err != nil {
		log.Fatal("error:", err)
	}

	json.NewEncoder(w).Encode(students)
}

func StoredStudentInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Post request is made...")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	body, err := ioutil.ReadAll(r.Body)

	var students []Student
	var msg Message

	msg.Response = false
	err = json.Unmarshal([]byte(body), &students)
	if err != nil {
		fmt.Println("JSON Decoding Not Happened")
		w.Header().Set("Content-Type", "application/json")
		msg.Content = "JSON Decoding Not Happened at server side"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(msg)
		return
	}

	file, errIndent := json.MarshalIndent(students, "", "  ")
	if errIndent != nil {
		fmt.Println("Indentendation is not happend")
		w.Header().Set("Content-Type", "application/json")
		msg.Content = "Indentendation is not happend"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(msg)
		return
	}
	err = ioutil.WriteFile("students.json", file, 0644)
	if err != nil {
		fmt.Println("File is not updated..")
		w.Header().Set("Content-Type", "application/json")
		msg.Content = "File is not updated"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(msg)
		return
	}

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	msg.Content = "File is successfully updated..."
	msg.Response = true
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(msg)
}

func SetRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/api/students", GetStudentsInfo).Methods("GET")
	router.HandleFunc("/api/students", StoredStudentInfo).Methods("POST")
	http.ListenAndServe(":8080", router)
	log.Printf("Server is started at 8080 port")
}
