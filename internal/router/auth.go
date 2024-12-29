package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"florent-haxhiu/todoInGo/internal/model"
)

func Login(w http.ResponseWriter, r *http.Request) {}

func Register(w http.ResponseWriter, r *http.Request) {
	var user model.UserRegister

	fmt.Println(user)
	body := json.NewDecoder(r.Body)
	body.DisallowUnknownFields()

	err := body.Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	// TODO hash password
	hashedPass := model.UserPassHashed{
		Id:       user.Id,
		Username: user.Username,
		Password: user.Password,
	}

	key, err := generateToken(hashedPass)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	fmt.Println("Token: ", key)

	//err = db.SaveUserToDB(saltPassword(user))
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusTeapot)
	//	return
	//}
}

func generateToken(user model.UserPassHashed) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"iss": "issuer",
		"exp": time.Now().Add(time.Hour).Unix(),
		"data": map[string]any{
			"userId":   user.Id.String(),
			"username": user.Username,
		},
	})

	return token.SignedString([]byte("random"))
}

// func saltPassword(user model.UserRegister) model.UserPassHashed {
// 	var salt []byte
//
// 	key := argon.IDKey([]byte(user.Password), salt, 1, 64*1024, 4, 32)
//
// 	fmt.Println(user)
// 	fmt.Println(string(key))
//
// 	return model.UserPassHashed{
// 		Id:       user.Id,
// 		Username: user.Username,
// 		Password: string(key),
// 	}
// }

func getTokenPayload(token string) (model.TokenData, error) {
	var tokenPayload model.TokenData

	// TODO Get a better secret but it works for now
	payload, err := jwt.Parse(token, func(jwtTok *jwt.Token) (interface{}, error) {
		return []byte("random"), nil
	})

	claims := payload.Claims.(jwt.MapClaims)

	tokenPayload.UserId = uuid.MustParse(claims["userId"].(string))
	tokenPayload.Username = claims["username"].(string)

	if err != nil {
		return tokenPayload, err
	}

	return tokenPayload, nil
}

func verifyToken(token string) (model.TokenData, error) {
	newToken := strings.ReplaceAll(token, "Bearer ", "")

	tokenPayload, err := getTokenPayload(newToken)
	if err != nil {
		return tokenPayload, err
	}

	return tokenPayload, nil
}
