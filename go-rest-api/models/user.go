package models

import (
	"errors"
	"go-rest-api/db"
	"go-rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

// Save user to the DATABASE
func (u User) Save() error {
	query := `
	INSERT INTO users(email, password)
	VALUES(?, ?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	u.ID = id

	return err
}

func (u User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return errors.New("credentials invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("credentials invalid")
	}

	return nil
}
