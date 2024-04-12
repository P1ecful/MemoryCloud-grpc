package service

import (
	"context"
	"memorycloud/protos/gen"
)

type GRPCService struct {
	gen.UnimplementedMemoryCloudServer
}

func CreateNewService() *GRPCService {
	return &GRPCService{}
}

// !FIXME Create methods
func (s *GRPCService) Create(context.Context, *gen.CreateRequest) (*gen.Response, error) {
	return &gen.Response{
		Status:          "200",
		RemainingWeight: "5 ГБ",
	}, nil
}

func (s *GRPCService) Delete(context.Context, *gen.DeleteRequest) (*gen.Response, error) {
	return &gen.Response{
		Status:          "200",
		RemainingWeight: "3.5 ГБ",
	}, nil
}
