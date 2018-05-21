package controllers

import (
	"net/http"

	"github.com/Pimmytrousers/GoSwarm/server/views"
)

func NewBot() *Bots {
	return &Bots{
		NewView: views.NewView("bootstrap", "bots/new"),
	}

}

func (b *Bots) New(w http.ResponseWriter, r *http.Request) {
	if err := b.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

type Bots struct {
	NewView *views.View
}
