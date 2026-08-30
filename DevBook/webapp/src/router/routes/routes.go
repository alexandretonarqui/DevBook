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
	routes = append(routes, usersRoutes...)
	routes = append(routes, homePageRouter)

	for _, route := range routes {
		router.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}