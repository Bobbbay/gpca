package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Standard struct {
    Content string `json:"content"`
}

type ReturnCode struct {
    // TODO: refactor for better name
    Response string `json:"response"`
    ErrorCode string `json:"errorcode"`
    Description string `json:"description"`
}

type Person struct {
    // id INTEGER PRIMARY KEY, name TEXT, points INTEGER, cryptocurrency TEXT)
    Id int `json:"id"`
    Name string `json:"name"`
    Points int `json:"points"`
    Cryptocurrency string `json:"cryptocurrency"`
}

func main(){
    fmt.Println("GPCA v0.0.1")

    // Routers for all pages
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", home) // TODO: change these to a class with functions (?)
    router.HandleFunc("/verify", verify).Methods("POST") // TODO: Throw better error at unknown method
    router.HandleFunc("/new", newEntry).Methods("POST")
    router.HandleFunc("/status", status)
    log.Fatal(http.ListenAndServe(":10000", router))
}
