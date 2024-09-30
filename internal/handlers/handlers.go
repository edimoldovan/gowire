package handlers

import (
	"embed"
	"fmt"
	"gowire/internal/handlers/templates"
	"html/template"
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
	RenderHTML(w, "home", map[string]interface{}{
		"Title": "Home",
	})
}

func (h Handlers) About(w http.ResponseWriter, r *http.Request) {
	RenderHTML(w, "about", map[string]interface{}{
		"Title": "About",
	})
}

func (h Handlers) Contact(w http.ResponseWriter, r *http.Request) {
	RenderHTML(w, "contact", map[string]interface{}{
		"Title": "Contact",
	})
}

func (h Handlers) Private(w http.ResponseWriter, r *http.Request) {
	RenderHTML(w, "private", map[string]interface{}{
		"Title": "Private",
	})
}

func (h Handlers) ServeJS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")
	http.ServeFile(w, r, "web/static/js/light.js")
}
