package main

import (
	"fmt"
	"net/http"

	"github.com/Pimmytrousers/GoSwarm/server/controllers"
	"github.com/Pimmytrousers/GoSwarm/server/models"
	"github.com/gorilla/mux"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "Pim"
	password = ""
	dbname   = "GoSwarm"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	us.AutoMigrate()

	usersC := controllers.NewUsers(us)
	staticC := controllers.NewStatic()
	botsC := controllers.NewBots(us)

	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")

	r.Handle("/contact", staticC.Contact).Methods("GET")

	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")

	r.HandleFunc("/bots", botsC.New).Methods("GET")

	r.Handle("/login", usersC.LoginView).Methods("GET")
	r.HandleFunc("/login", usersC.Login).Methods("POST")
	r.HandleFunc("/cookietest", usersC.CookieTest).Methods("GET")
	http.ListenAndServe(":3000", r)
}
