package routers

import (
	"net/http"

	"github.com/Projectoutlast/space_service/space_web_app/internal/httphandlers/public"
	"github.com/Projectoutlast/space_service/space_web_app/internal/httphandlers/secure"
	"github.com/Projectoutlast/space_service/space_web_app/internal/middleware"
	"github.com/gorilla/mux"
)

type Routers struct {
	Mux                *mux.Router
	publicHTTPHandlers *public.PublicHTTPHandlers
	secureHTTPHandlers *secure.SecureHTTPHandlers
	fileServerDir      string
	middleware         *middleware.Middleware
	staticPrefix       string
}

func New(
	httpHandlers *public.PublicHTTPHandlers,
	secureHTTPHandlers *secure.SecureHTTPHandlers,
	fileServerDir string,
	middleware *middleware.Middleware,
	staticPrefix string,
) *Routers {
	return &Routers{
		Mux:                mux.NewRouter(),
		publicHTTPHandlers: httpHandlers,
		secureHTTPHandlers: secureHTTPHandlers,
		fileServerDir:      fileServerDir,
		middleware:         middleware,
		staticPrefix:       staticPrefix,
	}
}

func (r *Routers) SetUpHandlers() {
	// Public routers
	r.Mux.HandleFunc("/", r.middleware.Logging(r.publicHTTPHandlers.Index)).Methods("GET")
	r.Mux.HandleFunc("/registration", r.middleware.Logging(r.publicHTTPHandlers.Registration)).Methods("GET")
	r.Mux.HandleFunc("/registration-process", r.middleware.Logging(r.publicHTTPHandlers.RegistrationProcess)).Methods("POST")
	r.Mux.HandleFunc("/login", r.middleware.Logging(r.publicHTTPHandlers.Login)).Methods("GET")
	r.Mux.HandleFunc("/login-process", r.middleware.Logging(r.publicHTTPHandlers.LoginProcess)).Methods("POST")

	// Secure routers
	r.Mux.HandleFunc("/home", r.middleware.SecureLogging(r.secureHTTPHandlers.Index)).Methods("GET")
	r.Mux.HandleFunc("/random", r.middleware.SecureLogging(r.secureHTTPHandlers.GetRandomSpaseImage)).Methods("GET")
}

func (r *Routers) SetUpFileServer() {
	fs := http.FileServer(http.Dir(r.fileServerDir))
	r.Mux.Handle(r.staticPrefix, http.StripPrefix(r.staticPrefix, fs))
}
