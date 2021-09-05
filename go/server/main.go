package main

import (
	"github.com/asim/go-micro/v3"
	"log"
	helloPb "lqm.demo.grpc/server/proto/hello"
	helloService "lqm.demo.grpc/server/service_handler/hello"
)

func main() {
	service := micro.NewService(
		micro.Name("hello.server"),
	)
	service.Init()
	err2 := helloPb.RegisterHelloServiceHandler(service.Server(), &helloService.Server{})
	if err2 != nil {
		log.Fatal(err2)
	}
	err := service.Run()
	if err != nil {
		log.Fatal(err)
	}
}
