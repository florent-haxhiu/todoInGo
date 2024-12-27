package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"florent-haxhiu/todoInGo/internal/model"
)

const (
	dbname = "todo.db"
)

func CreateClient() *model.Client {
	db, err := sql.Open("sqlite3", dbname)
	if err != nil {
		log.Fatal(err)
	}
	return &model.Client{
		Connection: db,
	}
}

func GetNote(id int) model.Note {
	var note model.Note
	c := *CreateClient()

	sql_item := c.Connection.QueryRow("SELECT ? FROM Notes", id)
	sql_item.Scan(&note)

	return note
}

func CreateNote(createdNote model.Note, userId int) (model.Note, error) {
	c := *CreateClient()

	statement, err := c.Connection.Prepare("INSERT INTO Notes (id, title, body, userId) VALUES (?, ?, ?, ?)")
	if err != nil {
		return model.Note{}, err
	}

	statement.Exec(createdNote.Id, createdNote.Title, createdNote.Body, createdNote.UserId)

	return createdNote, nil
}

func DeleteNote(noteId int, userId int) model.Note {
	var note model.Note
	// _ := *CreateClient()
	return note
}

func UpdateNote(note model.Note, userId int) (model.Note, error) {
	c := *CreateClient()

    statement, err := c.Connection.Prepare("UPDATE Notes SET title=?, body=?")
    if err != nil {
        return model.Note{}, err
    }
    statement.Exec(note.Title, note.Body)
	return note, nil
}
