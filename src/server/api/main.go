package main

import (
    "log"
    "net/http"
    "os"
    "fmt"
)

func Test(num int)int{
    return num
}

func login(w http.ResponseWriter, req *http.Request){
    fmt.Fprint(w, "login information")
}

func index(w http.ResponseWriter, req *http.Request){
    fmt.Fprint(w, "index information")
}

func setupHandlers(mux *http.ServeMux){
    mux.HandleFunc("/", index)
    mux.HandleFunc("/login", login)
}

func main() {

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
        log.Printf("Defaulting to port %s", port)
    }
    mux := http.NewServeMux();
    setupHandlers(mux)
    log.Fatal(http.ListenAndServe(":"+port, mux))

}