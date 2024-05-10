package core

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

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

func GenerateVP(vc *VC, holderPvKey *ecdsa.PrivateKey) (*VP, error) {
	vp := &VP{
		Id:                   "vp-12345",
		Type:                 "VerifiablePresentation",
		VerifiableCredential: []VC{*vc}, // VP에 포함될 VC 목록
		Proof: VPProof{
			Type:               "JsonWebSignature2020",
			Created:            time.Now(),
			VerificationMethod: "https://example.com/keys/1",
			ProofPurpose:       "authentication",
			ProofValue:         "", // 이 값은 나중에 생성할 예정
		},
	}

	// VPProof 생성 및 할당
	if err := GenerateVPProof(vp, holderPvKey); err != nil {
		return nil, fmt.Errorf("failed to generate VP proof: %v", err)
	}

	return vp, nil
}

func GenerateVPProof(vp *VP, holderPvKey *ecdsa.PrivateKey) error {

	claims := jwt.MapClaims{
		"iss": "https://example.com",
		"sub": vp.Id,
		"iat": time.Now().Unix(),
		"vp":  vp,
	}

	// JWT 생성 및 서명
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	signedString, err := token.SignedString(holderPvKey)
	if err != nil {
		return fmt.Errorf("error signing JWT: %v", err)
	}

	vp.Proof.ProofValue = signedString
	return nil
}

func verifyVP(vp *VP, publicKey *ecdsa.PublicKey) error {
	// JWT 파싱 및 검증
	token, err := jwt.Parse(vp.Proof.ProofValue, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	if err != nil {
		return fmt.Errorf("error parsing token: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Claim에서 필요한 정보 검증
		fmt.Printf("Issuer: %v, Subject: %v\n", claims["iss"], claims["sub"])
		// 추가적인 클레임 검증이 필요한 경우 여기에 로직 추가
	} else {
		return fmt.Errorf("invalid token or claims")
	}

	// VCs 검증 (예시로 간단히 처리)
	for _, vc := range vp.VerifiableCredential {
		if vc.Id == "" {
			return fmt.Errorf("VC has no ID")
		}
	}

	return nil
}
