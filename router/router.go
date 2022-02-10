package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

// PokemonController is the interface that wraps the controller's methods
// Catch
type PokemonController interface {
	Catch(w http.ResponseWriter, r *http.Request)
}

// Setup returns a router instance
func Setup(c PokemonController) *mux.Router {
	r := mux.NewRouter()

	// versioning api
	v1 := r.PathPrefix("/api/v1/pokemon").Subrouter()

	// the endpoints
	v1.HandleFunc("/catch", c.Catch).
		Methods(http.MethodGet).Name("Catch")

	return r
}
