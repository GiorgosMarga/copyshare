package models

import (
	"database/sql"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	password     string
	Username     string `json:"username"`
	isRestricted bool
}

type UserModel struct {
	DB *sql.DB
}

func (u *UserModel) Register(email, password, username string) (int, error) {
	stmt := `INSERT INTO user(email, password, username,createdAt, updatedAt, isRestricted) VALUES(?,?,?,UTC_TIMESTAMP(),UTC_TIMESTAMP(),false)`
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return 0, err
	}
	res, err := u.DB.Exec(stmt, email, string(hashedPassword), username)
	if err != nil {
		code := strings.Split(err.Error(), " ")[1]
		if code == "1062" {
			return 0, ErrDuplicateEmail
		}
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (u *UserModel) Login(email, password string) (*User, error) {
	user := &User{}
	stmt := `SELECT id,email, username, password, isRestricted FROM user WHERE email=?`
	err := u.DB.QueryRow(stmt, email).Scan(&user.ID, &user.Email, &user.Username, &user.password, &user.isRestricted)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecond
		}
		return nil, err
	}
	if user.isRestricted {
		return nil, ErrRestrictedAcc
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.password), []byte(password))
	if err != nil {
		return nil, ErrInvalidCredentials
	}
	return user, nil
}

func (u *UserModel) Exists(id int) (bool, error) {
	var exists bool
	stmt := "SELECT EXISTS(SELECT true FROM user WHERE id=? AND isRestricted=false)"
	err := u.DB.QueryRow(stmt, id).Scan(&exists)
	return exists, err
}

func (u *UserModel) WhoAmI(id int) (string, error) {
	var username string
	stmt := "SELECT username FROM user WHERE id=?"
	err := u.DB.QueryRow(stmt, id).Scan(&username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", ErrNoRecond
		}
		return "", err
	}
	return username, nil

}
