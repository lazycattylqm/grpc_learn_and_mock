package main

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/v3"
	"log"
	helloPb "lqm.demo.grpc.gomicro/server/proto"
)

func main() {
	service := micro.NewService()
	service.Init()
	helloService := helloPb.NewHelloService("hello.server", service.Client())
	sayHello, err := helloService.SayHello(
		context.Background(), &helloPb.HelloRequest{
			Name: "liqiming client",
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Response is %v\n", sayHello)

}
