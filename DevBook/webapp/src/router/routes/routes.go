package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Route representa todas as rotas da Web Application
type Route struct {
	URI		    string
	Method	    string
	Function    func (http.ResponseWriter, *http.Request)
	RequestAuth bool
}

//Config coloca todas as rotas dentro do Router
func Config(router *mux.Router) *mux.Router {
	routes := loginRoutes

	for _, route := range routes {
		router.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return router
}