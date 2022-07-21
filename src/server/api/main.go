package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"server/datastore/queries"

	"google.golang.org/appengine/v2"
)

func Test(num int) int {
	return num
}

type user struct {
	Valid bool
}

func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		headers := req.Header.Get("Access-Control-Request-Headers")
		(res).Header().Set("Access-Control-Allow-Headers", headers)
		(res).Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(res, req)
	})
}

func checkUsername(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	_, err := queries.UserExists(ctx, "test")
	if err != nil {
		fmt.Fprint(res, err)
	} else {
		encoder := json.NewEncoder(res)
		encoder.Encode(user{Valid: true})
		if err != nil {
			panic(err)
		}
	}
}

func root(res http.ResponseWriter, req *http.Request) {

}

func setupHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/checkUsername", checkUsername)
	mux.HandleFunc("/", root)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	mux := http.NewServeMux()
	setupHandlers(mux)
	http.Handle("/", enableCors(mux))
	appengine.Main()
}
