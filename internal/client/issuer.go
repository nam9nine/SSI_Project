package client

import (
	"context"
	"fmt"
	"github.com/nam9nine/SSI_Project/config"
	"github.com/nam9nine/SSI_Project/protos/actors/issuer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func RequestVC(did string) (*issuer.IssuerRes, error) {
	cfg, err := config.LoadConfig("config/config.toml")

	if err != nil {

		return nil, fmt.Errorf("config파일 가져오기 실패 : %v", err)
	}

	addr := cfg.Servers.Issuer.Address()

	if addr == "" {
		return nil, fmt.Errorf("주소가 존재하지 않음")
	}

	conn, err := grpc.Dial("127.0.0.1:50054", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("issuer 서버와 연결 안 됨 : %v", err)
	}
	defer conn.Close()

	newClient := issuer.NewIssuerServiceClient(conn)

	res, err := newClient.CreateUniversityVC(context.Background(), &issuer.IssuerReq{
		HolderDID: did,
	})

	if err != nil {
		log.Printf("CreateUniversityVC 호출 실패: %v", err)
		return nil, err
	}

	return res, nil
}
