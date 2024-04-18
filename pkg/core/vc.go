package core

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type VC struct {
	Context           []string          `json:"@context"`
	ID                string            `json:"id"`
	Type              []string          `json:"type"`
	Issuer            string            `json:"issuer"`
	IssuanceDate      time.Time         `json:"issuanceDate"`
	CredentialSubject CredentialSubject `json:"credentialSubject"`
	Proof             SignatureProof    `json:"proof"`
	UsagePolicy       UsagePolicy
}

type CredentialSubject struct {
	StudentID   string `json:"studentId"`
	StudentName string `json:"studentName"`
	SchoolName  string `json:"schoolName"`
	Department  string `json:"department"`
	ClassYear   int    `json:"classYear"`
	Status      string
	UsagePolicy UsagePolicy `json:"usagePolicy"`
}

type SignatureProof struct {
	Type               string    `json:"type"`
	Created            time.Time `json:"created"`
	ProofPurpose       string    `json:"proofPurpose"`
	VerificationMethod string    `json:"verificationMethod"`
	JWS                string    `json:"jws"`
}

type UsagePolicy struct {
	UseCase      string `json:"useCase"`
	Restrictions string `json:"restrictions"`
	DataSharing  string `json:"dataSharing"`
}

func GenerateVCl() *VC {
	currentTime := time.Now()

	vc := &VC{
		Context: []string{
			"https://www.w3.org/2018/credentials/v1",
			"https://www.w3.org/2018/credentials/examples/v1",
		},
		ID:           "http://example.university.edu/credentials/5678",
		Type:         []string{"VerifiableCredential", "EnrollmentCredential"},
		Issuer:       "https://example.university.edu",
		IssuanceDate: currentTime,
		CredentialSubject: CredentialSubject{
			StudentID:   "did:example:123456",
			StudentName: "John Doe",
			SchoolName:  "Example University",
			Department:  "Computer Science",
			ClassYear:   2023,
			Status:      "Enrolled",
		},
		Proof: SignatureProof{
			Type:               "RsaSignature2018",
			Created:            currentTime,
			ProofPurpose:       "assertionMethod",
			VerificationMethod: "https://example.university.edu/keys/1",
			JWS:                "eyJhbGciOiJSUzI1NiIsImtpZCI...",
		},
		UsagePolicy: UsagePolicy{
			UseCase:      "Academic Verification",
			Restrictions: "No commercial use",
			DataSharing:  "Data cannot be shared with third party without consent",
		},
	}

	return vc
}

func createToken() string {
	claims := jwt.MapClaims{
		"iss":   "https://example.university.edu",
		"sub":   "did:example:123456",
		"name":  "John Doe",
		"admin": true,
		"iat":   time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("your-256-bit-secret"))
	if err != nil {
		fmt.Println("Error signing token: ", err)
		return ""
	}
	return signedToken
}

func verifyToken(signedToken, secretKey string) (*jwt.Token, error) {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	return token, err
}
