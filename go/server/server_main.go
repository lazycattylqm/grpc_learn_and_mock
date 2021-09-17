package main

import (
	pb "example.com/liqiming/proto"
	"example.com/liqiming/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

const port = ":3000"

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("listen 3000 port error, %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterHelloServiceServer(server, &service.HelloService{})
	log.Println("server listening on ", listen.Addr())
	listenErr := server.Serve(listen)
	if listenErr != nil {
		log.Fatalf("starting server err %v", listenErr)
	}

}
