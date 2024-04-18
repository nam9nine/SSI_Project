package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
)

type ServerConfig struct {
	IP    string `toml:"ip"`
	Ports string `toml:"ports"`
}

type Servers struct {
	Registrar ServerConfig `toml:"registrar"`
	Resolver  ServerConfig `toml:"resolver"`
	Holder    ServerConfig `toml:"holder"`
}

type Config struct {
	Servers Servers `toml:"servers"`
}

func (sc *ServerConfig) Address() string {
	return fmt.Sprintf("%s:%s", sc.IP, sc.Ports)
}

func LoadConfig(path string) (*Config, error) {

	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		log.Fatalf("Error parsing config file: %s", err)
		return nil, err
	}
	return &config, nil
}
