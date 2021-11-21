package websockets

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

type WebSocet struct{
}

func NewWebSockets() *WebSocet{
	return &WebSocet{}
}


var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}


func upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return ws, err
	}

	return ws, nil
}


func writer(conn *websocket.Conn) {

	for {
		ticker := time.NewTicker(5 * time.Second)
		for t := range ticker.C {
			fmt.Println(t)
		}
	}
}

func TestSocket(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}
	go writer(ws)
}