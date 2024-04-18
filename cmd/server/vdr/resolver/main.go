package main

import (
	"github.com/nam9nine/SSI_Project/config"
	"github.com/nam9nine/SSI_Project/internal/server"
)

func main() {
	cfg, err := config.LoadConfig("config/config.toml")

	if err != nil {
		panic(err)
	}

	server.StartDIDResolverServer(cfg)
}
