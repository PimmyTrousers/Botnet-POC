package main

import (
	"net/http"

	"github.com/Pimmytrousers/GoSwarm/server/controllers"
	"github.com/gorilla/mux"
)

func main() {
	usersC := controllers.NewUsers()
	staticC := controllers.NewStatic()
	botsC := controllers.NewBot()

	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	r.HandleFunc("/bots", botsC.New).Methods("GET")
	http.ListenAndServe(":3000", r)
}
