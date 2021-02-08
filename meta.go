package main

import (
    "fmt"
    "encoding/json"
    "net/http"
    "strconv"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

// In this file, you'll find regular meta functions.
// This ranges from the home function, to the database
// status function, and so on.

func home(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Hit on home") // TODO: don't require to do this on every function
    result := "Welcome to the GPCA."
    json.NewEncoder(w).Encode(result)
}

func status(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Hit on status")

    database, err := sql.Open("sqlite3", "./gpca.db")
    if err != nil { fmt.Fprintf(w, "Error opening database: %v", err) }

    rows, err := database.Query("SELECT id, name, points, cryptocurrency FROM people")
    if err != nil { fmt.Fprintf(w, "Error initializing database: %v", err) }

    switch r.Method {
    case "GET":
        var id int
        var name string
        var points int
        var cryptocurrency string

        for rows.Next() {
            rows.Scan(&id, &name, &points, &cryptocurrency)
            // fmt.Fprintf(w, strconv.Itoa(id) + " " + name + " " + strconv.Itoa(points) + " " + cryptocurrency)
            json.NewEncoder(w).Encode( Person{
                Id: id,
                Name: name,
                Points: points,
                Cryptocurrency: cryptocurrency})
        }

    case "POST":
        // Parse form for `name` and `key` values (/status?name=x&key=y)
        if err := r.ParseForm(); err != nil {
            fmt.Fprintf(w, "ParseForm() err: %v", err)
            return
        }
        // TODO: Keep these values in a beautiful struct, e.x. request.name
        name := r.FormValue("name")
        key := r.FormValue("key")

        fmt.Println(name, key)

        var id int
        var firstname string
        var lastname string
        for rows.Next() {
            rows.Scan(&id, &firstname, &lastname)
            fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
        }

        json.NewEncoder(w).Encode( ReturnCode{
            Response: "success",
            ErrorCode: "200",
            Description: "Error 200 OK"})
    default:
        // TODO: Return error
    }
}
