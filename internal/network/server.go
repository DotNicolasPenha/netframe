package network

import "com.dv.mp/internal/dispatcher"

type Server struct {
	Port       string
	Name       string
	Dispatcher *dispatcher.Dispatcher
	Cfg        *Config
}
type Config struct {
	DebugMode bool
	MaxConns  int
}

func NewServer(cfg *Config) *Server {
	d := dispatcher.NewDispatcher()
	return &Server{
		Port:       ":8080",
		Name:       "MineServer",
		Dispatcher: d,
		Cfg:        cfg,
	}
}
