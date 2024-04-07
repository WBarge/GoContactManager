package handlers

import "github.com/go-chi/chi/v5"

type PersonalPhoneNumberHandler struct {
	HandlerPatter string
}

func NewPersonalPhoneNumberHandler() *PersonalPhoneNumberHandler {
	return &PersonalPhoneNumberHandler{HandlerPatter: "/PersonalPhoneNumber"}
}

func (p *PersonalPhoneNumberHandler) AddHandlerRoutes(router chi.Router) {

}
