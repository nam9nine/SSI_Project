package main

import (
	"github.com/nam9nine/SSI_Project/config"
	"github.com/nam9nine/SSI_Project/internal/server"
)

func main() {
	//registrar 서버 시작
	cfg, err := config.LoadConfig("config/config.toml")

	if err != nil {
		panic(err)
	}

	server.StartRegisterServer(cfg)
}
