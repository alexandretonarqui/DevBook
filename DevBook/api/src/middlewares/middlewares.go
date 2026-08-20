package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

//Imprime no terminal as informações da requisição
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

//Verifica se o usuário fazendo a requisição, ele está autenticado
func Authentic(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Autenticando...")
		next(w, r)
	}
}