package client

import (
	"context"
	"fmt"
	"github.com/nam9nine/SSI_Project/config"
	"github.com/nam9nine/SSI_Project/protos/actors/verifier"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func RequestVP(vp string) (*verifier.VPRes, error) {
	cfg, err := config.LoadConfig("config/config.toml")

	if err != nil {

		return nil, fmt.Errorf("config파일 가져오기 실패 : %v", err)
	}

	addr := cfg.Servers.Verifier.Address()

	if addr == "" {
		return nil, fmt.Errorf("주소가 존재하지 않음")
	}

	conn, err := grpc.Dial("127.0.0.1:50055", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("verifier 서버와 연결 안 됨 : %v", err)
	}
	defer conn.Close()

	newClient := verifier.NewVerifierServiceClient(conn)

	res, err := newClient.VerifyVP(context.Background(), &verifier.VPReq{
		Vp: vp,
	})

	if err != nil {
		log.Printf("VerifyVP 호출 실패: %v", err)
		return nil, err
	}

	return res, nil
}
