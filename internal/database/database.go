package database

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"

	"florent-haxhiu/todoInGo/internal/model"
)

const (
	dbname = "todo.db"
)

func createClient() *model.Client {
	db, err := sql.Open("sqlite3", dbname)
	if err != nil {
		log.Fatal(err)
	}
	return &model.Client{
		Connection: db,
	}
}

func GetAllNotes() ([]model.Note, error) {
	var notes []model.Note
	c := *createClient()

	rows, err := c.Connection.Query("SELECT * FROM Notes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var note model.Note

		if err := rows.Scan(&note.Id, &note.Title, &note.Body, &note.UserId); err != nil {
			return notes, err
		}

		notes = append(notes, note)
	}

	if err = rows.Err(); err != nil {
		return notes, err
	}

	return notes, nil
}

func GetNote(id uuid.UUID, userId uuid.UUID) model.Note {
	var note model.Note
	c := *createClient()

	row := c.Connection.QueryRow("SELECT ? FROM Notes WHERE userId = ?", id, userId)
	row.Scan(&note)

	return note
}

func CreateNote(createdNote model.Note) (model.Note, error) {
	c := *createClient()

	statement, err := c.Connection.Prepare("INSERT INTO Notes (id, title, body, userId) VALUES (?, ?, ?, ?)")
	if err != nil {
		return model.Note{}, err
	}

	statement.Exec(createdNote.Id.String(), createdNote.Title, createdNote.Body, createdNote.UserId.String())

	return createdNote, nil
}

func DeleteNote(noteId uuid.UUID, userId uuid.UUID) model.Note {
	var note model.Note
	// _ := *createClient()
	return note
}

func UpdateNote(note model.Note, userId uuid.UUID) (model.Note, error) {
	c := *createClient()

	statement, err := c.Connection.Prepare("UPDATE Notes SET title=?, body=?")
	if err != nil {
		return model.Note{}, err
	}
	statement.Exec(note.Title, note.Body)
	return note, nil
}

func SaveUserToDB(user model.UserPassHashed) error {
	c := *createClient()

	statement, err := c.Connection.Prepare("INSERT INTO Users (id, username, password) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}

	statement.Exec(user.Id, user.Username, user.Password)

	return nil
}

func GetUserFromDB(user model.UserPassHashed) (model.UserPassHashed, error) {
	c := *createClient()

	statement, err := c.Connection.Prepare("INSERT INTO Users (id, username, password) VALUES (?, ?, ?)")
	if err != nil {
		return user, err
	}
	statement.Exec(user.Id, user.Username, user.Password)

	return user, nil
}
