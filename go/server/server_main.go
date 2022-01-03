package main

import (
	"context"
	pb "example.com/liqiming/proto"
	gatewayPb "example.com/liqiming/proto/gateway"
	"example.com/liqiming/service"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
)

const port = ":3000"

func main() {
	server2()

}

func server() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("listen 3000 port error, %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterHelloServiceServer(server, &service.HelloService{})
	fmt.Println("server listening on ", listen.Addr())
	listenErr := server.Serve(listen)
	if listenErr != nil {
		log.Fatalf("starting server err %v", listenErr)
	}
}

func server2() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("listen 3000 port error, %v", err)
	}
	s := grpc.NewServer()
	gatewayService := service.NewGatewayService()
	gatewayPb.RegisterGatewayServer(s, gatewayService)
	log.Println("Serving gRPC on 0.0.0.0:3000")
	go func() {
		log.Fatalln(s.Serve(listen))
	}()
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:3000",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	gwmux := runtime.NewServeMux()
	err = gatewayPb.RegisterGatewayHandler(context.Background(), gwmux, conn)

	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}
