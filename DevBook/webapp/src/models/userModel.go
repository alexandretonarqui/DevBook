package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"webapp/src/config"
	"webapp/src/requests"
)

// User representa uma pessoa utilizando a rede social
type User struct {
	ID           uint64        `json:"id"`
	Name         string        `json:"name"`
	Email        string        `json:"email"`
	Nick         string        `json:"nick"`
	CreatedAt    time.Time     `json:"createdAt"`
	Followers    []User        `json:"followers"`
	Following    []User        `json:"following"`
	Publications []Publication `json:"publications"`
}

// SearchFullUser faz 4 requisições na API para montar o usuário
func SearchFullUser(userID, uint64, r http.Request) (User, error) {
	userChannel := make(chan User)
	followersChannel := make(chan []User)
	followingChannel := make(chan []User)
	publicationsChannel := make(chan []Publication)

	go GetUserData(userChannel, userID, r)
	go GetFollowers(followersChannel, userID, r)
	go GetFollowing(followingChannel, userID, r)
	GetPublications(publicationsChannel, userID, r)

}

// GetUserData chama a API para buscar os dados base do usuário
func GetUserData(channel chan<- User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d", config.APIURL, userID)
	response, erro := requests.MakeAuthRequest(r, http.MethodGet, url, nil)
	if erro != nil {
		channel <- User{}
		return
	}
	defer response.Body.Close()

	var user User
	if erro = json.NewDecoder(response.Body).Decode(&user); erro != nil {
		channel <- User{}
		return
	}

	channel <- user
}

// GetFollowers chama a API para buscar os seguidores do usuário
func GetFollowers(channel chan<- []User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/followers", config.APIURL, userID)
	response, erro := requests.MakeAuthRequest(r, http.MethodGet, url, nil)
	if erro != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var followers []User
	if erro = json.NewDecoder(response.Body).Decode(&followers); erro != nil {
		channel <- nil
		return
	}

	channel <- followers
}

// GetFollowing chama a API para buscar os usuários seguidos por un usuário
func GetFollowing(channel chan<- []User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/following", config.APIURL, userID)
	response, erro := requests.MakeAuthRequest(r, http.MethodGet, url, nil)
	if erro != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var following []User
	if erro = json.NewDecoder(response.Body).Decode(&following); erro != nil {
		channel <- nil
		return
	}

	channel <- following
}

//GetPublications chama a API para buscar as publicações de um usuário
func GetPublications(channel chan<- []Publication, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/publications", config.APIURL, userID)
	response, erro := requests.MakeAuthRequest(r, http.MethodGet, url, nil)
	if erro != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var publications []Publication
	if erro = json.NewDecoder(response.Body).Decode(&publications); erro != nil {
		channel <- nil
		return
	}

	channel <- publications
}
