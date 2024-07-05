package routers

import (
	"net/http"

	"github.com/Projectoutlast/space_service/space_web_app/internal/httphandlers"
	"github.com/Projectoutlast/space_service/space_web_app/internal/middleware"
	"github.com/gorilla/mux"
)

type Routers struct {
	Mux           *mux.Router
	httpHandlers  *httphandlers.HTTPHandlers
	fileServerDir string
	middleware    *middleware.Middleware
	staticPrefix  string
}

func New(
	httpHandlers *httphandlers.HTTPHandlers,
	fileServerDir string,
	middleware *middleware.Middleware,
	staticPrefix string,
) *Routers {
	return &Routers{
		Mux:           mux.NewRouter(),
		httpHandlers:  httpHandlers,
		fileServerDir: fileServerDir,
		middleware:    middleware,
		staticPrefix:  staticPrefix,
	}
}

func (r *Routers) SetUpHandlers() {
	r.Mux.HandleFunc("/", r.middleware.Logging(r.httpHandlers.Index)).Methods("GET")
	r.Mux.HandleFunc("/random", r.middleware.Logging(r.httpHandlers.GetRandomSpaseImage)).Methods("GET")
	r.Mux.HandleFunc("/registration", r.middleware.Logging(r.httpHandlers.Registration)).Methods("GET")
	r.Mux.HandleFunc("/registration-process", r.middleware.Logging(r.httpHandlers.RegistrationProcess)).Methods("POST")
}

func (r *Routers) SetUpFileServer() {
	fs := http.FileServer(http.Dir(r.fileServerDir))
	r.Mux.Handle(r.staticPrefix, http.StripPrefix(r.staticPrefix, fs))
}
