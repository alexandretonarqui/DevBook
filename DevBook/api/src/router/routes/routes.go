package routes

import "net/http"

//Routes todas as rotas da API
type Routes struct {
	URI 		string
	Method 		string
	Function 	func(http.ResponseWriter, *http.Request)
	RequestAuth bool
}