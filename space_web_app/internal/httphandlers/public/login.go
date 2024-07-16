package httphandlers

import (
	"context"
	"net/http"
	"text/template"

	pb "github.com/Projectoutlast/nasa_proto/gen"
)

func (h *HTTPHandlers) Login(w http.ResponseWriter, r *http.Request) {
	session, _ := h.store.Get(r, "flash-session")
	errorFlashes := session.Flashes("error")
	session.Save(r, w)

	files := []string{
		"./assets/html/public/login.html",
		baseSpaceLayout,
	}

	tmpl := template.Must(template.ParseFiles(files...))

	tmpl.Execute(w, errorFlashes)
}

func (h *HTTPHandlers) LoginProcess(w http.ResponseWriter, r *http.Request) {
	email, password := r.FormValue("email"), r.FormValue("password")

	_, err := h.authClient.Login(context.Background(), &pb.LoginRequest{
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

		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	session.AddFlash("Successfully logged in!", "success")
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
