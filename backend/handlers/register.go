//real-time-forum/backend/handlers/register.go

// this file will contain the handler for the register route

package handlers

import (
    "database/sql"
    "net/http"
    "real-time-forum/backend/models"
    "log"
    "strconv"
)

func RegisterHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != "POST" {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }

        age, err := strconv.Atoi(r.FormValue("age"))
        if err != nil {
            http.Error(w, "Invalid age", http.StatusBadRequest)
            return
        }

        user := models.User{
            Nickname:  r.FormValue("nickname"),
            Age:       age,
            Gender:    r.FormValue("gender"),
            FirstName: r.FormValue("firstName"),
            LastName:  r.FormValue("lastName"),
            Email:     r.FormValue("email"),
            Password:  r.FormValue("password"),
        }

        err = models.RegisterUser(db, user)
        if err != nil {
            http.Error(w, "Error registering user", http.StatusInternalServerError)
            return
        }

        log.Println("User registered:", user.Email)
        http.Redirect(w, r, "/login", http.StatusSeeOther)
    }
}



