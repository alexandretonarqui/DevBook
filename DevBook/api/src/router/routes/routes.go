package routes

import (
	"api/src/middlewares"
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
	routes = append(routes, routesPublications...)

	for _, route := range routes {

		if route.RequestAuth {
			r.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Authentic(route.Function)),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}

	return r
}