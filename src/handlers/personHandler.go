package handlers

import "github.com/go-chi/chi/v5"

type PersonHandler struct {
	HandlerPatter string
}

func NewPersonHandler() *PersonHandler {
	return &PersonHandler{HandlerPatter: "/Person"}
}

func (p *PersonHandler) AddHandlerRoutes(router chi.Router) {

}
