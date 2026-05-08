package protocol

type Header struct {
	Opcode byte
}
type Packet struct {
	Header  Header
	Payload []byte
}

func NewPacket(opcode byte, payload []byte) *Packet {
	return &Packet{
		Header: Header{
			Opcode: opcode,
		},
		Payload: payload,
	}
}
func BuildPacket(packet *Packet) ([]byte, error) {
	return Encode(packet)
}
