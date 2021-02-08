package main

import (
    "fmt"
    "encoding/json"
    "net/http"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

// This file contains functions and routes regarding people
// and entries in the database. Focusing on the informational
// aspect, it provides a place for many common routes to live.

func newPerson(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Hit on update")

    if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
    }

    username := r.FormValue("username")
    name := r.FormValue("name")
    permissions := r.FormValue("permissions")

    database, err := sql.Open("sqlite3", "./gpca.db")
    if err != nil { fmt.Fprintf(w, "Error opening database: %v", err) }

    statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, username TEXT, name TEXT, permissions TEXT, points INTEGER, cryptocurrency TEXT)")
    if err != nil { fmt.Fprintf(w, "Error initializing database: %v", err) }

    statement.Exec()

    statement, err = database.Prepare("INSERT INTO people (username, name, permissions, points, cryptocurrency) VALUES (?, ?, ?, ?, ?)")
    if err != nil { fmt.Println("Error modifying database:", err) }

    statement.Exec(username, name, permissions, 0, 0)

    json.NewEncoder(w).Encode( ReturnCode{
        Response: "success",
        ErrorCode: "200",
        Description: "Error 200 OK"})
}

func update(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Hit on update")

    if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
    }

    database, err := sql.Open("sqlite3", "./gpca.db")
    if err != nil { fmt.Fprintf(w, "Error opening database: %v", err) }

    rows, err := database.Query("SELECT username, name, permissions, points, cryptocurrency FROM people")
    if err != nil { fmt.Fprintf(w, "Error initializing database: %v", err) }

    var username string
    var name string
    var permissions string
    var points string
    var cryptocurrency string

    for rows.Next() {
        rows.Scan(&username, &name, &permissions, &points, &cryptocurrency)
    }

    if (len(r.FormValue("username")) > 0) { username = r.FormValue("username") }
    if (len(r.FormValue("name")) > 0) { name = r.FormValue("name") }
    if (len(r.FormValue("permissions")) > 0) { permissions = r.FormValue("permissions") }
    if (len(r.FormValue("points")) > 0) { points = r.FormValue("points") }
    if (len(r.FormValue("cryptocurrency")) > 0) { cryptocurrency = r.FormValue("cryptocurrency") }

    statement, err := database.Prepare("UPDATE people SET name = '" + name + "', points = " + points + ", cryptocurrency = '" + cryptocurrency + "' WHERE username = '" + username + "'")
    if err != nil { fmt.Println("Error modifying database:", err) }

    statement.Exec()

    json.NewEncoder(w).Encode( ReturnCode{
        Response: "success",
        ErrorCode: "200",
        Description: "Error 200 OK"})
}

func status(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Hit on status")
    setupCORS(&w)

    database, err := sql.Open("sqlite3", "./gpca.db")
    if err != nil { fmt.Fprintf(w, "Error opening database: %v", err) }

    rows, err := database.Query("SELECT id, username, name, permissions, points, cryptocurrency FROM people")
    if err != nil { fmt.Fprintf(w, "Error initializing database: %v", err) }

    switch r.Method {
    case "GET":
        var id int
        var username string
        var name string
        var permissions string
        var points int
        var cryptocurrency string

        // lol

        fmt.Fprintf(w, "{\n")

        for rows.Next() {
            rows.Scan(&id, &username, &name, &permissions, &points, &cryptocurrency)
            if (id > 1) { fmt.Fprintf(w, ",\"" + username + "\": ") } else { fmt.Fprintf(w, "\"" + username + "\": ") }
            json.NewEncoder(w).Encode( Person{
                Id: id,
                Username: username,
                Name: name,
                Permissions: permissions,
                Points: points,
                Cryptocurrency: cryptocurrency})
        }

        fmt.Fprintf(w, "}")

    case "POST":
        if err := r.ParseForm(); err != nil {
            fmt.Fprintf(w, "ParseForm() err: %v", err)
            return
        }

        username := r.FormValue("username")

        rows, err := database.Query("SELECT id, name, permissions, points, cryptocurrency FROM people WHERE username = " + username)
        if err != nil { fmt.Fprintf(w, "Error initializing database: %v", err) }

        var id int
        var name string
        var permissions string
        var points int
        var cryptocurrency string

        for rows.Next() {
            rows.Scan(&id, &name, &name)
            json.NewEncoder(w).Encode( Person{
                Id: id,
                Username: username,
                Name: name,
                Permissions: permissions,
                Points: points,
                Cryptocurrency: cryptocurrency})
        }

        json.NewEncoder(w).Encode( ReturnCode{
            Response: "success",
            ErrorCode: "200",
            Description: "Error 200 OK"})
    default:
        // TODO: Return error
    }
}
