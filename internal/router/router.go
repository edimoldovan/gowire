package router

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"gowire/internal/handlers"
	"gowire/internal/middleware"
)

//go:embed config
var routesConfig embed.FS

type Route struct {
	Path       string `json:"path"`
	Handler    string `json:"handler"`
	Middleware string `json:"middleware"`
}

var (
	public = []middleware.Middleware{
		middleware.Logger,
	}
	private = []middleware.Middleware{
		middleware.Logger,
		middleware.Auth,
	}
)

func SetupRoutes() (*http.ServeMux, error) {
	h, err := handlers.NewHandlers()
	if err != nil {
		return nil, fmt.Errorf("error creating handlers: %v", err)
	}
	if h == nil {
		return nil, fmt.Errorf("handlers is nil after creation")
	}

	mux := http.NewServeMux()
	routes, err := loadConfig()
	if err != nil {
		return nil, fmt.Errorf("error loading config: %v", err)
	}

	for _, route := range *routes {
		handler, err := getHandlerFunc(h, route.Handler)
		if err != nil {
			log.Printf("Error getting handler for route %s: %v", route.Path, err)
			continue
		}

		var middlewareStack []middleware.Middleware
		switch route.Middleware {
		case "public":
			middlewareStack = public
		case "private":
			middlewareStack = private
		default:
			log.Printf("Unknown middleware group for route: %s", route.Path)
			continue
		}

		mux.Handle(route.Path, middleware.Chain(handler, middlewareStack...))
	}

	mux.HandleFunc("/light.js", h.ServeJS)
	mux.HandleFunc("/routes", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(routes)
	})

	return mux, nil
}

func loadConfig() (*[]Route, error) {
	routesJSON, err := routesConfig.ReadFile("config/routes.json")
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	var routes []Route
	err = json.Unmarshal(routesJSON, &routes)
	if err != nil {
		return nil, fmt.Errorf("error parsing config file: %v", err)
	}

	return &routes, nil
}

func getHandlerFunc(h *handlers.Handlers, handlerName string) (http.HandlerFunc, error) {
	if h == nil {
		return nil, fmt.Errorf("handlers is nil")
	}
	handlerValue := reflect.ValueOf(h).MethodByName(handlerName)
	if !handlerValue.IsValid() {
		return nil, fmt.Errorf("handler %s not found", handlerName)
	}
	return handlerValue.Interface().(func(http.ResponseWriter, *http.Request)), nil
}
