package grpc

import (
	"context"
	pb "github.com/cory-johannsen/gomud/generated/mud/api"
	"github.com/cory-johannsen/gomud/internal/config"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type GrpcServer struct {
	pb.UnimplementedMudServiceServer
	address string
}

func NewGrpcServer(cfg *config.Config) *GrpcServer {
	return &GrpcServer{
		address: cfg.GrpcAddress,
	}
}

func (s *GrpcServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{}, nil
}

func StartGRPCServer(server *GrpcServer) error {
	lis, err := net.Listen("tcp", server.address)
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	pb.RegisterMudServiceServer(s, server)
	log.Printf("gRPC server listening at %v", server.address)
	return s.Serve(lis)
}
