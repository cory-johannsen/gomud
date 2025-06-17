package rest

import (
	"context"
	gw "github.com/cory-johannsen/gomud/generated/mud/api"
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net/http"
)

type RestServer struct {
	GrpcAddress string
	RestAddress string
}

func NewRestServer(cfg *config.Config) *RestServer {
	return &RestServer{
		GrpcAddress: cfg.GrpcAddress,
		RestAddress: cfg.RestAddress,
	}
}

func StartRestServer(s *RestServer) error {
	ctx := context.Background()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	mux := runtime.NewServeMux()
	err := gw.RegisterMudServiceHandlerFromEndpoint(ctx, mux, s.GrpcAddress, opts)
	if err != nil {
		return err
	}
	log.Printf("Swagger UI server listening at %s", s.RestAddress)
	return http.ListenAndServe(s.RestAddress, mux)
}
