package main

import (
	"context"
	"encoding/json"
	pb "example.com/liqiming/proto"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

const (
	addr = "localhost:3000"
	name = "liqiming"
)

func main() {
	dial, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("grpc dial failed %v", err)
	}
	defer dial.Close()
	client := pb.NewHelloServiceClient(dial)
	hello, serviceErr := client.SayHello(
		context.Background(), &pb.HelloRequest{
			Name: name,
		},
	)
	if serviceErr != nil {
		log.Fatalf("calling service error %v", serviceErr)
	}
	marshal, jsonErr := json.Marshal(hello)
	if jsonErr != nil {
		log.Fatalf("json error ,  %v", jsonErr)
	}
	fmt.Printf("service response is %v\n", string(marshal))

}
