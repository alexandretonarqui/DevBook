package requests

import (
	"io"
	"net/http"
	"webapp/src/cookies"
)

//MakeAuthRequest é utilizado para colocar o token na requisição
func MakeAuthRequest(r *http.Request, method, url string, datas io.Reader) (*http.Response, error) {
	request, erro := http.NewRequest(method, url, datas)
	if erro != nil {
		return nil, erro
	}

	cookie, _ := cookies.Read(r)
	request.Header.Add("Authorization", "Bearer "+cookie["token"])

	client := &http.Client{}
	response, erro := client.Do(request)
	if erro != nil {
		return nil, erro
	}

	return response, nil
}