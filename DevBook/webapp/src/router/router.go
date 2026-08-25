package router

import "github.com/gorilla/mux"

//Generate retorna um router com todas as rotas configuradas
func Generate() *mux.Router {
	return mux.NewRouter()
}