package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var publicationsRoutes = []Route{
	{
		URI:         "/publications",
		Method:      http.MethodPost,
		Function:    controllers.CreatePublication,
		RequestAuth: true,
	},
	{
		URI:         "/publications/{publicationID}/like",
		Method:      http.MethodPost,
		Function:    controllers.LikePublication,
		RequestAuth: true,
	},
	{
		URI:         "/publications/{publicationID}/unlike",
		Method:      http.MethodPost,
		Function:    controllers.UnLikePublication,
		RequestAuth: true,
	},
	{
		URI:         "/publications/{publicationID}/update",
		Method:      http.MethodGet,
		Function:    controllers.LoadUpdatePublicationPage,
		RequestAuth: true,
	},
	{
		URI:         "/publications/{publicationID}",
		Method:      http.MethodPut,
		Function:    controllers.UpdatePublication,
		RequestAuth: true,
	},
	{
		URI:         "/publications/{publicationID}",
		Method:      http.MethodDelete,
		Function:    controllers.DeletePublication,
		RequestAuth: true,
	},
}
