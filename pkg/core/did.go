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
	scheme                   string
	method                   Method
	methodSpecificIdentifier string
	did                      string
}

func NewDID(scheme string, method Method, methodSpecificIdentifier string) (*DID, error) {
	if scheme != "did" || method == "" || methodSpecificIdentifier == "" {
		return nil, errors.New("invalid did parameters")
	}

	var newDID = new(DID)

	newDID.scheme = scheme
	newDID.method = method
	newDID.methodSpecificIdentifier = methodSpecificIdentifier
	newDID.did = fmt.Sprintf("%s:%s:%s", scheme, method, methodSpecificIdentifier)

	return newDID, nil
}
