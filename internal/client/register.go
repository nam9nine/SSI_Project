package client

import (
	"context"
	"errors"
	"github.com/nam9nine/SSI_Project/config"
	registrar "github.com/nam9nine/SSI_Project/protos/vdr/registrar"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

// RegistrarDID DID 등록 요청 코드 with generics
func RegistrarDID(did string, doc string, role registrar.Role) (*registrar.DIDRegistrarRes, error) {

	cfg, err := config.LoadConfig("config/config.toml")

	if err != nil {
		log.Printf("config파일 가져오기 실패 : %v", err)
		return nil, err
	}

	addr := cfg.Servers.Registrar.Address()

	if addr == "" {
		return nil, errors.New("주소가 존재하지 않음")
	}

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, errors.New("registrar 서버와 연결 안 됨")
	}
	defer conn.Close()

	newClient := registrar.NewDIDRegistrarClient(conn)

	// Convert role to string or handle serialization based on actual type requirements

	res, err := newClient.RegisterDidDoc(context.Background(), &registrar.DIDRegistrarReq{
		Did:    did,
		DidDoc: doc,
		Role:   role,
	})

	if err != nil {
		log.Fatalf("%v", err)
	}

	return res, nil
}
