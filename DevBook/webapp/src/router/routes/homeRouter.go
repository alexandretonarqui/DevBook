package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var homePageRouter = Route{
	URI:		 "/home",
	Method:		 http.MethodGet,
	Function:	 controllers.LoadHomePage,
	RequestAuth: true,
}