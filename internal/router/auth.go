package router

import (
	"encoding/json"
	"net/http"
	"time"

	db "florent-haxhiu/todoInGo/internal/database"
	"florent-haxhiu/todoInGo/internal/logger"
	"florent-haxhiu/todoInGo/internal/model"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user model.User

	body := json.NewDecoder(r.Body)
	body.DisallowUnknownFields()

	err := body.Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	exists, err := db.UserExists(user.Username)

	if !exists {
		err_resp, _ := json.Marshal(model.ErrorResponse{Message: "User does not exist with that username", Status: http.StatusUnauthorized})
		http.Error(w, string(err_resp), http.StatusUnauthorized)
		return
	}

	user_in_db, err := db.GetUser(user.Username)

	logger.InfoMsg("User from db", user_in_db)

	err = verifyPassword(user.Password, user_in_db.Password)

	if err != nil {
		err_resp, _ := json.Marshal(model.ErrorResponse{Message: "Details are incorrect", Status: http.StatusUnauthorized})
		http.Error(w, string(err_resp), http.StatusUnauthorized)
		return
	}

	expDate := time.Now().Add(time.Hour).Unix()

	key, err := generateToken(user_in_db, expDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
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

	w.WriteHeader(200)
	w.Write(b)
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user model.User

	body := json.NewDecoder(r.Body)
	body.DisallowUnknownFields()

	err := body.Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	exists, err := db.UserExists(user.Username)

	logger.InfoMsg("User in db", "user", exists)

	if err != nil {
		err_resp, _ := json.Marshal(model.ErrorResponse{Message: "Error checking user", Status: http.StatusConflict})
		http.Error(w, string(err_resp), http.StatusConflict)
		logger.ErrorMsg("User existence check failed", "error", err)
		return
	}

	if exists {
		err_resp, _ := json.Marshal(model.ErrorResponse{Message: "Username already exists with that name", Status: http.StatusConflict})
		http.Error(w, string(err_resp), http.StatusConflict)
		return
	}

	hashedPassUserModel, err := saltPassword(user)
	if err != nil {
		err_resp, _ := json.Marshal(model.ErrorResponse{Message: err.Error(), Status: http.StatusConflict})
		http.Error(w, string(err_resp), http.StatusConflict)
		return
	}

	err = db.SaveUserToDB(hashedPassUserModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusTeapot)
		return
	}

	resp := model.UserRegisterResponse{
		Message: "You have successfully registered",
	}

	b, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusTeapot)
		return
	}

	w.WriteHeader(201)
	w.Write(b)
}

func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	// var user model.User

	k := model.UserId("userId")

	logger.InfoMsg("What is the user id", r.Context().Value(k))

	// user_in_db, err := db.GetUser(user.Username)
}
