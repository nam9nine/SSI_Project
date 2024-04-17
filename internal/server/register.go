package server

import (
	"context"
	"github.com/nam9nine/SSI_Project/config"
	"github.com/nam9nine/SSI_Project/protos/vdr"
	"github.com/syndtr/goleveldb/leveldb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type RegisterDIDDoc struct {
	vdr.UnimplementedVDRServer
}

// RegisterDidDoc RegisterDidDoc(context.Context, *RegisterDidDocReq) (*RegisterDidDocRes, error) 구현
func (r *RegisterDIDDoc) RegisterDidDoc(ctx context.Context, req *vdr.RegisterDidDocReq) (*vdr.RegisterDidDocRes, error) {
	reqDocRes := vdr.RegisterDidDocRes{}
	db, err := leveldb.OpenFile("./db", nil)
	if err != nil {
		reqDocRes.State = vdr.State_UNKNOWN
		panic(err)
	}
	defer db.Close()

	err = db.Put([]byte(req.Did), []byte(req.DidDoc), nil)

	if err != nil {
		reqDocRes.State = vdr.State_FAILURE
	}
	reqDocRes.State = vdr.State_SUCCESS
	reqDocRes.Message = "good"
	return &reqDocRes, nil
}

// StartRegisterServer grpc 서버 생성
func StartRegisterServer() {
	cfg, err := config.LoadConfig("config/config.toml")

	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", cfg.Servers.VDR.Address())

	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	vdr.RegisterVDRServer(s, &RegisterDIDDoc{})
	err = s.Serve(lis)

	if err != nil {
		panic(err)
	}

}
