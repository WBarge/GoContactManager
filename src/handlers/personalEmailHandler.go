package handlers

import "github.com/go-chi/chi/v5"

type PersonalEmailHandler struct {
	HandlerPattern string
}

func NewPersonalEmailHandler() *PersonalEmailHandler {
	return &PersonalEmailHandler{HandlerPattern: "/PersonalEmail"}
}

func (p *PersonalEmailHandler) AddHandlerRoutes(router chi.Router) {

}
