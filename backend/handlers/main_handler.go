//real-time-forum/backend/handlers/main_handler.go

// This file will contain the handler for the main route.

package handlers

import (
    "net/http"
    "log"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("Serving the main page")
    http.ServeFile(w, r, "./frontend/index.html")
}
