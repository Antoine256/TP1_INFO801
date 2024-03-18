package global

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// Autoriser toutes les origines
			return true
		},
	}
	conn      *websocket.Conn
	writeChan = make(chan []byte, 1)
)

func SendToConn(message string) {
	writeChan <- []byte(message)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var err error
	conn, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	log.Printf("Client connected")

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		log.Printf("Received message: %s", p)

		if err := conn.WriteMessage(messageType, p); err != nil {
			fmt.Println(err)
			return
		}

		if EtatVentilateur == "activé" {
			SendToConn("Ventilateur activé")
		} else {
			SendToConn("Ventilateur désactivé")
		}
		if EtatPompe == "activée" {
			SendToConn("Pompe activée")
		} else {
			SendToConn("Pompe désactivée")
		}
	}
}

func HandleWrites() {
	println("HandleWrites is running...")
	for {
		select {
		case message := <-writeChan:
			if conn != nil {
				err := conn.WriteMessage(websocket.TextMessage, message)
				if err != nil {
					fmt.Println("Error writing to WebSocket:", err)
				}
			}
		}
	}
	println("HandleWrites is done")
}
