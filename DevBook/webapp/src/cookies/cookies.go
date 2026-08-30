package cookies

import (
	"net/http"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

//Configuration utiliza as varáveis de ambiente para a criação do SecureCookie
func Configuration() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

//Save registra as informações de autenticação
func Save(w http.ResponseWriter, ID, token string) error {
	data := map[string] string {
		"id": ID,
		"token": token,
	}

	codedData, erro := s.Encode("datas", data)
	if erro != nil {
		return erro
	}

	http.SetCookie(w, &http.Cookie{
		Name: "datas",
		Value: codedData,
		Path: "/",
		HttpOnly: true,
	})

	return nil
}

//Read retorna os valores armazenados no cookie
func Read(r *http.Request) (map[string]string, error) {
	cookie, erro := r.Cookie("datas")
	if erro != nil {
		return nil, erro
	}

	values := make(map[string]string)
	if erro = s.Decode("datas", cookie.Value, &values); erro != nil {
		return nil, erro
	}

	return values, nil
}