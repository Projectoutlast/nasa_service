package public

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/Projectoutlast/space_service/space_web_app/internal/authenticator"
	"github.com/gorilla/sessions"
)

func (p *PublicHTTPHandlers) Handler(auth *authenticator.Authenticator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		state, err := generateRandomState()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		session, err := p.store.Get(r, "state")
		session.Values["state"] = state
		err = sessions.Save(r, w)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, auth.AuthCodeURL(state), http.StatusTemporaryRedirect)
	}
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}
