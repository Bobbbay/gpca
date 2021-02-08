package main

import (
    "testing"
    "time"
    "os/exec"
    "net/http"
)

// This file contains unit tests for the GPCA package. It nurtures
// both the testing and benchmark functions. As Go convention states,
// tests are prefixed with the `Test` function name, and benchmarks
// with the `Benchmark` function name. Instructions for running them
// are found in README.md.

func TestPackages(t *testing.T) {
    _, err := exec.Command("sh", "-c", "go list github.com/gorilla/mux github.com/mattn/go-sqlite3").Output()
    if err != nil { t.Errorf(err.Error()) }
}

func BenchmarkEndpoints(b *testing.B) {
    // TODO: use b.N and avoid accept error
    for i := 0; i < 1000; i++ {
        client := http.Client{ Timeout: time.Second * 2 }
        res, err := client.Get("http://localhost:10000")
        if err != nil { b.Errorf(err.Error()) }
        defer res.Body.Close()
    }
}
