package handlers

import "github.com/go-chi/chi/v5"

type PersonHandler struct {
	HandlerPattern string
}

func NewPersonHandler() *PersonHandler {
	return &PersonHandler{HandlerPattern: "/Person"}
}

func (p *PersonHandler) AddHandlerRoutes(router chi.Router) {

}
