package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbname = "todo"
)

type Client struct {
	Connection *sql.DB
}

type Note struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Body string `json:"body"`
}

type NoteMethods interface {
	GetNote(id int) Note
	CreateNote(createdNote Note, userId int) Note
	DeleteNote(noteId int, userId int) Note
	UpdateNote(note Note, userId int) Note
}

func CreateClient() *Client {
	db, err := sql.Open("sqlite3", dbname)
	if err != nil {
		log.Fatal(err)
	}
	return &Client{
		Connection: db,
	}
}

func (c *Client) GetNote(id int) Note {
	var note Note

	sql_item := c.Connection.QueryRow("SELECT $1 FROM Inventory", id)
	sql_item.Scan(&note)

	return note
}

func (c *Client) CreateNote(createdNote Note, userId int) Note {
	return createdNote
}

func (c *Client) DeleteNote(noteId int, userId int) Note {
	var note Note
	return note
}

func (c *Client) UpdateNote(note Note, userId int) Note {
	return note
}
