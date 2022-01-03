package service

import (
	"context"
	"example.com/liqiming/proto/gateway"
	"fmt"
)

type gateService struct {
	gateway.UnimplementedGatewayServer
}

func NewGatewayService() *gateService {
	return &gateService{}
}

func (s gateService) PostMethod(ctx context.Context, request *gateway.Request) (*gateway.Response, error) {
	message := request.Message
	responseMsg := fmt.Sprintf("this is response, hello %s", message)
	return &gateway.Response{
		Message: responseMsg,
	}, nil
}
