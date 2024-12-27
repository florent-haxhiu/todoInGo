package model

import "database/sql"

type Client struct {
	Connection *sql.DB
}

type Note struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId string `json:"userId"`
}

type Response struct {
	Message     string `json:"message"`
	CreatedNote Note   `json:"createdNote"`
}

type ResponseNotes struct {
	Data []Note `json:"data"`
}
