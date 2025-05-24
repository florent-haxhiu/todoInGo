package router

import (
	"os"
	"reflect"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"florent-haxhiu/todoInGo/internal/model"
)

// Reflection-based approach
func isStructEmpty(v interface{}) bool {
	return reflect.DeepEqual(v, reflect.Zero(reflect.TypeOf(v)).Interface())
}

func saltPassword(user model.User) (model.UserPassHashed, error) {
	key, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.UserPassHashed{}, err
	}

	return model.UserPassHashed{
		Id:       user.Id,
		Username: user.Username,
		Password: string(key),
	}, nil
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

func getTokenPayload(token string) (model.TokenData, error) {
	var tokenPayload model.TokenData

	// TODO Get a better secret but it works for now
	payload, err := jwt.Parse(token, func(jwtTok *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNING_KEY")), nil
	})
	if err != nil {
		return tokenPayload, err
	}

	claims := payload.Claims.(jwt.MapClaims)

	data := claims["data"].(map[string]interface{})

	tokenPayload.UserId = uuid.MustParse(data["userId"].(string))
	tokenPayload.Username = data["username"].(string)

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
