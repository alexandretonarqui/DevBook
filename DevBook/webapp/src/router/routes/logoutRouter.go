package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var logoutRouter = Route{
	URI: "/logout",
	Method: http.MethodGet,
	Function: controllers.Logout,
	RequestAuth: true,
}