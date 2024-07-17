package secure

import (
	"html/template"
	"net/http"
)

func (s *SecureHTTPHandlers) Index(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./assets/html/secure/index.html",
		baseSpaceLayout,
	}

	tmpl, err := template.ParseFiles(files...)

	if err != nil {
		s.log.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}
