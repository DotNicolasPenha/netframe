package protocol

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

const (
	headerSize       = 5
	payloadMaxLength = 1024 * 1024 // 1MB
)

var (
	ErrPayloadExceedsMaximumAllowedSize = errors.New("payload exceeds maximum allowed size")
)

func Encode(p *Packet) ([]byte, error) {
	buf := bytes.NewBuffer(nil)

	err := buf.WriteByte(p.Header.Opcode)
	if err != nil {
		return nil, err
	}

	err = binary.Write(
		buf,
		binary.BigEndian,
		uint32(len(p.Payload)),
	)
	if err != nil {
		return nil, err
	}

	_, err = buf.Write(p.Payload)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
func Decode(r io.Reader) (*Packet, error) {
	var h Header

	headerbuf := make([]byte, headerSize)
	_, err := io.ReadFull(r, headerbuf)
	if err != nil {
		return nil, err
	}

	h.Opcode = headerbuf[0]
	payloadLength := binary.BigEndian.Uint32(headerbuf[1:headerSize])

	if payloadLength > payloadMaxLength {
		return nil, ErrPayloadExceedsMaximumAllowedSize
	}

	payloadBuffer := make([]byte, payloadLength)
	_, err = io.ReadFull(r, payloadBuffer)
	if err != nil {
		return nil, err
	}

	return &Packet{
		Header:  h,
		Payload: payloadBuffer,
	}, nil
}
