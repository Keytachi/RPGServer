package main

import (
	"dm-tools/game"
	"dm-tools/routes"
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	game := game.Game{}
	router := routes.Router{Game: game}

	ln, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatal(err)
	}
	go listenForPlayers(ln, game)
	http.HandleFunc("/dm", router.DmToolPage)
	http.HandleFunc("/dm_socket", router.DmWebSocket)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./client"))))
	fmt.Println("Listening on port 8000...")
	go log.Fatal(http.ListenAndServe(":8000", nil))

}

func listenForPlayers(ln net.Listener, game game.Game) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("Error Accepting Player")
		}
		game.ListenToPlayer(&conn)

	}
}
