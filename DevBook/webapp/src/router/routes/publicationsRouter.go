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
	{
		URI: "/publications/{publicationID}/like",
		Method: http.MethodPost,
		Function: controllers.LikePublication,
		RequestAuth: true,
	},
	{
		URI: "/publications/{publicationID}/unlike",
		Method: http.MethodPost,
		Function: controllers.UnLikePublication,
		RequestAuth: true,
	},
	{
		URI: "/publications/{publicationID}/edit",
		Method: http.MethodGet,
		Function: controllers.LoadEditPublicationPage,
		RequestAuth: true,
	},
}