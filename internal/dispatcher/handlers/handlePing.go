package handlers

import "com.dv.mp/internal/protocol"

func HandlePing(r protocol.Request) error {
	return r.Client.Send(
		protocol.PONG,
		[]byte("PONG! from server."),
	)
}
