package public

import (
	"context"
	"net/http"
	"text/template"

	pb "github.com/Projectoutlast/nasa_proto/gen"
)

func (h *PublicHTTPHandlers) Registration(w http.ResponseWriter, r *http.Request) {
	session, _ := h.store.Get(r, "flash-session")
	messageFlashes := session.Flashes("error")
	session.Save(r, w)

	files := []string{
		"./assets/html/public/registration.html",
		baseSpaceLayout,
	}

	tmpl := template.Must(template.ParseFiles(files...))

	tmpl.Execute(w, messageFlashes)
}

func (h *PublicHTTPHandlers) RegistrationProcess(w http.ResponseWriter, r *http.Request) {
	email, password := r.FormValue("email"), r.FormValue("password")

	_, err := h.authClient.Registration(context.Background(), &pb.RegistrationRequest{
		Email:    email,
		Password: password,
	})

	session, _ := h.store.Get(r, "flash-session")

	if err != nil {
		session.AddFlash(err.Error(), "error")
		err := session.Save(r, w)

		if err != nil {
			h.log.Error(err.Error())
		}

		http.Redirect(w, r, "/registration", http.StatusSeeOther)
		return
	}

	session.AddFlash("Successfully registered!", "success")
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
