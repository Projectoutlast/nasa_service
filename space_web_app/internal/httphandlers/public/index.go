package httphandlers

import (
	"net/http"
	"text/template"
)

type indexData struct {
	Content    string
	SuccessMsg []interface{}
	ErrorMsg   []interface{}
}

func (h *HTTPHandlers) Index(w http.ResponseWriter, r *http.Request) {
	session, _ := h.store.Get(r, "flash-session")

	successMessageFlashes := session.Flashes("success")
	errorMessageFlashes := session.Flashes("error")

	session.Save(r, w)

	data := indexData{
		Content:    "Welcome to NASA space sevice",
		SuccessMsg: successMessageFlashes,
		ErrorMsg:   errorMessageFlashes,
	}

	files := []string{
		"./assets/html/public/index.html",
		baseSpaceLayout,
	}

	tmpl := template.Must(template.ParseFiles(files...))

	tmpl.Execute(w, data)
}
