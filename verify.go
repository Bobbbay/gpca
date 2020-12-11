package main

import (
    "fmt"
    "encoding/json"
    "net/http"
    "strconv"
)

type ReturnCode struct { // TODO: refactor for better name
    Response string `json:"response"`
    ErrorCode string `json:"errorcode"`
    Description string `json:"description"`
}

func verify(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Hit on verify")

    // Parse form for `code`
    if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
    }
    code := r.FormValue("code")

    if len(code) > 5 {
        // Convert `code` to int, and return if it works
        if _, err := strconv.Atoi(code[:5]); err == nil {
            // TODO: Link to database to update at this point
            json.NewEncoder(w).Encode( ReturnCode{
                Response: "success",
                ErrorCode: "200",
                Description: "Error 200 OK"})
        } else {
            json.NewEncoder(w).Encode( ReturnCode{
                Response: "fail",
                ErrorCode: "GPCA.02",
                Description: "Error .02 code-invalid"})
        }
    } else {
        json.NewEncoder(w).Encode( ReturnCode{
            Response: "fail",
            ErrorCode: "GPCA.01",
            Description: "Error .01 less-than-6"})
    }
}
