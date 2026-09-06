package models

import (
	"net/http"
	"time"
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

//SearchFullUser faz 4 requisições na API para montar o usuário
func SearchFullUser(userID, uint64, r http.Request) (User, error){
	userChannel := make(chan User)
	followersChannel := make(chan []User)
	followingChannel := make(chan []User)
	publicationsChannel := make(chan []Publication)

	go GetUserData(userChannel, userID, r)
	go GetFollowers(followersChannel,userID, r)
	go GetFollowing(followingChannel, userID, r)
	GetPublications(publicationsChannel, userID, r)

}

func GetUserData(channel <- chan User, userID uint64, r *http.Request) {

}

func GetFollowers(channel <- chan []User, userID uint64, r *http.Request) {

}

func GetFollowing(channel <- chan []User, userID uint64, r *http.Request) {

}

func GetPublications(channel <- chan []Publication, userID uint64, r *http.Request) {

}