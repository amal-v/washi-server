package api

import (
	"net/http"

	"github.com/go-chi/chi"
)

/*Router ... Will start serving the routes.. */
func Router() http.Handler {
	r := chi.NewRouter()
	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))

	})
	return r
}
