package main

import (
	"log"
	"net/http"
	"os"
	"server/endpoints"
	"server/middleware"
	"server/validate"

	"google.golang.org/appengine/v2"
)

func Test(num int) int {
	return num
}

func root(res http.ResponseWriter, req *http.Request) {

}

func setupHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/checkUsername", validate.Json(endpoints.CheckUsername))
	mux.HandleFunc("/registerNewUser", validate.Json(endpoints.RegisterNewUser))
	mux.HandleFunc("/logIn", validate.Json(endpoints.LogIn))
	mux.HandleFunc("/getRegistrationCookies", endpoints.GetRegistrationCookies)
	mux.HandleFunc("/cookieAuthenticator", endpoints.AuthenticateCookie)
	//authorization required
	mux.HandleFunc("/listSetter", middleware.Auth(endpoints.ListSetter))
	mux.HandleFunc("/listGrabber", middleware.Auth(endpoints.ListGrabber))
	mux.HandleFunc("/listDeleter", middleware.Auth(endpoints.ListDeleter))
	mux.HandleFunc("/itemDeleter", middleware.Auth(endpoints.ItemDeleter))
	mux.HandleFunc("/collaborator", middleware.Auth(endpoints.Collaborator))
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
	http.Handle("/", middleware.Headers(mux))
	appengine.Main()
}
