package socket

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	conn   *websocket.Conn
	roomId string
	server *Server
	id     string
	name   string
}

func NewClient(conn *websocket.Conn, s *Server) *Client {
	return &Client{
		server: s,
		conn:   conn,
		id:     uuid.New().String(),
	}
}

func (c *Client) ReadMessages() {
	defer func() {
		c.server.mut.Lock()
		delete(c.server.clients, c.id)
		c.server.mut.Unlock()
	}()

	for {
		_, msg, err := c.conn.ReadMessage()
		if len(c.roomId) > 0 && c.server.rooms[c.roomId] != nil {
			c.server.rooms[c.roomId].broadcast <- msg
			continue
		}

		if err != nil {
			break
		}
	}
}

func (c *Client) WriteMessage(msg []byte) {
	c.conn.WriteMessage(websocket.BinaryMessage, msg)
}
