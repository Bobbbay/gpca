package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequest() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", home) // TODO: change these to a class with functions
    router.HandleFunc("/verify", verify)
    log.Fatal(http.ListenAndServe(":10000", router))
}

func home(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Hit on home") // TODO: don't require to do this on every function
    json.NewEncoder(w).Encode("{'Hello world': 'baz'}")
}

func verify(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Hit on verify")
    json.NewEncoder(w).Encode("{'Hello world':'Foo Baz'}")
}

func main(){
    fmt.Println("GPCA")
    handleRequest()
}
