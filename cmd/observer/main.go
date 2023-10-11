package main

import (
	"design-patterns/behavioral"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	http.HandleFunc("/ws", behavioral.HandleWebSocket)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "test.html")
	})

	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {

		fmt.Println("Server running on :8080")
		err := http.ListenAndServe(":8080", nil)

		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()

	for {
		select {
		case <-ticker.C:
			behavioral.NotifyAll([]byte("Yikes!"))
		case <-quit:
			fmt.Println("exiting the server")
			behavioral.WebsocketServer.CloseAllConnections()
			ticker.Stop()
			return
		}
	}
}
