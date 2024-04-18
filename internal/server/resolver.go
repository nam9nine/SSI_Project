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

func byte2string(b []byte) string {
	return string(b[:len(b)])
}

// ResolveDID ResolveDID(context.Context, *ResolveDIDReq) (*ResolveDIDRes, error) 구현
func (r *ResolverServer) ResolveDID(ctx context.Context, req *resolver.ResolveDIDReq) (*resolver.ResolveDIDRes, error) {
	log.Printf("Resolve DID: %s\n", req.Did)

	db, err := leveldb.OpenFile("./internal/db", nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	data, err := db.Get([]byte(req.Did), nil)
	didDocument := byte2string(data)
	return &resolver.ResolveDIDRes{DidDoc: didDocument}, nil
}

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
