package main

import (
    "fmt"
    "encoding/json"
    "net/http"
    "strconv"
)

func verify(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Hit on verify")

    // Parse form for `code` value (/verify?code=x)
    if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
    }
    // For TODOs, see update.go
    code := r.FormValue("code")

    if len(code) == 64 {
        // Convert `code` to int, and return if it works
        if _, err := strconv.Atoi(code[:5]); err == nil {
                json.NewEncoder(w).Encode( ReturnCode{
                Response: "success",
                ErrorCode: "200",
                Description: "Error 200 OK"})
        } else {
            json.NewEncoder(w).Encode( ReturnCode{
                Response: "fail",
                ErrorCode: "GPCA.02",
                Description: "Error .02 non-valid-block"})
        }
    } else {
        json.NewEncoder(w).Encode( ReturnCode{
            Response: "fail",
            ErrorCode: "GPCA.01",
            Description: "Error .01 not-a-sum"})
    }
}

