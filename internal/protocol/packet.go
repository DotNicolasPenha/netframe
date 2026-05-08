package protocol

type Header struct {
	Opcode byte
}
type Packet struct {
	Header  Header
	Payload []byte
}

func BuildEncPacket(opcode byte, payload []byte) (*[]byte, error) {
	packet := Packet{
		Header: Header{
			Opcode: opcode,
		},
		Payload: payload,
	}
	packetInBytes, err := Encode(&packet)
	if err != nil {
		return nil, err
	}
	return &packetInBytes, nil
}
