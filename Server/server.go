package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Home Page")
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Test Page")
}
func StartServer() {
	r := mux.NewRouter()

	//Define Routes
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/test", TestHandler)

	http.ListenAndServe(":8080", r)
}
