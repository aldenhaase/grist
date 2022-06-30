package main

import (
    "log"
    "net/http"
    "flag"
)

func main() {
    var devFlag = flag.Bool("dev",false, "change server root dir")
    flag.Parse()
    var path string
    if *devFlag {
        path = "../../client/dist/lystr/"
    } else{
        path = "dist/lystr/"
    }
    http.Handle("/", http.FileServer(http.Dir(path)))
    log.Fatal(http.ListenAndServe(":8081", nil))

}