package client

import (
	"context"
	"errors"
	"github.com/nam9nine/SSI_Project/config"
	resolver "github.com/nam9nine/SSI_Project/protos/vdr/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ResolverDID DIDResolver 서버에게 요청
func ResolverDID(didString string, cfg *config.Config) (*resolver.ResolveDIDRes, error) {
	req := &resolver.ResolveDIDReq{
		Did: didString,
	}

	conn, err := grpc.Dial(cfg.Servers.Resolver.Address(), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, errors.New("resolver 서버와 연결되지 않음")
	}

	defer conn.Close()

	newClient := resolver.NewDIDResolverClient(conn)

	res, err := newClient.ResolveDID(context.Background(), req)

	return res, nil
}
