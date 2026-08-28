package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var usersRoutes = []Route {
	{
		URI: 		 "/create-user",
		Method: 	 http.MethodGet,
		Function: 	 controllers.LoadUserSubmitPage,
		RequestAuth: false,
	},
	{
		URI: 		 "/users",
		Method: 	 http.MethodPost,
		Function: 	 controllers.CreateUser,
		RequestAuth: false,
	},
}