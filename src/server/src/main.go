package main

import (
    "log"
    "net/http"
    "os"
)

func main() {

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
        log.Printf("Defaulting to port %s", port)
    }

    path := "dist/"
    http.Handle("/", http.FileServer(http.Dir(path)))
    log.Fatal(http.ListenAndServe(":"+port, nil))

}