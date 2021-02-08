package main

import (
    "fmt"
    "encoding/json"
    "time"
    "math/rand"
    "strconv"
    "net/http"

    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func modifyCryptocurrency(difference int) {
    database, err := sql.Open("sqlite3", "./gpca.db")
    if err != nil { fmt.Printf("Error opening database: %v\n", err) }

    statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS cryptocurrency (date TEXT, points INTEGER)")
    if err != nil { fmt.Printf("Error initializing database: %v\n", err) }

    statement.Exec()

    var date string
    var points int
    rows, err := database.Query("SELECT * FROM cryptocurrency ORDER BY date DESC LIMIT 1")
    if err != nil { fmt.Printf("Error initializing database: %v\n", err) }
    for rows.Next() {
        rows.Scan(&date, &points)
    }

    fmt.Println(date, points)

    difference = difference + points

    statement, err = database.Prepare("INSERT INTO cryptocurrency (date, points) VALUES (?, ?)")
    if err != nil { fmt.Println("Error modifying database:", err) }

    statement.Exec(time.Now().Format("01-02-2006 3:4:5"), difference)
}

func claimCryptocurrency(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Hit on claimCryptocurrency")
    // TODO: make this take a name, and claim the cryptocurrency to the
    // account. After this, make sure to increment the cryptocurrency's
    // value by a random number between 0.5 and 1.

    // Parse form for `name` and `key` values (/cryptocurrency/claim?name=x&hash=y)
    if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
    }
    // TODO: Keep these values in a beautiful struct, e.x. request.name
    name := r.FormValue("name")
    hash := r.FormValue("hash")

    database, err := sql.Open("sqlite3", "./gpca.db")
    if err != nil { fmt.Fprintf(w, "Error opening database: %v", err) }

    rows, err := database.Query("SELECT id, name, cryptocurrency FROM people")
    if err != nil { fmt.Fprintf(w, "Error initializing database: %v", err) }

    var id int
    var userName string
    var cryptocurrency string

    for rows.Next() {
        rows.Scan(&id, &userName, &cryptocurrency)
    }

    statement, err := database.Prepare("UPDATE people SET cryptocurrency = '" + cryptocurrency + ", " + hash + "' WHERE name = '" + name + "'")
    if err != nil { fmt.Println("Error modifying database:", err) }

    statement.Exec()

    // Get random number (between 0.5 and 1) and increment the value of 
    // cryptocurrency
    rand.Seed(time.Now().UnixNano())
    // TODO: Currently, we generate a random number using rand.Intn.
    // Regrettably, this doesn't support floats - hence, we need multiply
    // by 10, get the random number, then divide by 10. Ugly, but it works.
    modifyCryptocurrency( (rand.Intn(10 - 5 + 10) + 5) / 10)

    json.NewEncoder(w).Encode( ReturnCode{
        Response: "success",
        ErrorCode: "200",
        Description: "Error 200 OK"})
}

func cryptocurrencyStatus(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Hit on status/cryptocurrency")
    setupCORS(&w)

    database, err := sql.Open("sqlite3", "./gpca.db")
    if err != nil { fmt.Fprintf(w, "Error opening database: %v", err) }

    statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS cryptocurrency (date TEXT, points INTEGER)")
    if err != nil { fmt.Fprintf(w, "Error initializing database: %v", err) }

    statement.Exec()

    rows, err := database.Query("SELECT date, points FROM cryptocurrency")
    if err != nil { fmt.Fprintf(w, "Error initializing database: %v", err) }

    // Given the following (sqlite) input:
    // ```
    // |   date   |   points   |
    // |----------|------------|
    // | 01-01-01 |      1     |
    // | 01-01-02 |      3     |
    // |----------|------------|
    // ```
    //
    // Produce the following (JSON) output:
    // ```
    // {
    //   "01-01-01": 1,
    //   "01-01-02": 3
    // }
    // ```
    //
    // Which in the frontend should be transformed into a struct:
    // ```
    // return CryptocurrencyStatus (
    //    dates: ["01-01-01", "01-01-02"],
    //    points: [1, 3]
    // )
    // ```
    // Although this may not be a beautiful solution - it works.

    // TODO: Currently, this just spits out a string. Ugly work. In the
    // future, this should use the json package to create a real json struct.

    // WARNING: I was not thinking straight here. This completely needs a rewrite.

    var date string
    var points int

    x := 0
    fmt.Fprintf(w, "{")
    for rows.Next() {
        if (x != 0) {
            fmt.Fprintf(w, ",\n")
        } else {
            fmt.Fprintf(w, "\n")
        }
        x++

        rows.Scan(&date, &points)
        row := "  \"" + date + "\": " + strconv.Itoa(points)
        fmt.Fprintf(w, row)
    }
    fmt.Fprintf(w, "\n}")
}

func setupCORS(w *http.ResponseWriter) {
    (*w).Header().Set("Access-Control-Allow-Origin", "*")
}


func verifyCryptocurrency(w http.ResponseWriter, r *http.Request) {
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
