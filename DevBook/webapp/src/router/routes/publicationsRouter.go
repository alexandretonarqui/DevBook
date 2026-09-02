package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var publicationsRoutes = []Route {
	{
		URI: "/publications",
		Method: http.MethodPost,
		Function: controllers.CreatePublication,
		RequestAuth: true,
	},
}