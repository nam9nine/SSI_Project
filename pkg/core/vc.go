package core

type VC struct {
	// Mendatory
	Context []string `json:"@context"`

	Id                string              `json:"id,omitempty"`
	Type              []string            `json:"type,omitempty"`
	Issuer            string              `json:"issuer,omitempty"`
	IssuanceDate      string              `json:"issuanceDate,omitempty"`
	CredentialSubject []CredentialSubject `json:"credentialSubject,omitempty"`
	Proof             *Proof              `json:"proof,omitempty"`
}

type Proof struct {
	Type               string `json:"type,omitempty"`
	Created            string `json:"created,omitempty"`
	ProofPurpose       string `json:"proofPurpose,omitempty"`
	VerificationMethod string `json:"verificationMethod,omitempty"`
	ProofValue         string `json:"proofValue,omitempty"`
	Jws                string `json:"jws,omitempty"`
}

type CredentialSubject struct {
}
