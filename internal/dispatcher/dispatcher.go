package dispatcher

import (
	"com.dv.mp/internal/protocol"
)

type HandlerFunc func(r *protocol.Request) error

type Dispatcher struct {
	handlers map[byte]HandlerFunc
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		handlers: make(map[byte]HandlerFunc),
	}
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
		return d.handlers[protocol.UNKNOWN](r)
	}
	return handler(r)
}
