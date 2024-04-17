package core

import (
	"encoding/json"
	"errors"
)

type (
	Type string
)

const (
	VERIFICATION_KEY_TYPE_SECP256R1 = "P256VerificationKey2018"
)

type DIDDocument struct {
	Context []string `json:"@context"`

	Id                   string               `json:"id"`
	AlsoKnownAs          []string             `json:"alsoKnownAs,omitempty"`
	Controller           string               `json:"controller,omitempty"`
	VerificationMethod   []VerificationMethod `json:"verificationMethod,omitempty"`
	Authentication       []Authentication     `json:"authentication,omitempty"`
	AssertionMethod      string               `json:"assertionMethod,omitempty"`
	KeyAgreement         string               `json:"keyAgreement,omitempty"`
	CapabilityInvocation string               `json:"capabilityInvocation,omitempty"`
	CapabilityDelegation string               `json:"capabilityDelegation,omitempty"`
	Service              []Service            `json:"service,omitempty"`
}

type VerificationMethod struct {
	Id                 string `json:"id"`
	Type               string `json:"type"`
	Controller         string `json:"controller"`
	PublicKeyMultibase string `json:"PublicKeyMultibase,omitempty"`
}

type Authentication struct {
	Id                 string `json:"id"`
	Type               string `json:"type"`
	Controller         string `json:"controller"`
	PublicKeyBase58    string `json:"publicKeyBase58,omitempty"`
	PublicKeyMultibase string `json:"PublicKeyMultibase,omitempty"`
}

type Service struct {
	Id              string `json:"id"`
	Type            string `json:"type"`
	ServiceEndpoint string `json:"serviceEndpoint"`
}

// GenerateDIDDocument 추후 json으로 매개변수 받기
func (doc *DIDDocument) GenerateDIDDocument(did string, vm []VerificationMethod) *DIDDocument {

	doc.Context = []string{"https://www.w3.org/ns/did/v1"}
	doc.Id = did
	doc.VerificationMethod = vm
	return doc
}

func (doc *DIDDocument) AppendVerificationMethod(id string, t string, controller string, publickey string) []VerificationMethod {
	newVm := VerificationMethod{
		Id:                 id,
		Type:               t,
		Controller:         controller,
		PublicKeyMultibase: publickey,
	}

	doc.VerificationMethod = append(doc.VerificationMethod, newVm)
	return doc.VerificationMethod
}

func (doc *DIDDocument) Produce() (string, error) {
	docStr, err := json.Marshal(doc)

	if err != nil {
		return "", errors.New("error marshalling DID document")
	}

	return string(docStr), nil
}

func (doc *DIDDocument) Consume(str string) (string, error) {
	err := json.Unmarshal([]byte(str), doc)
	if err != nil {
		return "", errors.New("error unmarshalling DID document")
	}
	return str, nil
}
