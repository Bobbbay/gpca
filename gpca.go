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

func main(){
    fmt.Println("GPCA v0.0.1")

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", home) // TODO: change these to a class with functions
    router.HandleFunc("/verify", verify).Methods("POST") // TODO: Throw better error at unknown method
    log.Fatal(http.ListenAndServe(":10000", router))
}
