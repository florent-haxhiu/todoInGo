package router

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	db "florent-haxhiu/todoInGo/internal/database"
	"florent-haxhiu/todoInGo/internal/model"
)

func Login(w http.ResponseWriter, r *http.Request) {}

func Register(w http.ResponseWriter, r *http.Request) {
	var user model.UserRegister

	body := json.NewDecoder(r.Body)
	body.DisallowUnknownFields()

	err := body.Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	userInDB, _ := db.GetUserFromDB(user.Username)

	if userInDB != (model.UserPassHashed{}) {
		http.Error(w, "Username already exists with that name", http.StatusConflict)
		return
	}

	hashedPassUserModel, err := saltPassword(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}

	expDate := time.Now().Add(time.Hour).Unix()

	key, err := generateToken(hashedPassUserModel, expDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	err = db.SaveUserToDB(hashedPassUserModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusTeapot)
		return
	}

	resp := model.UserLoginResponse{
		Token:      key,
		Expiration: expDate,
	}

	b, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusTeapot)
		return
	}

	w.WriteHeader(201)
	w.Write(b)
}

func generateToken(user model.UserPassHashed, expDate int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"iss": "issuer",
		"exp": expDate,
		"data": map[string]any{
			"userId":   user.Id.String(),
			"username": user.Username,
		},
	})

	signingKey := os.Getenv("SIGNING_KEY")

	return token.SignedString([]byte(signingKey))
}

func saltPassword(user model.UserRegister) (model.UserPassHashed, error) {
	key, err := bcrypt.GenerateFromPassword([]byte(user.Password), 0)
	if err != nil {
		return model.UserPassHashed{}, err
	}

	return model.UserPassHashed{
		Id:       user.Id,
		Username: user.Username,
		Password: string(key),
	}, nil
}

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

func verifyPassword(password_from_user string, password_in_db string) error {
	return bcrypt.CompareHashAndPassword([]byte(password_in_db), []byte(password_from_user))
}
