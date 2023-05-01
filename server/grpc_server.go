package server

import (
	"context"
	"net"

	"github.com/google/uuid"
	"github.com/johnson7543/pricefetcher/proto"
	"github.com/johnson7543/pricefetcher/service"
	"google.golang.org/grpc"
)

func MakeGRPCServerAndRun(listenAddr string, svc service.PriceService) error {
	grpcPriceFetcher := NewGRPCPriceServer(svc)

	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}

	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)
	proto.RegisterPriceFetcherServer(server, grpcPriceFetcher)

	return server.Serve(ln)
}

type GRPCPriceFetcherServer struct {
	svc service.PriceService
	proto.UnimplementedPriceFetcherServer
}

func NewGRPCPriceServer(svc service.PriceService) *GRPCPriceFetcherServer {
	return &GRPCPriceFetcherServer{
		svc: svc,
	}
}

func (s *GRPCPriceFetcherServer) FetchPrice(ctx context.Context, req *proto.PriceRequest) (*proto.PriceResponse, error) {
	const uuidKey contextKey = "uuid"
	ctx = context.WithValue(ctx, uuidKey, uuid.New().String())

	price, err := s.svc.FetchPrice(ctx, req.Ticker)
	if err != nil {
		return nil, err
	}

	resp := &proto.PriceResponse{
		Ticker: req.Ticker,
		Price:  float32(price),
	}

	return resp, err

}
