package templates

import (
	"embed"
	"fmt"
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
		"link": func(url, text, target string) template.HTML {
			return template.HTML(fmt.Sprintf(
				`<a href="%s" data-url="%s" data-target="%s">%s</a>`,
				url, url, target, text,
			))
		},
		"add": func(a, b int) int {
			return a + b
		},
	}
}
