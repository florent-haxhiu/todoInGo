package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"

	db "florent-haxhiu/todoInGo/internal/database"
	"florent-haxhiu/todoInGo/internal/model"
)

type AuthResources struct{}

func (ar AuthResources) Login(w http.ResponseWriter, r *http.Request) {}

func (ar AuthResources) Register(w http.ResponseWriter, r *http.Request) {
	var user model.UserRegister

	fmt.Println(user)
	body := json.NewDecoder(r.Body)
	body.DisallowUnknownFields()

	err := body.Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	done, err := db.SaveUserToDB(ar.saltPassword(user))
	if err != nil {
		http.Error(w, err.Error(), http.StatusTeapot)
		return
	}
}

func (ar AuthResources) generateToken(user model.UserPassHashed) *jwt.Token {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"username": user.Username,
		"password": user.Password,
	})

	return token
}

func (ar AuthResources) saltPassword(user model.UserRegister) model.UserPassHashed {
	var hashed model.UserPassHashed

	print(user)

	return hashed
}

func (ar AuthResources) unhashPassword() {}
