package main

import (
    "log"
    "net/http"
    "os"
    "io"
    "fmt"
    "strings"
)

func Test(num int)int{
    return num;
}

type logLine struct{
    UserIP string `json:"user_ip"`
    Event string `json:"event"`
}

func login(w http.ResponseWriter, req *http.Request){
    buf := new(strings.Builder)
    io.Copy(buf, req.Body)
    fmt.Fprintf(w,buf.String())
}

func setupHandlers(mux *http.ServeMux){
    mux.Handle("/", http.FileServer(http.Dir("dist/")))
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