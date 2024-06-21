// real-time-forum/backend/db/init_db.go

// This file will contain the initialization of the database schema.

package db

import (
    "database/sql"
    "log"
)

func InitDB(db *sql.DB) {
    query := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        nickname TEXT NOT NULL,
        age INTEGER NOT NULL,
        gender TEXT NOT NULL,
        firstName TEXT NOT NULL,
        lastName TEXT NOT NULL,
        email TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL
    );
    CREATE TABLE IF NOT EXISTS posts (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        userId INTEGER NOT NULL,
        title TEXT NOT NULL,
        content TEXT NOT NULL,
        createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY(userId) REFERENCES users(id)
    );
    `
    _, err := db.Exec(query)
    if err != nil {
        log.Fatal("Error initializing database:", err)
    }
}






