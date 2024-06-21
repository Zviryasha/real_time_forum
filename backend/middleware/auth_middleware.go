// real-time-forum/backend/middleware/auth_middleware.go

// This file contains the middleware for authenticating the user. 
// The middleware will be used in the main handler to ensure that only authenticated users can access the main page.
// The middleware will be used in the login handler to ensure that only unauthenticated users can log in.
// The middleware will be used in the logout handler to ensure that only authenticated users can log out.
// The middleware will be used in the post handler to ensure that only authenticated users can post messages.
// The middleware will be used in the WebSocket handler to ensure that only authenticated users can connect to the WebSocket.
// The middleware will be used in the register handler to ensure that only unauthenticated users can register.

// backend/middleware/auth_middleware.go

package middleware

import (
    "net/http"
    "real-time-forum/backend/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if !utils.IsAuthenticated(r) {
            http.Redirect(w, r, "/login", http.StatusSeeOther)
            return
        }
        next.ServeHTTP(w, r)
    })
}
