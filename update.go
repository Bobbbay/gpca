package main

import (
    "fmt"
    "encoding/json"
    "net/http"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func update(w http.ResponseWriter, r *http.Request) {
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

    // BIG TODO: Link to database to update at this point
    database, err := sql.Open("sqlite3", "./gpca.db")
    if err != nil { fmt.Fprintf(w, "Error opening database: %v", err) }

    statement, err := database.Prepare("UPDATE people SET name = '" + name + "', points = " + points + ", cryptocurrency = '" + cryptocurrency + "' WHERE name = '" + name + "'")
    if err != nil { fmt.Println("Error modifying database:", err) }

    statement.Exec()

    json.NewEncoder(w).Encode( ReturnCode{
        Response: "success",
        ErrorCode: "200",
        Description: "Error 200 OK"})
}

