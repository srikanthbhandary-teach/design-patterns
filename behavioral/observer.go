package behavioral

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type Observer interface {
	Notify(message []byte)
	Close()
}

type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers(message []byte)
}

type Client struct {
	conn     *websocket.Conn
	send     chan []byte
	server   Subject // Reference to the server as the Subject
	isClosed bool    // Flag to track if the connection is closed
}

func (c *Client) Close() {
	c.server.RemoveObserver(c) // Remove client from the server's observers
	c.conn.Close()
	close(c.send)

}
func (c *Client) Notify(message []byte) {
	if !c.isClosed {
		c.send <- message
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type WebSocketServer struct {
	observers []Observer
}

func (s *WebSocketServer) RegisterObserver(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *WebSocketServer) RemoveObserver(observer Observer) {
	for i, obs := range s.observers {
		if obs == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *WebSocketServer) NotifyObservers(message []byte) {
	for _, observer := range s.observers {
		observer.Notify(message)
	}
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Could not open websocket connection", err.Error())
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		return
	}

	client := &Client{
		conn:     conn,
		send:     make(chan []byte, 256),
		server:   WebsocketServer, // Reference to the server as the Subject
		isClosed: false,
	}

	// Register the client as an observer
	WebsocketServer.RegisterObserver(client)

	defer func() {
		// Unregister the client when the connection is closed
		WebsocketServer.RemoveObserver(client)
		client.conn.Close()
	}()

	go client.write()
	client.read()
}

func (c *Client) read() {
	for {
		_, _, err := c.conn.ReadMessage()
		if err != nil {
			c.isClosed = true
			break
		}
	}
}

func (c *Client) write() {
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				return
			}
			err := c.conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				return
			}
		}
	}
}

func NotifyAll(message []byte) {
	WebsocketServer.NotifyObservers(message)
}

var WebsocketServer *WebSocketServer

func init() {
	WebsocketServer = &WebSocketServer{
		observers: []Observer{},
	}
}

func (s *WebSocketServer) CloseAllConnections() {
	for _, observer := range s.observers {
		observer.Close()
	}
}
