package main

import (
	"fmt"
	"github.com/nam9nine/SSI_Project/config"
	"github.com/nam9nine/SSI_Project/pkg/actors"
	"log"
)

func main() {

	hldr := new(actors.Holder)

	cfg, err := config.LoadConfig("config/config.toml")

	if err != nil {
		log.Println(err)
		panic(err)
	}
	// DID, DID Document 등록
	hldr.InitHolder(cfg)
	res := hldr.ResolveHolderDID()

	fmt.Println(res)

}
