package main

import (
    "fmt"
    "encoding/json"
    "net/http"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func newEntry(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Hit on update")

    // Parse form for `name` and `key` values (/new?name=x&points=y&cryptocurrency=z)
    if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
    }
    // TODO: Keep these values in a beautiful struct, e.x. request.name
    name := r.FormValue("name")
    points := r.FormValue("points")
    cryptocurrency := r.FormValue("cryptocurrency")

    if true {
        // Convert `code` to int, and return if it works
        if true {
            // BIG TODO: Link to database to update at this point
            database, err := sql.Open("sqlite3", "./gpca.db")
            if err != nil { fmt.Fprintf(w, "Error opening database: %v", err) }

            // people database: id (key), name (str), points (int), cryptocurrency (string/json)
            statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, name TEXT, points INTEGER, cryptocurrency TEXT)")
            if err != nil { fmt.Fprintf(w, "Error initializing database: %v", err) }

            _, err = statement.Exec()
            if err != nil {fmt.Println("Hello world!")}

            statement, err = database.Prepare("INSERT INTO people (name, points, cryptocurrency) VALUES (?, ?, ?)")
            if err != nil { fmt.Println("Error modifying database:", err) }

            _, err = statement.Exec(name, points, cryptocurrency)
            if err != nil {fmt.Println("Hello world!")}

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

