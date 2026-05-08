package handlers

import (
	"fmt"

	"com.dv.mp/internal/protocol"
)

func HandleUnknown(r *protocol.Request) error {
	return r.Client.Send(
		protocol.UNKNOWN,
		[]byte(fmt.Sprintf(
			"unknown opcode: %d",
			r.Packet.Header.Opcode,
		)),
	)
}
