package handlers

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"
)

//go:embed templates/*.html
var templateFS embed.FS

type Handlers struct {
	templates *template.Template
}

func NewHandlers() (*Handlers, error) {
	tmpl, err := parseTemplates()
	if err != nil {
		return nil, err
	}
	return &Handlers{templates: tmpl}, nil
}

func parseTemplates() (*template.Template, error) {
	tmpl := template.New("")

	err := fs.WalkDir(templateFS, "templates", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			content, err := fs.ReadFile(templateFS, path)
			if err != nil {
				return err
			}
			_, err = tmpl.New(d.Name()).Parse(string(content))
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return tmpl, nil
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	h.templates.ExecuteTemplate(w, "home.html", nil)
}

func (h *Handlers) PublicPage(w http.ResponseWriter, r *http.Request) {
	h.templates.ExecuteTemplate(w, "private.html", nil)
}
