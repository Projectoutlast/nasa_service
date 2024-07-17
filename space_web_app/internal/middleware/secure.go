package middleware

import (
	"net/http"
	"strings"
	"time"
)

func (m *Middleware) SecureLogging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		m.log.Info("started handling request", "method", r.Method, "url", r.URL.String())

		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			m.log.Error("token header is empty")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		splitToken := strings.Split(tokenHeader, " ")
		if len(splitToken) != 2 || splitToken[0] != "Bearer" {
			m.log.Error("token header is malformed")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		token := splitToken[1]

		_, err := m.validator.GetToken(token)
		if err != nil {
			m.log.Error(err.Error())
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		f(w, r)
		m.log.Info("sent response to request", "method", r.Method, "url", r.URL.String(), "duration", time.Since(start))
	}
}
