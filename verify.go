package main

import (
    "fmt"
    "encoding/json"
    "net/http"
)
func verify(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Hit on verify")
    json.NewEncoder(w).Encode("{'Hello world':'Foo Baz'}")
}
