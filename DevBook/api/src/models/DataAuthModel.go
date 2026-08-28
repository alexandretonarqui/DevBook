package models

//DataAuth contém o Token e o ID do usuário autenticado
type DataAuth struct {
	ID string `json:"id"`
	Token string `json:"token"`
}