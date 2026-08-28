package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

//ERRO respresenta a resposta de erro da API
type ErroAPI struct {
	Erro string `json:"erro"`
}

//JSON retorna uma resposta em formato JSON para a requisição
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if erro := json.NewEncoder(w).Encode(data); erro != nil {
		log.Fatal(erro)
	}
}

//TreatStatusCodeError trata as requisições com status code 400 ou superior
func TreatStatusCodeError(w http.ResponseWriter, r *http.Response) {
	var erro ErroAPI
	json.NewDecoder(r.Body).Decode(&erro)
	JSON(w, r.StatusCode, erro)
}