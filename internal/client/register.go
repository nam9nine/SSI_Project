package client

import (
	"context"
	"errors"
	"github.com/nam9nine/SSI_Project/config"
	"github.com/nam9nine/SSI_Project/protos/vdr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func RegisterDIDDoc(did string, doc string) (*vdr.RegisterDidDocRes, error) {

	cfg, err := config.LoadConfig("config/config.toml")

	if err != nil {
		log.Printf("config파일 가져오기 실패 : %v", err)
		return nil, err
	}

	addr := cfg.Servers.VDR.Address()

	if addr == "" {
		return nil, errors.New("주소가 존재하지 않음")
	}

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, errors.New("서버와 연결 안 됨")
	}
	defer conn.Close()
	newClient := vdr.NewVDRClient(conn)

	res, err := newClient.RegisterDidDoc(context.Background(), &vdr.RegisterDidDocReq{
		Did:    did,
		DidDoc: doc,
	})

	if err != nil {
		log.Fatalf("%v", err)
	}

	return res, nil
}
