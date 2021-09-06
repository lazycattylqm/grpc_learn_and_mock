package main

import (
	"client/mock"
	"context"
	"fmt"
	"github.com/asim/go-micro/v3"
	"log"
	"lqm.demo.grpc/server/proto/hello"
)

func main() {
	service := micro.NewService()
	service.Init()
	helloService := mock.NewHelloService()
	sayHello, err := helloService.SayHello(
		context.Background(), &hello.HelloRequest{
			Name: "liqiming client",
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Response is %v\n", sayHello)

}
