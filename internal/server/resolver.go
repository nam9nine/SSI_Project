package server

import (
	"context"
	"errors"
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

	dbPath, err := DBPathRes(req)

	if err != nil {
		return nil, err
	}

	db, err := leveldb.OpenFile(dbPath, nil)
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
		panic(err)
	}

	s := grpc.NewServer()

	resolver.RegisterDIDResolverServer(s, &ResolverServer{})

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

func DBPathRes(req *resolver.ResolveDIDReq) (string, error) {

	cfg, err := config.LoadConfig("./././config/config.toml")

	if err != nil {
		panic(err)
	}
	var dbPath string

	switch req.Role {
	case resolver.Role_Issuer:
		dbPath = cfg.Servers.Issuer.DBPath
	case resolver.Role_Holder:
		dbPath = cfg.Servers.Holder.DBPath
	case resolver.Role_Verifier:
		dbPath = cfg.Servers.Verifier.DBPath
	default:
		return "", errors.New("invalid role")
	}
	return dbPath, nil
}
