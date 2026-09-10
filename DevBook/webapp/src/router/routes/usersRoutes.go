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
	{
		URI: 		 "/search-users",
		Method: 	 http.MethodGet,
		Function: 	 controllers.LoadUsersPage,
		RequestAuth: true,
	},
	{
		URI: 		 "/users/{userID}",
		Method: 	 http.MethodGet,
		Function: 	 controllers.LoadUsersProfile,
		RequestAuth: true,
	},
	{
		URI: 		 "/users/{userID}/unfollow",
		Method: 	 http.MethodPost,
		Function: 	 controllers.UnfollowUser,
		RequestAuth: true,
	},
	{
		URI: 		 "/users/{userID}/follow",
		Method: 	 http.MethodPost,
		Function: 	 controllers.FollowUser,
		RequestAuth: true,
	},
	{
		URI: 		 "/profile",
		Method: 	 http.MethodGet,
		Function: 	 controllers.LoadUserProfileLogged,
		RequestAuth: true,
	},
	{
		URI: 		 "/edit-user",
		Method: 	 http.MethodGet,
		Function: 	 controllers.LoadUserEditProfile,
		RequestAuth: true,
	},
	{
		URI: 		 "/edit-user",
		Method: 	 http.MethodPut,
		Function: 	 controllers.EditUser,
		RequestAuth: true,
	},
	{
		URI: 		 "/updatepassword",
		Method: 	 http.MethodGet,
		Function: 	 controllers.LoadUpdatePasswordPage,
		RequestAuth: true,
	},
	{
		URI: 		 "/updatepassword",
		Method: 	 http.MethodPost,
		Function: 	 controllers.UpdatePassword,
		RequestAuth: true,
	},
}