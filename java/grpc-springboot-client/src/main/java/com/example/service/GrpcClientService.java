package com.example.service;

import lqm.demo.grpc.server.hello.HelloServiceGrpc;
import lqm.demo.grpc.server.hello.HelloServiceOuterClass;
import net.devh.boot.grpc.client.inject.GrpcClient;
import org.springframework.stereotype.Service;

@Service
public class GrpcClientService {

    @GrpcClient("local-grpc-server")
    private HelloServiceGrpc.HelloServiceBlockingStub clientStub;

    public String sendMessage() {
        final HelloServiceOuterClass.HelloReply reply = clientStub.sayHello(
                HelloServiceOuterClass.HelloRequest.newBuilder().setName("liqiming spring client").build());
        return reply.getMessage();
    }
}
