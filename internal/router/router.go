package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"

	"gowire/internal/handlers"
	"gowire/internal/middleware"
)

type RouteConfig struct {
	Path       string `json:"path"`
	Handler    string `json:"handler"`
	Middleware string `json:"middleware"`
}

type Config struct {
	Routes []RouteConfig `json:"routes"`
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

func SetupRoutes(h *handlers.Handlers) *http.ServeMux {
	mux := http.NewServeMux()

	config, err := loadConfig()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		os.Exit(1)
	}

	for _, route := range config.Routes {
		handler := getHandlerFunc(h, route.Handler)
		if handler == nil {
			fmt.Printf("Handler not found for route: %s\n", route.Path)
			continue
		}

		var middlewareStack []middleware.Middleware
		switch route.Middleware {
		case "public":
			middlewareStack = public
		case "private":
			middlewareStack = private
		default:
			fmt.Printf("Unknown middleware group for route: %s\n", route.Path)
			continue
		}

		mux.Handle(route.Path, middleware.Chain(handler, middlewareStack...))
	}

	return mux
}

func loadConfig() (*Config, error) {
	// configPath := os.Getenv("ROUTES_CONFIG_PATH")
	// if configPath == "" {
	// 	configPath = "config/routes.json" // default path
	// }

	// data, err := os.ReadFile(configPath)
	// if err != nil {
	// 	return nil, fmt.Errorf("error reading config file: %v", err)
	// }

	routes := `
	{
  "routes": [
    {
      "path": "/",
      "handler": "Home",
      "middleware": "public"
    },
    {
      "path": "/private",
      "handler": "PublicPage",
      "middleware": "private"
    }
  ]
}`

	var config Config
	err := json.Unmarshal([]byte(routes), &config)
	if err != nil {
		return nil, fmt.Errorf("error parsing config file: %v", err)
	}

	return &config, nil
}

func getHandlerFunc(h *handlers.Handlers, handlerName string) http.HandlerFunc {
	handlerValue := reflect.ValueOf(h).MethodByName(handlerName)
	if !handlerValue.IsValid() {
		return nil
	}
	return handlerValue.Interface().(func(http.ResponseWriter, *http.Request))
}
