package socket

type Room struct {
	clients   map[string]*Client
	game      *Game
	broadcast chan []byte
}

func newRoom() *Room {
	return &Room{
		clients: make(map[string]*Client),
		game:    newGame(),
	}
}

func (r *Room) addClient(c *Client) {
	r.clients[c.id] = c
}

func (r *Room) removeClient(c *Client) {
	c.roomId = ""
	delete(r.clients, c.id)
}
