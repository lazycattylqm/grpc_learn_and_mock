package main

import (
	"github.com/asim/go-micro/v3"
	"log"
	helloPb "lqm.demo.grpc.gomicro/server/proto"
	serviceHandler "lqm.demo.grpc.gomicro/server/service_handler"
)

func main() {
	service := micro.NewService(
		micro.Name("hello.server"),
		micro.Address(":3000"),
	)
	service.Init()
	err2 := helloPb.RegisterHelloServiceHandler(service.Server(), &serviceHandler.Server{})
	if err2 != nil {
		log.Fatal(err2)
	}
	err := service.Run()
	if err != nil {
		log.Fatal(err)
	}
}
