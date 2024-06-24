package templates

import (
	"embed"
	"html/template"
)

func ParseTemplates(templateFS embed.FS) *template.Template {
	customTemplateFunctions := CustomTemplateFunctions()

	return template.Must(template.New("").Funcs(customTemplateFunctions).ParseFS(templateFS,
		"templates/*.*",
	))
}

func CustomTemplateFunctions() template.FuncMap {
	return template.FuncMap{
		//
	}
}
