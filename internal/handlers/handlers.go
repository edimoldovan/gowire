package handlers

import (
	"embed"
	"fmt"
	"gowire/internal/handlers/templates"
	"html/template"
	"log"
	"net/http"
)

//go:embed templates/*.html
var templateFS embed.FS

var Tmpl *template.Template

type Handlers struct{}

func init() {
	Tmpl = templates.ParseTemplates(templateFS)
}

func NewHandlers() (*Handlers, error) {
	return &Handlers{}, nil
}

func RenderHTML(w http.ResponseWriter, templateName string, params map[string]interface{}) {
	if err := Tmpl.ExecuteTemplate(w, templateName, params); err != nil {
		fmt.Printf("ERR: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h Handlers) Home(w http.ResponseWriter, r *http.Request) {
	log.Println("Home handler")
	if r.Header.Get("X-Requested-With") == "LightFramework" {
		RenderHTML(w, "home_content", nil)
	} else {
		RenderHTML(w, "home", map[string]interface{}{
			"Title": "Home",
		})
	}
}

func (h Handlers) PublicPage(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-Requested-With") == "LightFramework" {
		RenderHTML(w, "private_content", nil)
	} else {
		RenderHTML(w, "private", map[string]interface{}{
			"Title": "Private",
		})
	}
}

func (h Handlers) ServeJS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")
	fmt.Fprint(w, `
	class LightFramework {
		constructor(routes) {
			this.routes = routes;
			this.setupEventListeners();
		}

		setupEventListeners() {
			document.addEventListener('click', async (event) => {
				const target = event.target.closest('[data-light-action]');
				if (target) {
					event.preventDefault();
					const action = target.getAttribute('data-light-action');
					const route = this.routes.find(r => r.handler === action);
					if (route) {
						await this.handleAction(route);
					}
				}
			});
		}

		async handleAction(route) {
			try {
				const response = await fetch(route.path, {
					method: 'GET',
					headers: {
						'X-Requested-With': 'LightFramework'
					}
				});
				const html = await response.text();
				document.querySelector('#content').innerHTML = html;
			} catch (error) {
				console.error('Error:', error);
			}
		}
	}

	// Initialize the framework with the routes
	(async () => {
		try {
			const response = await fetch('/routes');
			const routes = await response.json();
			const light = new LightFramework(routes);
		} catch (error) {
			console.error('Error initializing LightFramework:', error);
		}
	})();
	`)
}
