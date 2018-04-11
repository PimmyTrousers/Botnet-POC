package main

import (
	"fmt"
	"net/http"

	"github.com/Pimmytrousers/GoSwarm/Server/Front-End/views"
	"github.com/gorilla/mux"
)

var homeView *views.View

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	homeView.Render(w, nil)
}

func notfound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "404 yo ")
}

func server() {
	homeView = views.NewView("bootstrap", "Front-End/views/home.gohtml")
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.NotFoundHandler = http.HandlerFunc(notfound)
	http.ListenAndServe(":3000", r)
}

func main() {
	server()
}
