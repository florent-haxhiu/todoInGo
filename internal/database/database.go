package database

import (
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"

	"florent-haxhiu/todoInGo/internal/model"
)

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
	row.Scan(&note.Id, &note.Title, &note.Body)

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

func UpdateNote(note model.Note, userId string) (model.Note, error) {
	c := *createClient()

	statement, err := c.Connection.Prepare("UPDATE Notes SET title=?, body=? WHERE userId = ?")
	if err != nil {
		return model.Note{}, err
	}
	statement.Exec(note.Title, note.Body, userId)
	return note, nil
}

func DeleteNote(noteId uuid.UUID, userId string) model.Note {
	var note model.Note
	// _ := *createClient()
	return note
}
