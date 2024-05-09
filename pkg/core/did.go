package core

import (
	"errors"
	"fmt"
)

type (
	Method string
)

const (
	KeyMethod  Method = "key"
	EthrMethod Method = "ethr"
	BtcrMethod Method = "btcr"
)

type DID struct {
	Scheme                   string
	Method                   Method
	MethodSpecificIdentifier string
	Did                      string
}

func (did *DID) GenerateDID(scheme string, method Method, methodSpecificIdentifier string) (string, error) {
	if scheme != "did" || method == "" || methodSpecificIdentifier == "" {
		return "", errors.New("invalid did parameters")
	}

	did.Scheme = scheme
	did.Method = method
	did.MethodSpecificIdentifier = methodSpecificIdentifier
	did.Did = fmt.Sprintf("%s:%s:%s", scheme, method, methodSpecificIdentifier)

	return did.String(), nil
}

func (did *DID) String() string {
	return did.Did
}
