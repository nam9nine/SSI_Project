package main

import (
	"fmt"
	"github.com/nam9nine/SSI_Project/pkg/core"
	"log"
)

func main() {
	key := new(core.EcdsaKeyStorage)

	err := key.GenerateKey()

	if err != nil {
		log.Println(err)
	}

	publicKey, err := key.PublicKeyToString()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("-----private Key--------")
	fmt.Println(publicKey)
}
