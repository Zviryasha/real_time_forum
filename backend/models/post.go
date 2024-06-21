// real-time-forum/backend/models/post.go

// This file will contain the model for the post entity.

package models

import (
    "database/sql"
    "time"
)

type Post struct {
    ID        int       `json:"id"`
    UserID    int       `json:"userId"`
    Title     string    `json:"title"`
    Content   string    `json:"content"`
    CreatedAt time.Time `json:"createdAt"`
}

func CreatePost(db *sql.DB, post Post) error {
    stmt, err := db.Prepare(`
        INSERT INTO posts (userId, title, content, createdAt)
        VALUES (?, ?, ?, ?)
    `)
    if err != nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(post.UserID, post.Title, post.Content, time.Now())
    return err
}

func GetPosts(db *sql.DB) ([]Post, error) {
    rows, err := db.Query("SELECT id, userId, title, content, createdAt FROM posts ORDER BY createdAt DESC")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var posts []Post
    for rows.Next() {
        var post Post
        err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CreatedAt)
        if err != nil {
            return nil, err
        }
        posts = append(posts, post)
    }
    return posts, nil
}
