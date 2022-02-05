package main

import (
	"fmt"
	"net/http"

	"github.com/sajjadth/websocket/socket"
)

func main() {
	hub := socket.NewHub()
	go hub.Run()

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "Home")
	})

	http.HandleFunc("/ws", func(rw http.ResponseWriter, r *http.Request) {
		socket.ServeWs(hub, rw, r)
	})
	http.ListenAndServe(":8080", nil)
}
