// real-time-forum/backend/models/user.go

// This file will contain the model for the user.

package models

import (
    "database/sql"
    "errors"
)

type User struct {
    ID        int
    Nickname  string
    Age       int
    Gender    string
    FirstName string
    LastName  string
    Email     string
    Password  string
}

func ValidateUser(db *sql.DB, email, password string) (User, error) {
    var user User
    query := "SELECT id, password FROM users WHERE email = ?"
    err := db.QueryRow(query, email).Scan(&user.ID, &user.Password)
    if err != nil {
        return user, errors.New("invalid credentials")
    }

    if user.Password != password {
        return user, errors.New("invalid credentials")
    }

    return user, nil
}

func RegisterUser(db *sql.DB, user User) error {
    query := `
    INSERT INTO users (nickname, age, gender, firstName, lastName, email, password)
    VALUES (?, ?, ?, ?, ?, ?, ?)
    `
    _, err := db.Exec(query, user.Nickname, user.Age, user.Gender, user.FirstName, user.LastName, user.Email, user.Password)
    return err
}
