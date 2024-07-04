package routers

import (
	"net/http"

	"github.com/Projectoutlast/space_service/space_web_app/internal/httphandlers"
	"github.com/gorilla/mux"
)

type Routers struct {
	Mux           *mux.Router
	httpHandlers  *httphandlers.HTTPHandlers
	fileServerDir string
	staticPrefix  string
}

func New(
	httpHandlers *httphandlers.HTTPHandlers,
	fileServerDir string,
	staticPrefix string,
) *Routers {
	return &Routers{
		Mux:           mux.NewRouter(),
		httpHandlers:  httpHandlers,
		fileServerDir: fileServerDir,
		staticPrefix:  staticPrefix,
	}
}

func (r *Routers) SetUpHandlers() {
	r.Mux.HandleFunc("/", r.httpHandlers.Index).Methods("GET")
	r.Mux.HandleFunc("/random", r.httpHandlers.GetRandomSpaseImage).Methods("GET")
	r.Mux.HandleFunc("/registration", r.httpHandlers.Registration).Methods("GET")
	r.Mux.HandleFunc("/registration-process", r.httpHandlers.RegistrationProcess).Methods("POST")
}

func (r *Routers) SetUpFileServer() {
	fs := http.FileServer(http.Dir(r.fileServerDir))
	r.Mux.Handle(r.staticPrefix, http.StripPrefix(r.staticPrefix, fs))
}
