package main

import (
	"fmt"
	"github.com/nam9nine/SSI_Project/pkg/core"
)

func main() {
	//key 생성

	key := new(core.EcdsaKeyStorage)
	did := new(core.DID)

	_, err := key.GenerateKey()
	if err != nil {
		fmt.Println(err)
	}

	publicKey := key.PublicKeyMultibase()

	didData, err := did.GenerateDID("did", core.KeyMethod, publicKey)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("DID 생성")
	fmt.Println(didData.Did)
}
