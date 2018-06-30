package controllers

import (
	"net/http"

	"github.com/Pimmytrousers/GoSwarm/server/models"
	"github.com/Pimmytrousers/GoSwarm/server/views"
)

func NewBots(us models.UserService) *Bots {
	return &Bots{
		NewView: views.NewView("bootstrap", "bots/new"),
		us:      us,
	}

}

type Bots struct {
	NewView *views.View
	us      models.UserService
}

func (b *Bots) New(w http.ResponseWriter, r *http.Request) {
	if err := b.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}
