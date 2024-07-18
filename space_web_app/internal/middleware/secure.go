package middleware

import (
	"net/http"
	"time"
)

func (m *Middleware) SecureLogging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		m.log.Info("started handling request", "method", r.Method, "url", r.URL.String())

		authToken, err := r.Cookie("auth-token")
		if err != nil {
			m.log.Error(err.Error())
			http.Redirect(w, r, "/login", http.StatusUnauthorized)
			return
		}

		_, err = m.validator.GetToken(authToken.Value)
		if err != nil {
			m.log.Error(err.Error())
			http.Redirect(w, r, "/login", http.StatusUnauthorized)
			return
		}

		f(w, r)
		m.log.Info("sent response to request", "method", r.Method, "url", r.URL.String(), "duration", time.Since(start))
	}
}
