package src

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Handler struct {
	DS *DataSources
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	// Parse and decode the request body into a new `LoginRequest` instance
	request := &LoginRequest{}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "Invalid request body",
		})
		return
	}

	// Validate the request body
	if request.Username == "" || request.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "Username and password are required",
		})
		return
	}

	// Check if the username and password are correct
	// For the sake of simplicity, we will consider the username and password correct if they are not empty
	if request.Username != "username" || request.Password != "password" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "Invalid username or password",
		})
		return
	}

	// Create a new session token
	sessionToken := uuid.New().String()

	// Send the response

	json.NewEncoder(w).Encode(map[string]interface{}{
		"session_token": sessionToken,
	})

	// Log the successful login
	h.DS.Log.Info("User logged in: " + request.Username)
}

func (h *Handler) ServeWs(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		h.DS.Log.Error("Error upgrading connection to websocket")
		h.DS.Log.Error(err.Error())
		return
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Infinite loop to handle websocket messages
	for {
		// Read in a new message
		_, msg, err := ws.ReadMessage()
		if err != nil {
			break
		}
		// Write message back to browser
		err = ws.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			break
		}
	}
}
