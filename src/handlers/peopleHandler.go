package handlers

import (
	"github.com/go-chi/chi/v5"
)

type PeopleHandler struct {
	HandlerPatter string
}

func NewPeopleHandler() *PeopleHandler {
	return &PeopleHandler{HandlerPatter: "/People"}
}

func (p *PeopleHandler) AddHandlerRoutes(router chi.Router) {

}
