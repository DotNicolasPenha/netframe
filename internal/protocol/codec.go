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
		uint32(p.Header.Length),
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

	hbuf := make([]byte, headerSize)
	_, err := io.ReadFull(r, hbuf)
	if err != nil {
		return nil, err
	}

	h.Opcode = hbuf[0]
	h.Length = binary.BigEndian.Uint32(hbuf[1:headerSize])

	if h.Length > payloadMaxLength {
		return nil, ErrPayloadExceedsMaximumAllowedSize
	}

	payloadBuffer := make([]byte, h.Length)
	_, err = io.ReadFull(r, payloadBuffer)
	if err != nil {
		return nil, err
	}

	return &Packet{
		Header:  h,
		Payload: payloadBuffer,
	}, nil
}
