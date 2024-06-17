package main

import (
	"net/http"
	"real_time_forum/src"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Handler struct {
	ds *src.DataSources
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	h.ds.Log.Info("Login")
}

func main() {
	ds := src.InitDataSources()

	handler := &Handler{
		ds: ds,
	}
	r := mux.NewRouter()

	r.HandleFunc("/login", handler.Login)
	r.HandleFunc("/ws", handler.handleConnections)

	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, ".js") {
			w.Header().Set("Content-Type", "application/javascript")
		} else if strings.HasSuffix(r.URL.Path, ".css") {
			w.Header().Set("Content-Type", "text/css")
		}
		http.FileServer(http.Dir("./static")).ServeHTTP(w, r)
	})

	handler.ds.Log.Info("Server started on :8080")
	http.ListenAndServe(":8080", r)
}

func (h *Handler) handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		h.ds.Log.Error("Error upgrading connection to websocket")
		h.ds.Log.Error(err.Error())
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
