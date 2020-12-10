package main

import (
    "fmt"
    "net/http"
    "encoding/json"
)

func home(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Hit on home") // TODO: don't require to do this on every function
    result := Standard{Content: "Hello world"}
    json.NewEncoder(w).Encode(result)
}

