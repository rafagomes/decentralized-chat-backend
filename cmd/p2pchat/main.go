package main

import (
	"flag"
	"log"
	"net/http"
	"p2pchat/internal/interface/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	flag.Parse()
	//ctx := context.Background()

	hub := websocket.NewHub()
	go hub.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websocket.ServeWs(hub, w, r)
	})

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	//go network.InitP2pNode()
	//
	//chatUsecase := &chat.ChatUsecase{}
	//chatHandler := &rest.ChatHandler{ChatUsecase: chatUsecase}
	//
	//http.HandleFunc("/send", chatHandler.SendMessage)
	//
	//log.Fatal(http.ListenAndServe(":"+*port, nil))
}
