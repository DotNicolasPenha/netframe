package dispatcher

import (
	"com.dv.mp/internal/dispatcher/handlers"
	"com.dv.mp/internal/protocol"
)

type HandlerFunc func(r *protocol.Request) error

type Dispatcher struct {
	handlers map[byte]HandlerFunc
}

func NewDispatcher() *Dispatcher {
	d := Dispatcher{
		handlers: make(map[byte]HandlerFunc),
	}
	d.Register(protocol.PING, handlers.HandlePing)
	return &d
}

func (d *Dispatcher) Register(
	opcode byte,
	handler HandlerFunc,
) {
	d.handlers[opcode] = handler
}

func (d *Dispatcher) Dispatch(r *protocol.Request) error {
	handler, ok := d.handlers[r.Packet.Header.Opcode]
	if !ok {
		return handlers.HandleUnknown(r)
	}
	return handler(r)
}
