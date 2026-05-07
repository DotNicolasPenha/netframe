package protocol

type Header struct {
	Opcode byte
	Length uint32
}
type Packet struct {
	Header  Header
	Payload []byte
}
