package route

import "net/http"

// Representa todas as rotas da API
type Rota struct {
	URI                    string
	Method                 string
	Function               func(w http.ResponseWriter, r *http.Request)
	RequiresAuthentication bool
}
