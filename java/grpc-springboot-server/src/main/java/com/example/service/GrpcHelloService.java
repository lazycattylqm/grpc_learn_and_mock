package com.example.service;

import io.grpc.stub.StreamObserver;
import lombok.extern.slf4j.Slf4j;
import lqm.demo.grpc.server.hello.HelloServiceGrpc;
import lqm.demo.grpc.server.hello.HelloServiceOuterClass;
import net.devh.boot.grpc.server.service.GrpcService;
import org.springframework.stereotype.Service;

@Slf4j
@Service
@GrpcService
public class GrpcHelloService extends HelloServiceGrpc.HelloServiceImplBase {
    @Override
    public void sayHello(HelloServiceOuterClass.HelloRequest request,
                         StreamObserver<HelloServiceOuterClass.HelloReply> responseObserver) {
        final String name = request.getName();
        log.info("request name is {}", name);
        final String helloMsg = String.format("Hello %s, this is spring java server", name);
        final HelloServiceOuterClass.HelloReply reply =
                HelloServiceOuterClass.HelloReply.newBuilder().setMessage(helloMsg).build();
        responseObserver.onNext(reply);
        responseObserver.onCompleted();
    }
}
