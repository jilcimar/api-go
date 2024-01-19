package route

import "github.com/gorilla/mux"

// Retornar o route
func Generate() *mux.Router {
	return mux.NewRouter()
}
