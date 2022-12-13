package routes

import (
	"net/http"

	"github.com/VinayakBagaria/auth-micro-service/api/middlewares"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Route struct {
	Method       string
	Path         string
	Handler      http.HandlerFunc
	AuthRequired bool
}

func Install(router *mux.Router, routeList []*Route) {
	for _, route := range routeList {
		if route.AuthRequired {
			router.
				Handle(route.Path, middlewares.LogRequests(middlewares.Authenticate(route.Handler))).
				Methods(route.Method)
		} else {
			router.
				Handle(route.Path, middlewares.LogRequests(route.Handler)).
				Methods(route.Method)
		}
	}
}

func WithCORS(router *mux.Router) http.Handler {
	headers := handlers.AllowedHeaders([]string{"X-Requested-with", "Content-Type", "Accept", "Authorization"})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete})
	return handlers.CORS(headers, origins, methods)(router)
}
