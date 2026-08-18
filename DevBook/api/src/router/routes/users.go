package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesUsers = []Routes {
	{
		URI: 		 "/users",
		Method: 	 http.MethodPost,
		Function: 	 controllers.CreateUser,
		RequestAuth: false,
	},

	{
		URI:		 "/users",
		Method:	 	 http.MethodGet,
		Function: 	 controllers.FindAll,
		RequestAuth: false,
	},

	{
		URI:		 "/users/{userID}",
		Method:		 http.MethodGet,
		Function:	 controllers.FindByID,
		RequestAuth: false,
	},

	{
		URI:		 "/users/{userID}",
		Method:		 http.MethodPut,
		Function:	 controllers.UpdateUser,
		RequestAuth: false,
	},

	{
		URI:		 "/users/{userID}",
		Method:		 http.MethodDelete,
		Function:	 controllers.DeleteUser,
		RequestAuth: false,
	},
}