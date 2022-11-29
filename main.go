package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"net/http"
)

type M mapp[string]interface{}

const MESSAGE_NEW_USER = "New User"
const MESSAGE_CHAT = "Chat"
const MESSAGE_LEAVE = "Leave"

var connections = make([]*WebSocketConnection,0)

type SocketPayload struct {
	Message string
}

type SocketResponse struct {
	From string
	Type string
	Message string
}

type WebSocketConnection struct {
	*websocket.Conn
	Username string

}

func main(){
	http.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		content, err := ioutil.ReadFile("index.html")
		if err != nil {
			http.Error(w,"Cold not open requested file",http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w,"%s",content)
	})

	http.HandleFunc("/ws", func(writer http.ResponseWriter, r *http.Request) {
	currentGorillaConn, err := websocket.Upgrade(w,r,w.Header(),1024,1024)
	if err !=nil {
		http.Error(w,"Could not open websocket connection",http.StatusBadRequestq)

	}

	username := r.URL.Query().Get("username")
	currentConn := WebSocketConnection{Conn: currentGorillaConn, Username: username}
	connections = append(connections,&currentConn)

	go handleIO(&currentConn,connections)
	})


	fmt.Println("server starting at :8080")
	http.ListenAndServe(":8080",nil)
}