package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "User Service is up!")
    })

    http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "User registration endpoint")
    })

    http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "User login endpoint")
    })

    fmt.Println("Starting User Service on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
