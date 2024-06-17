package src

import (
	"database/sql"
	"log/slog"

	_ "github.com/mattn/go-sqlite3"
)

type DataSources struct {
	DB  *sql.DB
	Log *slog.Logger
}

func initDB() (*sql.DB, error) {

	db, err := sql.Open("sqlite3", "forum.db")

	if err != nil {
		return nil, err
	}

	const createCredentials string = `
	CREATE TABLE IF NOT EXISTS credentials (
	id INTEGER NOT NULL PRIMARY KEY,
	username TEXT,
	email TEXT,
	password TEXT,
	sessionToken TEXT,
	expiry DATETIME
	);`

	const createPosts string = `
	CREATE TABLE IF NOT EXISTS posts (
	id INTEGER NOT NULL PRIMARY KEY,
	user_id INTEGER NOT NULL,
	label TEXT,
	text TEXT,
	category TEXT,
	time DATETIME,
	likes INTEGER NOT NULL,
	dislikes INTEGER NOT NULL,
	comments INTEGER NOT NULL,
	FOREIGN KEY (user_id) REFERENCES credentials (id) ON DELETE CASCADE ON UPDATE NO ACTION
	);`

	const createComments string = `
	CREATE TABLE IF NOT EXISTS comments (
	id INTEGER NOT NULL PRIMARY KEY,
	user_id INTEGER NOT NULL,
	post_id INTEGER NOT NULL,
	text TEXT,
	time DATETIME,
	likes INTEGER NOT NULL,
	dislikes INTEGER NOT NULL,
	FOREIGN KEY (user_id) REFERENCES credentials (id) ON DELETE CASCADE ON UPDATE NO ACTION,
	FOREIGN KEY (post_id) REFERENCES posts (id) ON DELETE CASCADE ON UPDATE NO ACTION
	);`

	const createLikes string = `
	CREATE TABLE IF NOT EXISTS likes (
	id INTEGER NOT NULL PRIMARY KEY,
	user_id INTEGER NOT NULL,
	post_id INTEGER NOT NULL,
	comment_id INTEGER,
	value INTEGER,
	FOREIGN KEY (user_id) REFERENCES credentials (id) ON DELETE CASCADE ON UPDATE NO ACTION,
	FOREIGN KEY (post_id) REFERENCES posts (id) ON DELETE CASCADE ON UPDATE NO ACTION,
	FOREIGN KEY (comment_id) REFERENCES comments (id) ON DELETE CASCADE ON UPDATE NO ACTION
	);`

	_, err = db.Exec(createCredentials)
	if err != nil {
		slog.Error("Error creating credentials table:", err)
		db.Close()
		return nil, err
	}
	_, err = db.Exec(createPosts)
	if err != nil {
		slog.Error("Error creating posts table:", err)
		db.Close()
		return nil, err
	}
	_, err = db.Exec(createComments)
	if err != nil {
		slog.Error("Error creating comments table:", err)
		db.Close()
		return nil, err
	}
	_, err = db.Exec(createLikes)
	if err != nil {
		slog.Error("Error creating likes table:", err)
		db.Close()
		return nil, err
	}
	slog.Info("Database initialized")
	return db, nil
}

func InitDataSources() *DataSources {
	db, err := initDB()
	if err != nil {
		slog.Error("Error initializing database:", err)
	}
	return &DataSources{
		DB:  db,
		Log: slog.Default(),
	}
}
