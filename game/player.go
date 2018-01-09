package game

import (
	"net"
)

type Player struct {
	Conn *net.Conn
	Name string
}
