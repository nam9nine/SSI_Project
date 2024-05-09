package core

import (
	"time"
)

type VC struct {
	Context           []string          `json:"@context"`
	ID                string            `json:"id"`
	Type              []string          `json:"type"`
	Issuer            string            `json:"issuer"`
	Name              string            `json:"name"`
	CredentialSubject CredentialSubject `json:"credentialSubject"`
	Proof             VCProof           `json:"proof"`
}

type CredentialSubject struct {
	Id       string            `json:"id"`
	MtDetail map[string]string `json:"attend"`
}

type VCProof struct {
	Type               string    `json:"type"`
	Created            time.Time `json:"created"`
	VerificationMethod string    `json:"verificationMethod"`
	ProofPurpose       string    `json:"proofPurpose"`
	ProofValue         string    `json:"proofValue"`
}

func ExampleVC() *VC {
	currentTime := time.Now()

	vc := &VC{
		Context: []string{
			"https://www.w3.org/2018/credentials/v1",
			"https://www.w3.org/2018/credentials/examples/v1",
		},
		ID:     "http://example.university.edu/credentials/5678",
		Type:   []string{"VerifiableCredential", "EnrollmentCredential"},
		Issuer: "https://example.university.edu",
		Name:   "example university",
		CredentialSubject: CredentialSubject{
			Id: "did:key:sadfsf",
			MtDetail: map[string]string{
				"id":                  "did:key:subjectDID",
				"2023-10-01T09:00:00": "Yes",
			},
		},
		Proof: VCProof{
			Type:               "RsaSignature2018",
			Created:            currentTime,
			ProofPurpose:       "assertionMethod",
			VerificationMethod: "https://example.university.edu/keys/1",
			ProofValue:         "as;ldkfjal;skfdj",
		},
	}

	return vc
}
