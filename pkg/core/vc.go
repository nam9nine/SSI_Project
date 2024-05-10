package core

import (
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

// VC 요청(Holder)(verificationMethod 제시) -> Holder DID -> DID Document -  ->
type VC struct {
	Context           []string           `json:"@context"`
	Id                string             `json:"id"`
	Type              []string           `json:"type"`
	Issuer            string             `json:"issuer"`
	Name              string             `json:"name"`
	CredentialSubject *CredentialSubject `json:"credentialSubject"`
	Proof             *VCProof           `json:"proof"`
}

type CredentialSubject struct {
	Id       string            `json:"id"`
	MtDetail map[string]string `json:"mtDetail"`
}

type VCProof struct {
	Type               string    `json:"type"`
	Created            time.Time `json:"created"`
	VerificationMethod string    `json:"verificationMethod"`
	ProofPurpose       string    `json:"proofPurpose"`
	JWS                string    `json:"proofValue"`
}

func ExampleVC() *VC {
	currentTime := time.Now()

	vc := &VC{
		Context: []string{
			"https://www.w3.org/ns/credentials/v2",
			"https://www.exampleUniversity.com",
		},
		Id:     "http://example.university.edu/credentials/5678",
		Type:   []string{"VerifiableCredential", "EnrollmentCredential"},
		Issuer: "https://example.university.edu",
		Name:   "example university",
		CredentialSubject: &CredentialSubject{
			Id: "did:key:sadfsf",
			MtDetail: map[string]string{
				"id":                  "did:key:subjectDID",
				"2023-10-01T09:00:00": "Yes",
			},
		},
		Proof: &VCProof{
			Type:               "RsaSignature2018",
			Created:            currentTime,
			ProofPurpose:       "assertionMethod",
			VerificationMethod: "https://example.university.edu/keys/1",
			JWS:                "as;ldkfjal;skfdj",
		},
	}

	return vc
}

// GenerateUniversityVC  -> VC 생성
func GenerateUniversityVC(issuerDID string, cs *CredentialSubject, issuerPvKey *ecdsa.PrivateKey) *VC {
	//DID Auth verificationMethodKey string 매개변수, publicKey 암호화, privateKey 복호화 -> (Challenge, Response)
	subVC := VC{
		Context: []string{
			"https://www.w3.org/ns/credentials/v2",
			"https://www.exampleUniversity.com",
		},
		Id:                "1",
		Type:              []string{"VerifiableCredential", "EnrollmentCredential"},
		Issuer:            issuerDID,
		Name:              "example university",
		CredentialSubject: cs,
		Proof:             nil,
	}
	proof := GenerateVCProof(&subVC, issuerDID, issuerPvKey)

	var vc *VC = &VC{
		Context:           subVC.Context,
		Id:                subVC.Id,
		Type:              subVC.Type,
		Issuer:            subVC.Issuer,
		Name:              subVC.Name,
		CredentialSubject: subVC.CredentialSubject,
		Proof:             proof,
	}
	return vc
}

func GenerateJWS(subVC *VC, issuerPvKey *ecdsa.PrivateKey) (string, error) {
	claims := jwt.MapClaims{
		"iss": subVC.Issuer,                                        // 발급자
		"sub": fmt.Sprintf("%s, %s", subVC.Type[0], subVC.Type[1]), // 주체
		"aud": subVC.CredentialSubject.Id,                          // 수신자
		"exp": time.Now().Add(time.Hour * 24).Unix(),               // 만료 시간
		"iat": time.Now().Unix(),                                   // 발급 시간
		"jti": 1,                                                   // 토큰 ID
		"vc":  subVC,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	jws, err := token.SignedString(issuerPvKey)

	if err != nil {
		return "", err
	}

	return jws, nil
}

func GenerateVCProof(subVC *VC, issuerDID string, issuerPvKey *ecdsa.PrivateKey) *VCProof {
	jws, err := GenerateJWS(subVC, issuerPvKey)

	if err != nil {
		panic(err)
	}

	return &VCProof{
		Type:               "JsonWebSignature2020",
		Created:            time.Now(),
		ProofPurpose:       "assertionMethod",
		VerificationMethod: VCVerficationMethod(issuerDID),
		JWS:                jws,
	}
}

func VCVerficationMethod(IssuerDID string) string {
	return fmt.Sprintf("%s#keys-1", IssuerDID)
}

func ConvertVCtoJSON(vc *VC) string {
	jsonData, err := json.Marshal(vc)
	if err != nil {
		log.Fatalf("Error marshalling VC to JSON: %v", err)
	}
	return string(jsonData)
}

func UnmarshalVC(jsonData string) (*VC, error) {
	var vc VC
	err := json.Unmarshal([]byte(jsonData), &vc)
	if err != nil {
		return nil, err
	}
	return &vc, nil
}

func verifyVC(vc *VC, publicKey *ecdsa.PublicKey) error {
	token, err := jwt.Parse(vc.Proof.JWS, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("token is invalid")
	}

	return nil
}
