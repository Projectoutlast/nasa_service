package httpsessions

import (
	"net/http"

	"github.com/gorilla/sessions"
)

func New() *sessions.CookieStore {
	store := sessions.NewCookieStore([]byte("something-very-secret"))
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600 * 8,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}

	return store
}
