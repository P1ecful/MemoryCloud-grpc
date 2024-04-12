package service

import (
	"context"
	"memorycloud/proto/gen"
)

type GRPCService struct {
	gen.UnimplementedMemoryCloudServer
}

func CreateNewService() *GRPCService {
	return &GRPCService{}
}

func (s *GRPCService) Create(context.Context, *gen.CreateRequest) (*gen.Response, error) {
	return &gen.Response{}, nil
}

func (s *GRPCService) Delete(context.Context, *gen.DeleteRequest) (*gen.Response, error) {
	return &gen.Response{}, nil
}
