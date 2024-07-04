package httphandlers

import (
	"context"
	"net/http"
	"text/template"

	pb "github.com/Projectoutlast/nasa_proto/gen"
)

func (h *HTTPHandlers) Registration(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./assets/html/registration.html",
		baseSpaceLayout,
	}

	tmpl := template.Must(template.ParseFiles(files...))

	tmpl.Execute(w, &RegistrationResponse{PageTitle: "Registration"})
}

func (h *HTTPHandlers) RegistrationProcess(w http.ResponseWriter, r *http.Request) {
	email, password := r.FormValue("email"), r.FormValue("password")

	if email == "" || password == "" {
		h.log.Error("email or password is empty")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	_, err := h.authClient.Register(context.Background(), &pb.RegistrationRequest{
		Email:    email,
		Password: password,
	})

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
