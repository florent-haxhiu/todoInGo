package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"

	"florent-haxhiu/todoInGo/internal/logger"
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

func GetAllNotes(userId string) ([]model.Note, error) {
	var notes []model.Note
	c := *createClient()

	rows, err := c.Connection.Query("SELECT * FROM Notes WHERE userId = ?", userId)
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

func GetNote(id uuid.UUID, userId string) model.Note {
	var note model.Note
	c := *createClient()

	row := c.Connection.QueryRow("SELECT ? FROM Notes WHERE userId = ?", id, userId)
	row.Scan(&note)

	return note
}

func CreateNote(createdNote model.Note, userId string) (model.Note, error) {
	c := *createClient()

	statement, err := c.Connection.Prepare("INSERT INTO Notes (id, title, body, userId) VALUES (?, ?, ?, ?)")
	if err != nil {
		return model.Note{}, err
	}

	statement.Exec(createdNote.Id.String(), createdNote.Title, createdNote.Body, userId)

	return createdNote, nil
}

func DeleteNote(noteId uuid.UUID, userId string) model.Note {
	var note model.Note
	// _ := *createClient()
	return note
}

func UpdateNote(note model.Note, userId string) (model.Note, error) {
	c := *createClient()

	statement, err := c.Connection.Prepare("UPDATE Notes SET title=?, body=? WHERE userId = ?")
	if err != nil {
		return model.Note{}, err
	}
	statement.Exec(note.Title, note.Body, userId)
	return note, nil
}

func SaveUserToDB(user model.UserPassHashed) error {
	c := *createClient()

	logger.DebugMsg("User to be saved in db", "user", user)

	statement, err := c.Connection.Prepare("INSERT INTO Users (id, username, password) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}

	statement.Exec(user.Id, user.Username, user.Password)

	return nil
}

func UserExists(username string) (bool, error) {
	var count int 
	c := *createClient()

	err := c.Connection.QueryRow("SELECT COUNT(*) FROM Users WHERE username = ?", username).Scan(&count)

	if err != nil {
		return false, fmt.Errorf("Error checking if user exists: %w", err)
	}

	return count > 0, nil
}
