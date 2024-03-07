package main

import (
    "log"
    "net/http"
    "os"
    "github.com/oshaw1/SyncScribe/backend/handlers"
)

func main() {
    http.HandleFunc("/api/notes", handlers.HandleNotes)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    log.Fatal(http.ListenAndServe(":"+port, nil))
}