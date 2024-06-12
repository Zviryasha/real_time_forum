package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// Initialize the websocket connection
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// --------------
// create variable websocket connection
var clients []*websocket.Conn

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/first", First)
	r.HandleFunc("/second", Second)
	//create webpoint for connect websocket
	r.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		//uinitialize config websocket
		conn, _ := upgrader.Upgrade(w, r, nil)
		clients = append(clients, conn)
		//loop if client send to server
		for {
			//read massage from browser
			msgType, msg, err := conn.ReadMessage()
			//if error
			if err != nil {
				break
			}
			//send massage in your console terminal
			fmt.Printf("%s send: %s\n", conn.RemoteAddr(), string(msg))
			//loop if massege found and send again to client for
			//write in your browser
			for _, client := range clients {
				if err = client.WriteMessage(msgType, []byte("HI")); err != nil {
					return
				}
			}

		}
	})
	//----
	//send you html file for open to browser
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
		//w, r is write and delete your html file
	})
	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8000", r)
}

func First(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "first")
}

func Second(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "second")
}
