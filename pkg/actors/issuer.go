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

	vc := core.GenerateUniversityVC(i.Issuer.DID.Did, cre, i.Issuer.Key.PrivateKey)

	res.VC = vc.Proof.JWS

	return &res, nil
}

func (i *Issuer) InitIssuer(cfg *config.Config) {
	// PKI
	i.Cfg = cfg
	i.Key = new(core.EcdsaKeyStorage)

	publicKey, err := i.Key.GenerateKey()

	if err != nil {
		log.Printf("error generating key: %s", err)
		panic(err)
	}
	//DID
	i.DID = new(core.DID)
	did, err := i.DID.GenerateDID("did", core.KeyMethod, publicKey)

	if err != nil {
		log.Fatalf("새 DID 생성 실패: %v", err)
	}
	//추후 추가 할 때마다 key 번호 증가
	//DID DID Document 등록 및 doc 생성
	verficationId := fmt.Sprintf("%s#keys-1", did)

	i.DIDDoc = new(core.DIDDocument)
	vm := i.DIDDoc.AppendVerificationMethod(verficationId, core.VERIFICATION_KEY_TYPE_SECP256R1, did, publicKey)
	i.DIDDoc.GenerateDIDDocument(did, vm)

	res, err := i.RegistIssuerDID()

	if err != nil {
		panic(err)
	}

	log.Printf(res.String())

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
