package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Routes todas as rotas da API
type Routes struct {
	URI 		string
	Method 		string
	Function 	func(http.ResponseWriter, *http.Request)
	RequestAuth bool
}

//Coloca todas as rotas dentro do Router
func Config(r *mux.Router) *mux.Router {
	routes := routesUsers
	routes = append(routes, loginRoute)

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}