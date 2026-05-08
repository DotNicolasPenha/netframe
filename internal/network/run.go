package network

import (
	"fmt"
	"net"

	"com.dv.mp/internal/logger"
	"com.dv.mp/internal/protocol"
)

func Run() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	for {
		clientConn, err := ln.Accept()
		if err != nil {
			logger.LogError(err)
			continue
		}
		go handleClientConn(clientConn)
	}
}

func handleClientConn(clientConn net.Conn) {
	defer clientConn.Close()

	for {
		// this is temporary
		packet, err := protocol.Decode(clientConn)
		if err != nil {
			logger.LogError(err)
			return
		}

		logger.LogInfo(
			fmt.Sprintf(
				"connection %s: packet [LEN %d][OPCODE %d]",
				clientConn.RemoteAddr().String(),
				len(packet.Payload),
				packet.Header.Opcode,
			),
		)

		packetResponse, err := protocol.BuildEncPacket(protocol.PONG, []byte("pong from server"))
		if err != nil {
			logger.LogError(err)
			return
		}
		clientConn.Write(*packetResponse)
	}
}
