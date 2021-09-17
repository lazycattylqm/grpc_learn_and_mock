package service

import (
	"context"
	pb "example.com/liqiming/proto"
	"fmt"
)

type HelloService struct {
	pb.UnimplementedHelloServiceServer
}

func (s *HelloService) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error)  {
	name := request.Name
	helloMsg := fmt.Sprintf("Hello %s, this is golang server", name)
	return &pb.HelloReply{
		Message:  helloMsg,
	}, nil
}