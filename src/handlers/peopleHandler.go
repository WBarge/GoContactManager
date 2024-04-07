package handlers

import (
	"github.com/go-chi/chi/v5"
)

type PeopleHandler struct {
	HandlerPattern string
}

func NewPeopleHandler() *PeopleHandler {
	return &PeopleHandler{HandlerPattern: "/People"}
}

func (p *PeopleHandler) AddHandlerRoutes(router chi.Router) {

}
