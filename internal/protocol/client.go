package protocol

import (
	"net"
)

type Client struct {
	conn net.Conn
}

func NewClient(conn net.Conn) *Client {
	return &Client{
		conn: conn,
	}
}
func (c *Client) Send(
	opcode byte,
	payload []byte,
) error {

	packet := NewPacket(
		opcode,
		payload,
	)

	resp, err := BuildPacket(packet)
	if err != nil {
		return err
	}

	_, err = c.conn.Write(resp)

	return err
}
