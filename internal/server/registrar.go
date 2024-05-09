package server

import (
	"context"
	"github.com/nam9nine/SSI_Project/config"
	registrar "github.com/nam9nine/SSI_Project/protos/vdr/registrar"
	"github.com/syndtr/goleveldb/leveldb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type RegistrarServer struct {
	registrar.UnimplementedDIDRegistrarServer
}

// RegisterDidDoc RegisterDidDoc(context.Context, *DIDRegistrarReq) (*DIDRegistrarRes, error)
func (r *RegistrarServer) RegisterDidDoc(ctx context.Context, req *registrar.DIDRegistrarReq) (*registrar.DIDRegistrarRes, error) {

	res := registrar.DIDRegistrarRes{}
	db, err := leveldb.OpenFile("./internal/db", nil)
	if err != nil {
		res.State = registrar.State_UNKNOWN
		panic(err)
	}
	defer db.Close()

	err = db.Put([]byte(req.Did), []byte(req.DidDoc), nil)

	if err != nil {
		res.State = registrar.State_FAILURE
	}
	res.State = registrar.State_SUCCESS
	res.Message = "good"

	log.Printf("registrar DID: %s\n", req.Did)
	return &res, nil
}

// StartRegisterServer grpc 서버 생성
func StartRegisterServer(cfg *config.Config) {

	lis, err := net.Listen("tcp", cfg.Servers.Registrar.Address())

	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	registrar.RegisterDIDRegistrarServer(s, &RegistrarServer{})

	err = s.Serve(lis)

	if err != nil {
		panic(err)
	}
}
