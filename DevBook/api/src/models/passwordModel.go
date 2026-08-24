package models

//Representa o formata da requisição de alteração de senha
type Password struct {
	New string `json:"new"`
	Current string `json:"current"`
}