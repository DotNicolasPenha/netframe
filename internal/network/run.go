package network

import (
	"net"

	"com.dv.mp/internal/logger"
)

func (s *Server) Run() error {
	ln, err := net.Listen("tcp", s.Port)
	if err != nil {
		return err
	}
	defer ln.Close()

	logger.LogInfo("Server TCP is running")

	if s.Cfg.DebugMode {
		logger.LogInfo("Debug mode is ON")

	}

	for {
		clientConn, err := ln.Accept()
		if err != nil {
			logger.LogError(err)
			continue
		}
		go s.handleClientConn(clientConn)
	}
}
