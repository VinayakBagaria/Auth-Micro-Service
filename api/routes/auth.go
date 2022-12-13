package routes

import (
	"net/http"

	"github.com/VinayakBagaria/auth-micro-service/api/resthandlers"
)

func NewAuthRoutes(authHandlers resthandlers.AuthHandlers) []*Route {
	return []*Route{
		{Path: "/signup", Method: http.MethodPost, Handler: authHandlers.SignUp},
		{Path: "/signin", Method: http.MethodPost, Handler: authHandlers.SignIn},
		{Path: "/users", Method: http.MethodGet, Handler: authHandlers.GetUsers, AuthRequired: true},
		{Path: "/user/{id}", Method: http.MethodGet, Handler: authHandlers.GetUser, AuthRequired: true},
		{Path: "/user/{id}", Method: http.MethodPut, Handler: authHandlers.PutUser, AuthRequired: true},
		{Path: "/user/{id}", Method: http.MethodDelete, Handler: authHandlers.DeleteUser, AuthRequired: true},
	}
}
