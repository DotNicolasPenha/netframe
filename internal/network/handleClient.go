package network

import (
	"fmt"
	"io"
	"net"
	"time"

	"com.dv.mp/internal/logger"
	"com.dv.mp/internal/protocol"
)

func (s *Server) handleClientConn(clientConn net.Conn) {
	defer clientConn.Close()

	addr := clientConn.RemoteAddr().String()
	defer logger.LogInfo(fmt.Sprintf("Client %s disconnected", addr))
	logger.LogInfo(fmt.Sprintf("Client %s connected", addr))

	client := protocol.NewClient(clientConn)

	for {
		err := clientConn.SetReadDeadline(time.Now().Add(30 * time.Second))
		if err != nil {
			return
		}

		packet, err := protocol.Decode(clientConn)
		if err != nil {
			if err != io.EOF {
				logger.LogError(err)
				return
			}
		}

		request := protocol.NewRequest(client, packet)

		err = s.dispatcher.Dispatch(request)
		if err != nil {
			logger.LogError(err)
			return
		}
	}
}
