package models

import (
	"errors"

	"github.com/events/db"
	"github.com/events/utils"
)

type User struct {
	ID int64
	Email string `binding:"required"`
	Password string `binding:"required"`
}

func (s *User) Save() error {
	query := "INSERT INTO users(email, password) VALUES(?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(s.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(s.Email, hashedPassword)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	s.ID = userId
	return err
}

func (s *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, s.Email)

	var retrievedPassword string
	err := row.Scan(&s.ID, &retrievedPassword)

	if err != nil {
		return err
	}

	passwordIsValid := utils.CheckPasswordHash(s.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("CREDENTIALS INVALID")
	}

	return nil
}