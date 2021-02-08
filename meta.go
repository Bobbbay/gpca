package main

import (
    "fmt"
    "encoding/json"
    "net/http"
)

// In this file, you'll find regular meta functions.
// This ranges from the home function, to the database
// status function, and so on.

func home(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Hit on home") // TODO: don't require to do this on every function
    result := "Welcome to the GPCA."
    json.NewEncoder(w).Encode(result)
}
