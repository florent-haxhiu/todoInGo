package database

import (
	"florent-haxhiu/todoInGo/internal/logger"
	"florent-haxhiu/todoInGo/internal/model"
	"fmt"
)

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

func GetUser(username string) (model.UserPassHashed, error) {
	var user model.UserPassHashed
	c := *createClient()

	err := c.Connection.QueryRow("SELECT * FROM Users WHERE username = ?", username).Scan(&user.Id, &user.Username, &user.Password)

	if err != nil {
		return user, fmt.Errorf("Error checking if user exists: %w", err)
	}

	return user, nil
}
