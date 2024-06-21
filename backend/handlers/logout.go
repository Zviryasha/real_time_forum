//real-time-forum/backend/handlers/logout.go

// This file will contain the handler for the logout route.

package handlers

import (
    "net/http"
    "real-time-forum/backend/utils"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
    utils.ClearSession(w, r)
    w.WriteHeader(http.StatusOK)
}
