package game

import (
	"github.com/gorilla/websocket"
)

type Master struct {
	Conn *websocket.Conn
	Name string
}
