package actors

import (
	"context"
	"fmt"
	"github.com/nam9nine/SSI_Project/protos/actors/issuer"

	"github.com/nam9nine/SSI_Project/config"
	client "github.com/nam9nine/SSI_Project/internal/client"
	"github.com/nam9nine/SSI_Project/pkg/core"
	registrar "github.com/nam9nine/SSI_Project/protos/vdr/registrar"
	resolver "github.com/nam9nine/SSI_Project/protos/vdr/resolver"
	"log"
)

type Issuer struct {
	DID    *core.DID
	DIDDoc *core.DIDDocument
	Key    *core.EcdsaKeyStorage
	Cfg    *config.Config
}

type IssuerServer struct {
	issuer.UnimplementedIssuerServiceServer

	Issuer *Issuer
}

func (i *IssuerServer) CreateUniversityVC(context context.Context, req *issuer.IssuerReq) (*issuer.IssuerRes, error) {

	res := issuer.IssuerRes{}
	cre := &core.CredentialSubject{}

	cfg, err := config.LoadConfig("./././config/config.toml")
	iss := NewIssuer(cfg)

	i.Issuer = iss
	if err != nil {
		return nil, err
	}

	vc := core.GenerateUniversityVC(i.Issuer.DID.Did, cre, i.Issuer.Key.PrivateKey)

	res.VC = core.ConvertVCtoJSON(vc)

	return &res, nil
}
func NewIssuer(cfg *config.Config) *Issuer {
	iss := &Issuer{
		Cfg:    cfg,
		Key:    new(core.EcdsaKeyStorage),
		DID:    new(core.DID),
		DIDDoc: new(core.DIDDocument),
	}
	err := iss.InitIssuer()

	if err != nil {
		panic(err)
	}
	return iss
}

func (i *Issuer) InitIssuer() error {
	publicKey, err := i.Key.GenerateKey()
	if err != nil {
		return fmt.Errorf("error generating key: %v", err)
	}

	did, err := i.DID.GenerateDID("did", core.KeyMethod, publicKey)
	if err != nil {
		return fmt.Errorf("new DID creation failed: %v", err)
	}

	verificationId := fmt.Sprintf("%s#keys-1", did)
	i.DIDDoc.AppendVerificationMethod(verificationId, core.VERIFICATION_KEY_TYPE_SECP256R1, did, publicKey)
	i.DIDDoc.GenerateDIDDocument(did, nil)

	log.Println("Issuer DID 생성, DID Document 생성 및 VDR에 DID Document 등록 완료.")
	return nil
}

func (i *Issuer) RegistIssuerDID() (*registrar.DIDRegistrarRes, error) {

	docString, err := i.DIDDoc.Produce()
	if err != nil {
		log.Printf("Error producing DID document: %v", err)
		return nil, err
	}

	// VDR 클라이언트를 통해 DID 문서를 등록합니다.
	req, err := client.RegistrarDID(i.DID.Did, docString, registrar.Role_Issuer)
	if err != nil {

		log.Printf("Error registering DID document: %v", err)
		return nil, err
	}

	return req, nil
}

func (i *Issuer) ResolveIssuerDID() (*resolver.ResolveDIDRes, error) {
	did := i.DID.String()

	if i.Cfg == nil {
		return nil, fmt.Errorf("issuer did not have a configuration")

	}

	res, err := client.ResolverDID(did, i.Cfg, resolver.Role_Issuer)

	if err != nil {

		return nil, fmt.Errorf("error resolving DID document: %v", err)
	}
	if res.DidDoc == "" {
		return nil, fmt.Errorf("resolved DID document is empty")
	}
	return res, nil
}

func (i *Issuer) RequestVC(did string) (*issuer.IssuerRes, error) {
	res, err := client.RequestVC(did)

	if err != nil {
		return nil, err
	}

	return res, nil
}
