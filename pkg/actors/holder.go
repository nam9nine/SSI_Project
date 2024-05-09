package actors

import (
	"fmt"
	"github.com/nam9nine/SSI_Project/config"
	client "github.com/nam9nine/SSI_Project/internal/client"
	"github.com/nam9nine/SSI_Project/pkg/core"
	registrar "github.com/nam9nine/SSI_Project/protos/vdr/registrar"
	resolver "github.com/nam9nine/SSI_Project/protos/vdr/resolver"
	"log"
)

type Holder struct {
	DID    *core.DID
	DIDDoc *core.DIDDocument
	Key    *core.EcdsaKeyStorage
	Cfg    *config.Config
}

func (h *Holder) InitHolder(cfg *config.Config) {
	// PKI
	h.Cfg = cfg
	h.Key = new(core.EcdsaKeyStorage)

	publicKey, err := h.Key.GenerateKey()

	if err != nil {
		log.Printf("error generating key: %s", err)
		panic(err)
	}
	//DID
	h.DID = new(core.DID)
	did, err := h.DID.GenerateDID("did", core.KeyMethod, publicKey)

	if err != nil {
		log.Fatalf("새 DID 생성 실패: %v", err)
	}
	//추후 추가 할 때마다 key 번호 증가
	//DID DID Document 등록 및 doc 생성
	verficationId := fmt.Sprintf("%s#keys-1", did)

	h.DIDDoc = new(core.DIDDocument)
	vm := h.DIDDoc.AppendVerificationMethod(verficationId, core.VERIFICATION_KEY_TYPE_SECP256R1, did, publicKey)
	h.DIDDoc.GenerateDIDDocument(did, vm)

	res, err := h.RegistHolderDID()

	if err != nil {
		panic(err)
	}

	log.Printf(res.String())

}

func (h *Holder) RegistHolderDID() (*registrar.DIDRegistrarRes, error) {

	docString, err := h.DIDDoc.Produce()
	if err != nil {
		log.Printf("Error producing DID document: %v", err)
		return nil, err
	}

	// VDR 클라이언트를 통해 DID 문서를 등록합니다.
	req, err := client.RegistrarDID(h.DID.Did, docString)
	if err != nil {

		log.Printf("Error registering DID document: %v", err)
		return nil, err
	}

	return req, nil
}

func (h *Holder) ResolveHolderDID() *resolver.ResolveDIDRes {
	did := h.DID.String()

	if h.Cfg == nil {
		log.Printf("cfg is nil")
	}

	res, err := client.ResolverDID(did, h.Cfg)

	if err != nil {
		log.Printf("Error resolving DID document: %v", err)
		return nil
	}
	return res
}
