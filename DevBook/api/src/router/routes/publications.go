package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesPublications = []Routes {
	{
		URI:		 "/publications",
		Method:		 http.MethodPost,
		Function:	 controllers.CreatePublication,
		RequestAuth: true,
	},
	{
		URI:		 "/publications",
		Method:		 http.MethodGet,
		Function:	 controllers.FindPublications,
		RequestAuth: true,
	},
	{
		URI:		 "/publications/{publicationID}",
		Method:		 http.MethodGet,
		Function:	 controllers.FindPublication,
		RequestAuth: true,
	},
	{
		URI:		 "/publications/{publicationID}",
		Method:		 http.MethodPut,
		Function:	 controllers.UpdatePublication,
		RequestAuth: true,
	},
	{
		URI:		 "/publications/{publicationID}",
		Method:		 http.MethodDelete,
		Function:	 controllers.DeletePublication,
		RequestAuth: true,
	},
	{
		URI:		 "/users/{userID}/publications",
		Method:		 http.MethodGet,
		Function:	 controllers.GetPublicationsByUser,
		RequestAuth: true,
	},
	{
		URI: "/publications/{publicationID}/like",
		Method: http.MethodPost,
		Function: controllers.Like,
		RequestAuth: true,
	},
}