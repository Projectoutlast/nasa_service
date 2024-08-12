package middleware

import (
	"net/http"
	"time"
)

func (m *Middleware) SecureLogging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		m.log.Info("started handling request", "method", r.Method, "url", r.URL.String())

		session, err := m.store.Get(r, "state")

		if err != nil {
			m.log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if session.Values["profile"] == nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		f(w, r)
		m.log.Info("sent response to request", "method", r.Method, "url", r.URL.String(), "duration", time.Since(start))
	}
}
