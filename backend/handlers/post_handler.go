//real-time-forum/backend/handlers/post_handler.go

// This file will contain the handler for the post route.

package handlers

import (
    "database/sql"
    "encoding/json"
    "net/http"
    "real-time-forum/backend/models"
    "real-time-forum/backend/utils"
)

func PostsHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" {
            userID := utils.GetUserIDFromSession(r)
            if userID == 0 {
                http.Error(w, "Unauthorized", http.StatusUnauthorized)
                return
            }

            post := models.Post{
                UserID:   userID,
                Title:    r.FormValue("title"),
                Content:  r.FormValue("content"),
            }

            err := models.CreatePost(db, post)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }

            w.WriteHeader(http.StatusCreated)
        } else if r.Method == "GET" {
            if !utils.IsAuthenticated(r) {
                http.Error(w, "Unauthorized", http.StatusUnauthorized)
                return
            }

            posts, err := models.GetPosts(db)
            if err != nil {
                http.Error(w, "Error getting posts from database: "+err.Error(), http.StatusInternalServerError)
                return
            }

            jsonData, err := json.Marshal(posts)
            if err != nil {
                http.Error(w, "Error marshaling posts to JSON: "+err.Error(), http.StatusInternalServerError)
                return
            }

            w.Header().Set("Content-Type", "application/json")
            w.Write(jsonData)
        } else {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    }
}

