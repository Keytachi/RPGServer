package game

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net"
)

type Game struct {
	DM      Master
	Players []Player
}

func (m *Game) ListenToDm(c *websocket.Conn) {

	if m.DM.Conn != nil {
		m.DM.Conn.Close()
	}
	m.DM = Master{Conn: c}

	m.handleDm()

}

func (m *Game) ListenToPlayer(c *net.Conn) {
	p := Player{Conn: c}
	m.Players = append(m.Players, p)
	m.handlePlayer(p)

}

func (m *Game) handleDm() {
	currentDM := &(m.DM)
	fmt.Println(currentDM)
	for {
		_, p, err := (*currentDM).Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(string(p))
	}
}

func (m *Game) handlePlayer(p Player) {
	r := bufio.NewReader(*p.Conn)
	w := bufio.NewWriter(*p.Conn)
	fmt.Printf("Handling Player\n")
	for {
		w.Write([]byte("harhar\n"))
		w.Flush()
		msg, _, err := r.ReadLine()
		if err != nil {
			fmt.Println(err)
			for i := 0; i < len(m.Players); i++ {
				if p == m.Players[i] {
					m.Players = append(m.Players[:i], m.Players[i+1:]...)
				}
			}
			break
		}
		fmt.Printf("%s\n", string(msg))

	}
}
