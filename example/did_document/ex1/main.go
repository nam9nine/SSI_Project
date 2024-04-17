package main

import (
	"fmt"
	"github.com/nam9nine/SSI_Project/pkg/core"
	"log"
)

func main() {

	kms := new(core.EcdsaKeyStorage)
	kms.GenerateKey()

	did, err := core.NewDID("did", core.KeyMethod, kms.PublicKeyMultibase())

	if err != nil {
		log.Printf("Failed to generate DID, error: %v\n", err)
	}

	// Verification Method 생성.
	verificationId := fmt.Sprintf("%s#keys-1", did)
	verificationMethod := []core.VerificationMethod{
		{
			Id:                 verificationId,
			Type:               "EcdsaSecp256k1VerificationKey2019",
			Controller:         did.String(),
			PublicKeyMultibase: kms.PublicKeyMultibase(),
		},
	}

	// DID Document 생성.
	didDocument := core.NewDIDDocument(did.String(), verificationMethod)

	fmt.Println("### New DID, DID Document ###")
	fmt.Printf("did => %s\n", did)
	fmt.Printf("did document => %+v\n", didDocument)

}
