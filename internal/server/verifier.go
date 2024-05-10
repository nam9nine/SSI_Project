package server

import (
	"github.com/nam9nine/SSI_Project/config"
	"github.com/nam9nine/SSI_Project/pkg/actors"
	"github.com/nam9nine/SSI_Project/protos/actors/verifier"
	"google.golang.org/grpc"
	"net"
)

func StartVerifierServer(cfg *config.Config) {

	lis, err := net.Listen("tcp", "127.0.0.1:50055")

	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	verifier.RegisterVerifierServiceServer(s, &actors.VerifierServer{})

	err = s.Serve(lis)

	if err != nil {
		panic(err)
	}
}
