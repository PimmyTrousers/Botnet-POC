package Front_End

import "github.com/gorilla/mux"
import (
	"github.com/Pimmytrousers/GoSwarm/Server/Front-End/views"
	"net/http"
)

var homeView *views.View

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	homeView.Render(w, nil)
}

func ControlPanel() {
	homeView = views.NewView("bootstrap", "views/home.html")
	r := mux.NewRouter()
	r.HandleFunc("/", home)
}