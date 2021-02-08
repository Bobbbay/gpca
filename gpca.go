package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){
    fmt.Println("Welcome to GPCA with love.")

    // Routers for all pages
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", home)
    router.HandleFunc("/verify", verifyCryptocurrency).Methods("POST") // TODO: Throw better error at unknown method
    router.HandleFunc("/new", newPerson).Methods("POST")
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
