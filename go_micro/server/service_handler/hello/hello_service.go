package hello

import (
	"context"
	"fmt"
	helloPb "lqm.demo.grpc/server/proto/hello"
)

type Server struct{}

func (h *Server) SayHello(ctx context.Context, in *helloPb.HelloRequest, out *helloPb.HelloReply) error {
	name := in.Name
	sprintf := fmt.Sprintf("hello %s", name)
	out.Message = sprintf
	return nil
}
