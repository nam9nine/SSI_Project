package core

import "time"

type VP struct {
	Id                   string
	Type                 string
	VerifiableCredential []VC
	Proof                VPProof
}

type VPProof struct {
	Type               string    `json:"type"`
	Created            time.Time `json:"created"`
	VerificationMethod string    `json:"verificationMethod"`
	ProofPurpose       string    `json:"proofPurpose"`
	ProofValue         string    `json:"proofValue"`
}
