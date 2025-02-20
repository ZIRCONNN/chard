package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func goodbyeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Goodbye, World!")
}

func main() {
    http.HandleFunc("/hello", helloHandler)
    http.HandleFunc("/goodbye", goodbyeHandler)

    fmt.Println("Server is listening on port 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Failed to start server:", err)
    }
}
