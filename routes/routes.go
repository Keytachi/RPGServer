package routes

import (
	"dm-tools/game"
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"net/http"
)

//Router holds route definitions
type Router struct {
	Game game.Game
}

func (m *Router) DmToolPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	data, err := ioutil.ReadFile("./client/dm_client.html")
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Length", fmt.Sprint(len(data)))
	fmt.Fprint(w, string(data))
}

func (m *Router) DmWebSocket(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Origin") != "http://"+r.Host {
		http.Error(w, "Origin Not Allowed", 403)
		return
	}

	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		http.Error(w, "Could Not Open websocket Connection", http.StatusBadRequest)
	}

	//fmt.Println("Games Count: %d", m.Game.Count)
	m.Game.ListenToDm(conn)
}
