package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type UserId string

type Client struct {
	Connection *sql.DB
}

type Note struct {
	Id     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	Body   string    `json:"body"`
	UserId uuid.UUID `json:"userId"`
}

type User struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string
}

type UserPassHashed struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string
}

type TokenData struct {
	UserId   uuid.UUID `json:"userId"`
	Username string    `json:"username"`
}

type Response struct {
	Message string `json:"message"`
	Note    Note   `json:"createdNote"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Status int `json:"statusCode"`
}

type UserLoginResponse struct {
	Token      string `json:"token"`
	Expiration int64  `json:"expiration"`
}

type UserRegisterResponse struct {
	Message string `json:"message"`
}

type ResponseNotes struct {
	Data []Note `json:"data"`
}
