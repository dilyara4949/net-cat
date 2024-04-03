package main

const (
	maxConn = 10
)

type Chat struct {
	MaxConn int
	Clients []*Client
	History []string
}

func NewChat() *Chat {
	return &Chat{
		MaxConn: maxConn,
		Clients: make([]*Client, 0),
		History: make([]string, 0),
	}
}

func (c *Chat) Broadcast(sender *Client, msg string) {
	c.History = append(c.History, msg)
	for _, m := range c.Clients {
		if sender.Conn.RemoteAddr() != m.Conn.RemoteAddr() {
			m.Msg(msg)

		}
	}
}
