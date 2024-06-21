// real-time-forum/backend/utils/utils.go

// This file will contain the utility functions for the application.

package utils

import (
    "net/http"
    "github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secret-key"))

func IsAuthenticated(r *http.Request) bool {
    session, _ := store.Get(r, "session-name")
    auth, ok := session.Values["authenticated"].(bool)
    return ok && auth
}

func GetUserIDFromSession(r *http.Request) int {
    session, _ := store.Get(r, "session-name")
    userID, ok := session.Values["user_id"].(int)
    if !ok {
        return 0
    }
    return userID
}

func ClearSession(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "session-name")
    session.Values["authenticated"] = false
    session.Values["user_id"] = 0
    session.Save(r, w)
}

func SetUserSession(w http.ResponseWriter, r *http.Request, userID int) {
    session, _ := store.Get(r, "session-name")
    session.Values["authenticated"] = true
    session.Values["user_id"] = userID
    session.Save(r, w)
}

