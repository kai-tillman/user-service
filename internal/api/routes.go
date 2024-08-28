package api

import (
	"github.com/gorilla/mux"
	"net/http"
)


type Route struct {
	Name    string
	Path    string
	Method  string
	Handler http.HandlerFunc
}

type Routes []Route

func (h *APIHandler) AddRoutesToGorillaMux(router *mux.Router) {
	for _, route := range h.GetRoutes() {
		router.
			Name(route.Name).
			Path(route.Path).
			Methods(route.Method).
			Handler(route.Handler)
	}
}

func (h *APIHandler) GetRoutes() Routes {
	return Routes{
		{
			"registerUser",
			"/register",
			"POST",
			h.HandleRegisterUser,
		},{
			"deleteUserById",
			"/users/{userId}",
			"DELETE",
			h.HandleDeleteUserById,
		},{
			"getUserById",
			"/users/{userId}",
			"GET",
			h.HandleGetUserById,
		},{
			"updateUserById",
			"/users/{userId}",
			"PUT",
			h.HandleUpdateUserById,
		},
	}
}

