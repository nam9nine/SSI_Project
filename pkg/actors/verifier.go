package actors

import (
	"context"
	"fmt"
	"github.com/nam9nine/SSI_Project/protos/vdr/registrar"
	"github.com/nam9nine/SSI_Project/protos/vdr/resolver"

	"github.com/nam9nine/SSI_Project/config"
	client "github.com/nam9nine/SSI_Project/internal/client"
	"github.com/nam9nine/SSI_Project/pkg/core"
	"github.com/nam9nine/SSI_Project/protos/actors/verifier"
	"log"
)

type Verifier struct {
	DID    *core.DID
	DIDDoc *core.DIDDocument
	Key    *core.EcdsaKeyStorage
	Cfg    *config.Config
}

type VerifierServer struct {
	verifier.VerifierServiceServer

	Verifier *Verifier
}

func (v *VerifierServer) VerifyVP(context context.Context, req *verifier.VPReq) (*verifier.VPRes, error) {
	res := &verifier.VPRes{}

	cfg, err := config.LoadConfig("./././config/config.toml")
	ver := NewVerifier(cfg)

	v.Verifier = ver

	if err != nil {
		return nil, err
	}

	res.State = "검증됨"
	return res, nil
}

func NewVerifier(cfg *config.Config) *Verifier {
	ver := &Verifier{
		Cfg:    cfg,
		Key:    new(core.EcdsaKeyStorage),
		DID:    new(core.DID),
		DIDDoc: new(core.DIDDocument),
	}
	err := ver.InitVerifier()

	if err != nil {
		panic(err)
	}
	return ver
}

func (v *Verifier) InitVerifier() error {
	publicKey, err := v.Key.GenerateKey()
	if err != nil {
		return fmt.Errorf("error generating key: %v", err)
	}

	did, err := v.DID.GenerateDID("did", core.KeyMethod, publicKey)
	if err != nil {
		return fmt.Errorf("new DID creation failed: %v", err)
	}

	verificationId := fmt.Sprintf("%s#keys-1", did)
	v.DIDDoc.AppendVerificationMethod(verificationId, core.VERIFICATION_KEY_TYPE_SECP256R1, did, publicKey)
	v.DIDDoc.GenerateDIDDocument(did, nil)

	log.Println("Verifier DID 생성, DID Document 생성 및 VDR에 DID Document 등록 완료.")
	return nil
}

func (v *Verifier) RegistVerifierDID() (*registrar.DIDRegistrarRes, error) {

	docString, err := v.DIDDoc.Produce()
	if err != nil {
		log.Printf("Error producing DID document: %v", err)
		return nil, err
	}

	// VDR 클라이언트를 통해 DID 문서를 등록합니다.
	req, err := client.RegistrarDID(v.DID.Did, docString, registrar.Role_Verifier)
	if err != nil {

		log.Printf("Error registering DID document: %v", err)
		return nil, err
	}

	return req, nil
}

func (v *Verifier) ResolveVerifierDID() (*resolver.ResolveDIDRes, error) {
	did := v.DID.String()

	if v.Cfg == nil {
		return nil, fmt.Errorf("verfier did not have a configuration")

	}

	res, err := client.ResolverDID(did, v.Cfg, resolver.Role_Verifier)

	if err != nil {

		return nil, fmt.Errorf("error resolving DID document: %v", err)
	}
	if res.DidDoc == "" {
		return nil, fmt.Errorf("resolved DID document is empty")
	}
	return res, nil
}

func (v *Verifier) RequestVP(vp string) (*verifier.VPRes, error) {
	res, err := client.RequestVP(vp)

	if err != nil {
		return nil, err
	}

	return res, nil
}
