package routes

import (
	"net/http"

	"github.com/VinayakBagaria/auth-micro-service/api/resthandlers"
)

func NewAuthRoutes(authHandlers resthandlers.AuthHandlers) []*Route {
	return []*Route{
		{Path: "/signup", Method: http.MethodPost, Handler: authHandlers.SignUp},
		{Path: "/user/{id}", Method: http.MethodPut, Handler: authHandlers.PutUser},
		{Path: "/user/{id}", Method: http.MethodGet, Handler: authHandlers.GetUser},
		{Path: "/users", Method: http.MethodGet, Handler: authHandlers.GetUsers},
		{Path: "/user/{id}", Method: http.MethodDelete, Handler: authHandlers.DeleteUser},
	}
}
