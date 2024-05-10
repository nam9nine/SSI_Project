package server

import (
	"github.com/nam9nine/SSI_Project/config"
	"github.com/nam9nine/SSI_Project/pkg/actors"
	"github.com/nam9nine/SSI_Project/protos/actors/issuer"
	"google.golang.org/grpc"
	"log"
	"net"
)

func StartIssuerServer(cfg *config.Config) {
	log.Println(cfg.Servers.Issuer.Address())
	lis, err := net.Listen("tcp", "127.0.0.1:50054")

	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	issuer.RegisterIssuerServiceServer(s, &actors.IssuerServer{})

	err = s.Serve(lis)

	if err != nil {
		panic(err)
	}
}
