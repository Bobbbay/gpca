package main

type Standard struct {
    Content string `json:"content"`
}

type ReturnCode struct {
    // TODO: refactor for better name
    Response string `json:"response"`
    ErrorCode string `json:"errorcode"`
    Description string `json:"description"`
}

type Person struct {
    // id INTEGER PRIMARY KEY, name TEXT, points INTEGER, cryptocurrency TEXT)
    Id int `json:"id"`
    Name string `json:"name"`
    Points int `json:"points"`
    Cryptocurrency string `json:"cryptocurrency"`
}

type CryptocurrencyStatus struct {
    Date string `json:"date"`
    Points int `json:"points"`
}
