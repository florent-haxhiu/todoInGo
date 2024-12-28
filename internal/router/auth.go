package router

import (
	"encoding/json"
	"fmt"
    "strings"
	"net/http"

	"github.com/golang-jwt/jwt/v5"

    argon "golang.org/x/crypto/argon2"
    "crypto/rsa"
    "crypto/rand"

	//db "florent-haxhiu/todoInGo/internal/database"
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

    hashed := saltPassword(user)

    key, err := generateToken(hashed)
    if err != nil {
        http.Error(w, err.Error(), 422)
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
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
        "userId": user.Id,
		"username": user.Username,
	})

    secret, err := generateKey()
    if err != nil {
        return "", err
    }

    return token.SignedString(secret)
}

func generateKey() (*rsa.PrivateKey, error) {
    key, err := rsa.GenerateKey(rand.Reader, 2048)
    if err != nil {
        return nil, err
    }

    return key, nil
}

func saltPassword(user model.UserRegister) model.UserPassHashed {
    var salt []byte

    key := argon.IDKey([]byte(user.Password), salt, 1, 64*1024, 4, 32)

	fmt.Println(user)
    fmt.Println(key)


	return model.UserPassHashed{
        Id: user.Id,
        Username: user.Username,
        Password: string(key),
    }
}

func verifyToken(token string) (string) {
    fmt.Println(token)
    return strings.ReplaceAll(token, "Bearer ", "")
}
