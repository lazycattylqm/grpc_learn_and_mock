package mock

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/v3/client"
	helloPb "lqm.demo.grpc/server/proto/hello"
)

type MockServer struct{}

func (c *MockServer) SayHello(
	ctx context.Context, in *helloPb.HelloRequest, opts ...client.CallOption,
) (*helloPb.HelloReply, error) {
	mockingMsg := fmt.Sprintf("this is mocking hello. Hello %s", in.Name)
	return &helloPb.HelloReply{
		Message: mockingMsg,
	}, nil
}

func NewHelloService() helloPb.HelloService {
	return new(MockServer)
}
