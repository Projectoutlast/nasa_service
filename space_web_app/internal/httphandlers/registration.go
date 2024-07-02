package httphandlers

import (
	"net/http"
	"text/template"
)

func (h *HTTPHandlers) Registration(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./assets/html/registration.html",
		baseSpaceLayout,
	}

	tmpl := template.Must(template.ParseFiles(files...))

	tmpl.Execute(w, &RegistrationResponse{PageTitle: "Registration"})
}
