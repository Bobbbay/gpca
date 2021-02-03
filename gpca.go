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

type CryptocurrencyStatus struct {
    Date string `json:"date"`
    Points int `json:"points"`
}

func main(){
    fmt.Println("Welcome to GPCA with love.")

    // Routers for all pages
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", home)
    router.HandleFunc("/verify", verify).Methods("POST") // TODO: Throw better error at unknown method
    router.HandleFunc("/new", newEntry).Methods("POST")
    router.HandleFunc("/status", status)
    router.HandleFunc("/update", update).Methods("POST")
    router.HandleFunc("/status/cryptocurrency", cryptocurrencyStatus).Methods("GET")
    router.HandleFunc("/cryptocurrency/status", cryptocurrencyStatus).Methods("GET")
    router.HandleFunc("/cryptocurrency/claim", claimCryptocurrency).Methods("POST")

    fmt.Println("Established routes: ")
    router.Walk(func(route *mux.Route, r *mux.Router, ancestors []*mux.Route) error {
        t, _ := route.GetPathTemplate()
        fmt.Println("", t)
        return nil
    })

    port := ":10000"
    fmt.Println("GPCA listening on port", port)
    log.Fatal(http.ListenAndServe(port, router))
}
