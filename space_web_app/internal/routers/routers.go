package routers

import (
	"github.com/Projectoutlast/space_service/space_web_app/internal/httphandlers"
	"github.com/gorilla/mux"
)

type Routers struct {
	Mux         *mux.Router
	httpHandlers *httphandlers.HTTPHandlers
}

func New(httpHandlers *httphandlers.HTTPHandlers) *Routers {
	return &Routers{
		Mux: mux.NewRouter(),
		httpHandlers: httpHandlers,
	}
}

func (r *Routers) SetUpHandlers() {
	r.Mux.HandleFunc("/random", r.httpHandlers.GetRandomSpaseImage).Methods("GET")
}
