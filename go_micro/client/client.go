package main

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/v3"
	"log"
	helloPb "lqm.demo.grpc.gomicro/proto"
)

func main() {
	service := micro.NewService(
		micro.Name("hello.server"),
		micro.Address(":3000"),
		)
	service.Init()
	helloService := helloPb.NewHelloService("hello.server", service.Client())
	sayHello, err := helloService.SayHello(
		context.Background(), &helloPb.HelloRequest{
			Name: "liqiming client go micro",
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Response is %v\n", sayHello)

}
