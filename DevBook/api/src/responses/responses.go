package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

//Retorna uma resposta em Json para a requisição
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(statusCode)

	if erro := json.NewEncoder(w).Encode(data); erro != nil {
		log.Fatal(erro)
	}
}

//Retorna um erro em Json
func Erro(w http.ResponseWriter, statusCode int, erro error)  {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}