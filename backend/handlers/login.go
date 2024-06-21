//real-time-forum/backend/handlers/login.go

// this file will contain the handler for the login route

package handlers

import (
    "database/sql"
    "net/http"
    "real-time-forum/backend/models"
    "real-time-forum/backend/utils"
    "log"
)

func LoginHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != "POST" {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }

        email := r.FormValue("email")
        password := r.FormValue("password")

        user, err := models.ValidateUser(db, email, password)
        if err != nil {
            http.Error(w, "Invalid credentials", http.StatusUnauthorized)
            return
        }

        utils.SetUserSession(w, r, user.ID)
        log.Println("User logged in:", user.ID)
        w.WriteHeader(http.StatusOK)
    }
}


