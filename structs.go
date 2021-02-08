package main

// In this file, you'll find all the structs used throughout the
// project. They're a great reference point to go if you don't
// understand precisely what one block of code is meant to do.

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
    Id          int     `json:"id"`
    Username    string  `json:"username"`
    Permissions string `json:"permissions"`
    Name        string  `json:"name"`
    Points      int     `json:"points"`
    Cryptocurrency string `json:"cryptocurrency"`
}

type CryptocurrencyStatus struct {
    Date string `json:"date"`
    Points int `json:"points"`
}
