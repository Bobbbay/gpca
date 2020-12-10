package main

import (
    "fmt"
    "encoding/json"
    "net/http"
)
func verify(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Hit on verify")

    if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
    }
    code := r.FormValue("code")
    fmt.Println(code)


    json.NewEncoder(w).Encode("{'Hello world':'Foo Baz'}")
}
