package main

import (
	"log"
	"net/http"
	"path/filepath"
	"database/sql"
	"real-time-forum/backend/handlers"
	"real-time-forum/backend/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	dbConn, err := sql.Open("sqlite3", "./database/forum.db")
	if err != nil {
		log.Fatal("Error opening database:", err)
	}
	defer dbConn.Close()

	// Serve static files
	fs := http.FileServer(http.Dir("./frontend"))
	http.Handle("/frontend/", http.StripPrefix("/frontend", fs))

	// API Handlers
	http.HandleFunc("/api/login", handlers.LoginHandler(dbConn))
	http.HandleFunc("/api/logout", handlers.LogoutHandler)
	http.HandleFunc("/api/register", handlers.RegisterHandler(dbConn))
	http.Handle("/api/post", middleware.AuthMiddleware(http.HandlerFunc(handlers.PostsHandler(dbConn))))
	http.Handle("/api/posts", http.HandlerFunc(handlers.PostsHandler(dbConn))) // No middleware for testing
	http.HandleFunc("/ws", handlers.WsHandler)

	// Serve index.html for any other route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" || filepath.Ext(r.URL.Path) == "" || r.URL.Path == "/home" || r.URL.Path == "/login" || r.URL.Path == "/register" || r.URL.Path == "/post" {
			http.ServeFile(w, r, "./frontend/index.html")
		} else {
			fs.ServeHTTP(w, r)
		}
	})
	

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
