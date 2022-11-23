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
	mux.HandleFunc("/checkAuth", endpoints.CheckAuth)
	mux.HandleFunc("/registerNewUser", validate.Json(endpoints.RegisterNewUser))
	mux.HandleFunc("/logIn", validate.Json(endpoints.LogIn))
	mux.HandleFunc("/getUserList", endpoints.GetUserList)
	mux.HandleFunc("/deleteUserList", endpoints.DeleteUserList)
	mux.HandleFunc("/setUserList", endpoints.SetUserList)
	mux.HandleFunc("/deleteListItem", endpoints.DeleteListItem)
	mux.HandleFunc("/createUserList", endpoints.CreateUserList)
	mux.HandleFunc("/enumerateLists", endpoints.EnumerateLists)
	mux.HandleFunc("/getRegistrationCookies", endpoints.GetRegistrationCookies)
	mux.HandleFunc("/checkForUpdates", endpoints.CheckForUpdates)
	mux.HandleFunc("/checkListArray", endpoints.CheckListArray)
	////NEW
	mux.HandleFunc("/cookieAuthenticator", endpoints.AuthenticateCookie)
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
	http.Handle("/", middleware.All(mux))
	appengine.Main()
}
