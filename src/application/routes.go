package application

import (
	"net/http"

	"github.com/WBarge/GoContactManager/handlers"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func (w *WebServer) RegisterHandlers() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	})

	prefix := "/api"

	router.Get(prefix, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("The contact service is up and running"))
	})

	peopleHandler := handlers.NewPeopleHandler()
	personHandler := handlers.NewPersonHandler()
	personalEmailHandler := handlers.NewPersonalEmailHandler()
	personalPhoneNumberHandler := handlers.NewPersonalPhoneNumberHandler()

	router.Route(prefix+peopleHandler.HandlerPatter, peopleHandler.AddHandlerRoutes)
	router.Route(prefix+personHandler.HandlerPatter, personHandler.AddHandlerRoutes)
	router.Route(prefix+personalEmailHandler.HandlerPatter, personalEmailHandler.AddHandlerRoutes)
	router.Route(prefix+personalPhoneNumberHandler.HandlerPatter, personalPhoneNumberHandler.AddHandlerRoutes)

	return router
}
