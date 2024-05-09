package server

import (
	"context"
	"github.com/nam9nine/SSI_Project/config"
	resolver "github.com/nam9nine/SSI_Project/protos/vdr/resolver"
	"github.com/syndtr/goleveldb/leveldb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type ResolverServer struct {
	resolver.UnimplementedDIDResolverServer
}

func byteToString(b []byte) string {
	return string(b[:len(b)])
}

// ResolveDID ResolveDID(context.Context, *ResolveDIDReq) (*ResolveDIDRes, error) RPC메서드 구현
func (r *ResolverServer) ResolveDID(ctx context.Context, req *resolver.ResolveDIDReq) (*resolver.ResolveDIDRes, error) {

	db, err := leveldb.OpenFile("./internal/db", nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	data, err := db.Get([]byte(req.Did), nil)
	didDocument := byteToString(data)

	log.Printf("Resolve DID: %s\n", req.Did)

	return &resolver.ResolveDIDRes{DidDoc: didDocument}, nil
}

// StartDIDResolverServer 서버 시작 함수
func StartDIDResolverServer(cfg *config.Config) {
	lis, err := net.Listen("tcp", cfg.Servers.Resolver.Address())

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	resolver.RegisterDIDResolverServer(s, &ResolverServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
