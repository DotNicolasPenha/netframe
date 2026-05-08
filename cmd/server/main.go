package main

import "com.dv.mp/internal/network"

func main() {
	cfg := network.Config{
		DebugMode: false,
		MaxConns:  100,
	}
	port := ":8080"
	name := "MineServer"
	srv := network.NewServer(&cfg, port, name)
	srv.Run()
}
