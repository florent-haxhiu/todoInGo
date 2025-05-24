package database

import (
	"database/sql"
	"florent-haxhiu/todoInGo/internal/model"
	"log"
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
