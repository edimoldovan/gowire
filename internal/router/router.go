package router

import (
	"net/http"

	"gowire/internal/handlers"
	"gowire/internal/middleware"
)

var (
	public = []middleware.Middleware{
		middleware.Logger,
	}

	private = []middleware.Middleware{
		middleware.Logger,
		middleware.Auth,
	}
)

func SetupRoutes(h *handlers.Handlers) *http.ServeMux {
	mux := http.NewServeMux()

	// Home route with multiple middleware
	mux.Handle("/", middleware.Chain(
		http.HandlerFunc(h.Home),
		public...,
	))

	// Example of a route with only logger middleware
	mux.Handle("/private", middleware.Chain(
		http.HandlerFunc(h.PublicPage),
		private...,
	))

	return mux
}
