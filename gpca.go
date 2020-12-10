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

func handleRequest() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", home) // TODO: change these to a class with functions
    router.HandleFunc("/verify", verify)
    log.Fatal(http.ListenAndServe(":10000", router))
}

func main(){
    fmt.Println("GPCA")
    handleRequest()
}
