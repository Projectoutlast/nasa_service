package httphandlers

import (
	"net/http"
	"text/template"
)

type indexData struct {
	PageTitle string
	Content   string
}

func (h *HTTPHandlers) Index(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./assets/html/index.html",
		baseSpaceLayout,
	}

	tmpl := template.Must(template.ParseFiles(files...))

	data := indexData{
		PageTitle: "Space Web App",
		Content:   "Welcome to my magic space",
	}

	tmpl.Execute(w, data)
}
