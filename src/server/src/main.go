package main

import (
    "log"
    "net/http"
    "os"
    "io"
    "fmt"
    "strings"
    "regexp"
)

var dir = http.Dir("dist/")
var fileserver = http.FileServer(dir)
var jsType = regexp.MustCompile("\\.js$")

func Test(num int)int{
    return num;
}

func login(w http.ResponseWriter, req *http.Request){
    buf := new(strings.Builder)
    io.Copy(buf, req.Body)
    fmt.Fprintf(w,buf.String())
}

func index(w http.ResponseWriter, req *http.Request){
    reqURI := req.RequestURI
    if(jsType.MatchString(reqURI)){
        w.Header().Set("Content-Type","text/javascript")
    }
    fileserver.ServeHTTP(w,req)
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