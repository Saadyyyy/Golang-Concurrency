package routes

import (
	"net/http"

	"github.com/go-chi/chi"
)

type caRouter struct {
	router *chi.Mux
}

func NewChRouter() Router {
	return &caRouter{}
}
func (cr *caRouter) SERVE(port string) {
	http.ListenAndServe(port, cr.router)
}

func (cr *caRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	cr.router.Get(uri, f)
}
