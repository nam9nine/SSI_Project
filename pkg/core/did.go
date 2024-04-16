package core

import (
	"errors"
	"fmt"
)

package core

import (
"errors"
"fmt"
)

type DID struct {
	scheme                   string
	method                   string
	methodSpecificIdentifier string
	DIDKey
}

type (
	DIDKey string
)

func NewDID(scheme string, method string, methodSpecificIdentifier string) (*DID, error) {
	if scheme != "did" || method == "" || methodSpecificIdentifier == "" {
		return nil, errors.New("invalid did parameters")
	}

	var newDID = new(DID)

	newDID.scheme = scheme
	newDID.method = method
	newDID.methodSpecificIdentifier = methodSpecificIdentifier
	newDID.DIDKey = DIDKey(fmt.Sprintf("%s:%s:%s", scheme, method, methodSpecificIdentifier))

	return newDID, nil
}
