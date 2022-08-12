package main

import (
    "fmt"
    "net/http"
    "os"
)

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}

func main() {
    fmt.Println("Starting application...")

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello World")
    })

    http.ListenAndServe(":"+getEnv("PORT", "8080"), nil)
}
