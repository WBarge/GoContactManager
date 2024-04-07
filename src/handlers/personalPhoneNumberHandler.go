package handlers

import "github.com/go-chi/chi/v5"

type PersonalPhoneNumberHandler struct {
	HandlerPattern string
}

func NewPersonalPhoneNumberHandler() *PersonalPhoneNumberHandler {
	return &PersonalPhoneNumberHandler{HandlerPattern: "/PersonalPhoneNumber"}
}

func (p *PersonalPhoneNumberHandler) AddHandlerRoutes(router chi.Router) {

}
