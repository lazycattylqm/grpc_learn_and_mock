package service_handler

import (
	"context"
	"fmt"
	helloPb "lqm.demo.grpc.gomicro/server/proto"
)

type Server struct{}

func (h *Server) SayHello(ctx context.Context, in *helloPb.HelloRequest, out *helloPb.HelloReply) error {
	name := in.Name
	sprintf := fmt.Sprintf("hello %s", name)
	out.Message = sprintf
	return nil
}
