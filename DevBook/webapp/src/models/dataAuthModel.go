package models

//DataAuth contém o ID e o Token do usuário autenticado
type DataAuth struct {
	ID string 	 `json:"id"`
	Token string `json:"token"`
}