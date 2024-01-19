package route

import (
	"github.com/gorilla/mux"
	"net/http"
)

// Representa todas as rotas da API
type Rota struct {
	URI                    string
	Method                 string
	Function               func(w http.ResponseWriter, r *http.Request)
	RequiresAuthentication bool
}

// Coloando todas as rotas dentro do router
func Configure(r *mux.Router) *mux.Router {
	routes := routesUsers

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
