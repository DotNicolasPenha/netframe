package network

import "com.dv.mp/internal/dispatcher"

type Server struct {
	Port       string
	Name       string
	dispatcher *dispatcher.Dispatcher
	Cfg        *Config
}
type Config struct {
	DebugMode bool
	MaxConns  int
}

func NewServer(cfg *Config, port string, name string) *Server {
	d := dispatcher.NewDispatcher()
	return &Server{
		Port:       port,
		Name:       name,
		dispatcher: d,
		Cfg:        cfg,
	}
}
