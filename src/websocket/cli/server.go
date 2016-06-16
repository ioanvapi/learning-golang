package main

import (
	"io"
	"net/http"

	"golang.org/x/net/websocket"
)

func echoHandler(ws *websocket.Conn) {
	// msg := make([]byte, 512)
	// n, err := ws.Read(msg)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Receive: %s\n", msg[:n])
	//
	// m, err := ws.Write(msg[:n])
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Send: %s\n", msg[:m])

	io.Copy(ws, ws)
}

func main() {
	http.Handle("/echo", websocket.Handler(echoHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
