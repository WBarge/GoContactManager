package handlers

import "github.com/go-chi/chi/v5"

type PersonalEmailHandler struct {
	HandlerPatter string
}

func NewPersonalEmailHandler() *PersonalEmailHandler {
	return &PersonalEmailHandler{HandlerPatter: "/PersonalEmail"}
}

func (p *PersonalEmailHandler) AddHandlerRoutes(router chi.Router) {

}
