package protocol

type Request struct {
	Client *Client
	Packet *Packet
}

func NewRequest(client *Client, packet *Packet) *Request {
	return &Request{
		Client: client,
		Packet: packet,
	}
}
